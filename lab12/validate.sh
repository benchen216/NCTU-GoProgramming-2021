#!/bin/bash
export PATH=$PATH:/usr/bin/go/bin
test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab12-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

rm -rf * 
cp $solution_path/* .
go version
pwd
ls
chromium --version
chromedriver --version
go run lab12.go command >out 2>&1 &
curl --retry-connrefused --retry 4 --connect-timeout 5 http://0.0.0.0:8899 --verbose
wget http://0.0.0.0:8899/index.html
cat index.html
python3  $solution_path/../validate.py

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0