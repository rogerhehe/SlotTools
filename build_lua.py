#%%
import os,re

proto_map = {}
pattern = re.compile(r'.+\.proto$')
out_dir = './gen/lua'
lua_proto_import_path = 'Game.Common.Proto'
proto_lua_head = """local proto = {"""
proto_lua_tail = """ }\nreturn proto """


os.makedirs(out_dir,exist_ok=True)
for file in os.listdir('./proto'):
    if re.match(pattern,file):
        print(f'gen lua file {file} -> ')
    

#%%

__pkg_pattern = re.compile(r'(?<=package )(\S+)(?=;)')
def to_lua(content):
    __search = __pkg_pattern.search(content)
    __pkg = ''
    if __search:
        __pkg = __pkg_pattern.search(content).group() #match first
    
    



#%% test script

#%%

with open('proto/msg.proto','r') as f:
    file = f.read()


#%%
file = 'package pb.base; \n package pb.rr;'
__pkg_pattern = re.compile(r'(?<=package )(\S+)(?=;)')
__pkg_pattern.search(file)
# %%
