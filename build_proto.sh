# proto out go py
out_dir='protogen'

# lua go gen
rm -rf $out_dir

python3 gen_proto.py --dir=$out_dir

mkdir ./$out_dir/go
mkdir ./$out_dir/py
mkdir ./$out_dir/cs
protoc --go_out=./$out_dir/go/ ./proto/*.proto
protoc --python_out=./$out_dir/py/ ./proto/*.proto
protoc --csharp_out=./$out_dir/cs/ ./proto/*.proto

#TODO copy files to repo