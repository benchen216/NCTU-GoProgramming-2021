#!/bin/bash
export PATH=$PATH:/usr/bin/go/bin
test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab3-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

rm -rf *
cp $solution_path/lab3.go .
go mod init $(basename $tmp_dir)
go get gopl.io/ch4/github
go version
go run lab3.go &
ps -al
if [ "$(curl --retry-connrefused --retry 10 --connect-timeout http://localhost:8080 )" != "$(curl http://nctu.is-geek.com:8083/ )" ] ; then
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0