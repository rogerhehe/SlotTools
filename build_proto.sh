# proto out go py
POJECT_PATH="../UnityClient"


src_dir='proto'
out_dir='out/protogen'

CURPATH=$(cd "$(dirname "$0")"; pwd)
out_path=$CURPATH/$out_dir
# client
C_OUTPUT_PATH=$POJECT_PATH'/Assets/_LuaScripts/Game/Net/Proto'
# lua go gen
rm -rf $out_dir

python3 $CURPATH/bin/gen_proto.py --dir=$out_path --script_path='/Assets/_LuaScripts/Game/Net/Proto'

mkdir $out_dir/go
mkdir $out_dir/py
mkdir $out_dir/cs
protoc --go_out=$out_dir/go/ $src_dir/*.proto
protoc --python_out=$out_dir/py/ $src_dir/*.proto
protoc --csharp_out=$out_dir/cs/ $src_dir/*.proto

#TODO copy files to repo
rm -rf $C_OUTPUT_PATH 
cp -r $out_dir/lua $C_OUTPUT_PATH

echo "CPOY CONFIG DONE"