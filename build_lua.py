#%%
import os,re

proto_map = {}
pattern = re.compile(r'.+\.proto$')
out_dir = './gen/lua'
lua_proto_import_path = 'Game.Common.Proto'
#proto_lua_head = """local proto = {"""
#proto_lua_tail = """ }\nreturn proto """
file_fmt = """local protoc = require "Core.Protobuf.Protoc"

assert(protoc:load [[
{0}
]])"""

limporter = []

os.makedirs(out_dir,exist_ok=True)

#%%

__pkg_pattern = re.compile(r'(?<=package )(\S+)(?=;)')
def to_lua(content,filename):
    __search = __pkg_pattern.search(content)
    __pkg = ''
    if __search:
        __pkg = __pkg_pattern.search(content).group() #match first
        __fn = f"{__pkg.replace('.','_')}_{filename}"
        __dir = os.path.join(out_dir,f"{__fn}.lua")
        content = re.sub(r'//.+','',content)
        content = re.sub(r'(?<=\n\n)(\n+)','',content)
        with open(__dir,'w') as f:
            f.write(file_fmt.format(content))
        print(f'gen lua file {filename} -> {__dir}')
        f.close()
        limporter.append(f'require "{lua_proto_import_path}.{__fn}"\n')
#%%
for file in os.listdir('./proto'):
    if re.match(pattern,file):
        with open(f'./proto/{file}','r') as f:
            content = f.read()
            to_lua(content,file.split('.')[0])     

with open(f'{out_dir}/importer.lua','w')as f:
    f.writelines(limporter)

