# proto out go py
out_dir='gen'

# lua go gen
rm -rf $out_dir

python3 build_lua.py

mkdir ./$out_dir/go
mkdir ./$out_dir/py
protoc --go_out=./$out_dir/go/ ./proto/*.proto
protoc --python_out=./$out_dir/py/ ./proto/*.proto
#protoc --csharp_out=./cs/ *.proto

#TODO copy files to repo