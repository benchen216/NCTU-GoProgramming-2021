#!/bin/bash
export PATH=$PATH:/usr/bin/go/bin
test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab11-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

cp $solution_path/* .
go run lab11.go &> /dev/null
go run lab11.go &> result1.txt 
go run lab11.go -w &> result2.txt
go run lab11.go -max &> result3.txt
go run lab11.go -max 5 &> result4.txt
go run lab11.go -w ntu &> result5.txt
go run lab11.go -w ntu -max 3 &> result6.txt
go run lab11.go -max 100 -w ntu &> result7.txt
go run lab11.go -max 30 -w ptt &> result8.txt
python3 $solution_path/../validate.py

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0