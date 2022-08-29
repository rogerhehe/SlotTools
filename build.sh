# proto out go py
out_dir='gen'

# lua go gen
mkdir  gen/go
mkdir  gen/py
protoc --go_out=./gen/go/ ./proto/*.proto
protoc --python_out=./gen/py/ ./proto/*.proto

# 

#protoc --csharp_out=./cs/ *.proto



#TODO copy files to repo