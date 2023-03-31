# %%
try:
    import os,math
    import xmltodict,easydict
except:
    os.system("pip3 install xmltodict easydict")
    import xmltodict,easydict

## todo shell传进来
XML_DIR = "./proto/msg_map"
OUT_DIR = "./out/protomap"

os.makedirs(f'{OUT_DIR}/go/',exist_ok=True)
os.makedirs(f'{OUT_DIR}/lua/',exist_ok=True)

# %%

## MsgID.lua
cli_id_title="""MsgID = {{
	{0}
}}
"""

## MsgPB.lua
cli_proto_title="""require "Assets/_LuaScripts/Game/Net/Proto/Importer"

local pb_scheme = {{
	{0}
}}

return pb_scheme
"""

## msg_name.go
srv_msg_title = """package msgid

var Name = map[uint16]string{{
	{0}
}}
"""

## *_msgproto.go
srv_proto_title = """package msgproto

import (
	"slotserver/common/msgid"
	"slotserver/common/pb"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var {protoName}MsgProto = map[uint16]protoreflect.MessageType{{
	{proto}
}}
"""

## *_msgid.go
srv_id_title = """package msgid

const (
	{0}
)
"""

fmt_name = "{a}2{b}_{name}"

fmt_cli_id = "{pname:{maxlen}s} = {pid},"
#fmt_cli_proto_name = '{pname}= "{pname}"'
fmt_cli_proto = '[MsgID.{pname}] = "{proto}",'

fmt_srv_id = "{pname:{maxlen}s} = {pid}"
fmt_srv_proto_name = '{pname}:{space:{maxlen}s}"{pname}",'
fmt_srv_proto = "msgid.{pname}:{space:{maxlen}s}(&{proto}).ProtoReflect().Type(),"




# %%

lcid = []
lcproto = []
lsid = []
lsprotoname = []
lsproto = []

for fn in ["backend","frontend"]:
    parse = {}
    
    with open(f"{XML_DIR}/{fn}.xml",'r') as f:
        parse = easydict.EasyDict(xmltodict.parse(f.read()))
    for cfg in parse.protoCfg.cfg:
        maxlen = 0
        up = cfg.get("@up",'')
        down = cfg.get("@down",'')
        resp_offset = int(cfg.get("@resp_offset",0))
        ctype = cfg.get("@type","normal")
        dsc = cfg.get("@desc","")
        protos = cfg.proto
        if not isinstance(protos,list):
            protos = [protos]
        
        lcid.append(f'-- {dsc}')
        lcproto.append(f'-- {dsc}')
        lsid.append(f'//{dsc}')
        lsproto.append(f'//{dsc}')
        lsprotoname.append(f'//{dsc}')
        
        for proto in protos:
            proto_id = int(proto['@id'])
            name = proto['@name']
            req = proto.get('@req', None)
            resp = proto.get('@resp', None)
            if req:
                pname = fmt_name.format( a = up, b = down, name = name)
                maxlen = max(len(pname), maxlen)
            if resp:
                pname = fmt_name.format( a = down, b = up, name = name)
                maxlen = max(len(pname), maxlen)

        for proto in protos:
            proto_id = int(proto['@id'])
            name = proto['@name']
            req = proto.get('@req', None)
            resp = proto.get('@resp', None)
            if req:
                pname = fmt_name.format( a = up, b = down, name = name)
                lendiff = max(0,(maxlen - len(pname))) + 1
                cid = fmt_cli_id.format( pname = pname, pid = str(proto_id), maxlen = maxlen)
                cproto = fmt_cli_proto.format( pname = pname, proto = req, maxlen = maxlen)
                sid = fmt_srv_id.format( pname = pname, pid = str(proto_id), maxlen = maxlen)
                sproto = fmt_srv_proto.format( pname = pname, proto = req, maxlen = lendiff, space = "")
                sprotoname = fmt_srv_proto_name.format(pname = pname, space = "", maxlen=lendiff) 

                lcid.append(cid)
                lcproto.append(cproto)
                lsid.append(sid)
                lsproto.append(sproto)
                lsprotoname.append(sprotoname)
            if resp:
                pname = fmt_name.format( a = down, b = up, name = name)
                pid = str( proto_id + resp_offset )
                lendiff = max(0,(maxlen - len(pname))) + 1
                cid = fmt_cli_id.format( pname = pname, pid = pid, maxlen = maxlen)
                cproto = fmt_cli_proto.format( pname = pname, proto = req, maxlen = maxlen)
                sid = fmt_srv_id.format( pname = pname, pid = pid, maxlen = maxlen)
                sproto = fmt_srv_proto.format( pname = pname, proto = resp, maxlen = lendiff, space = "")
                sprotoname = fmt_srv_proto_name.format(pname = pname, space = "", maxlen = lendiff) 
                lcid.append(cid)
                lcproto.append(cproto)
                lsid.append(sid)
                lsproto.append(sproto)
                lsprotoname.append(sprotoname)
        

    

    with open(f'{OUT_DIR}/go/{fn}_msgid.go','w') as f:
        f.write(srv_id_title.format(str.join('\n\t',lsid)))

    with open(f'{OUT_DIR}/go/{fn}_msgproto.go','w') as f:
        f.write(srv_proto_title.format(protoName=fn[0].upper()+fn[1:].lower(), proto=str.join('\n\t',lsproto)))

    if fn == "frontend":
        with open(f'{OUT_DIR}/lua/MsgID.lua','w') as f:
            f.write(cli_id_title.format(str.join('\n\t',lcid)))

        with open(f'{OUT_DIR}/lua/MsgPB.lua','w') as f:
            f.write(cli_proto_title.format(str.join('\n\t',lcproto)))
    lcid = []
    lcproto = []
    lsid = []
    lsproto = []

# %%

with open(f'{OUT_DIR}/go/msg_name.go','w') as f:
    f.write(srv_msg_title.format(str.join('\n\t',lsprotoname)))


# %%



