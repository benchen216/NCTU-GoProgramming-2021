#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab3-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

rm -rf *
cp $solution_path/lab3.go .
go mod init $(basename $tmp_dir)
go get github.com/adonovan/gopl.io/ch4/github
go run lab3.go&
if [ "$(curl http://localhost:8080/ )" != "$(curl http://nctu.is-geek.com:8083/ )" ] ; then
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0