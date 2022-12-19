#!/bin/bash
protoc --go_out=. *.proto

base_path=$(cd `dirname $0`; pwd)
src_dir_path=$base_path"/pb"

cd `dirname $0`/..
out_dir_path=$(cd `dirname $0`; dirname $(pwd))
out_dir_path=$out_dir_path"/SlotServer/common"

cp -r $src_dir_path $out_dir_path
rm -r $src_dir_path