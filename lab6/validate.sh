#!/bin/bash
export PATH=$PATH:/usr/bin/go/bin
test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab6-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

rm -rf *
cp $solution_path/* .
go version
pwd
ls
chromium --version
go run server.go & python3  $solution_path/../validate.py

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0