#!/bin/bash
export PATH=$PATH:/usr/bin/go/bin
test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab12-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

rm -rf * 
cp -r $solution_path/* .
go version
pwd
ls
chromium --version
chromedriver --version
chmod 777 static
chmom 777 static/index.html
go run lab12.go command >out 2>&1 &
curl --retry-connrefused --retry 4 --connect-timeout 5 http://0.0.0.0:8899/ --verbose
echo "wget"
wget http://0.0.0.0:8899/index.html
echo "cat index.html"
cat index.html
echo "cat out"
cat out
ls -al
ls -al static
python3  $solution_path/../validate.py

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0