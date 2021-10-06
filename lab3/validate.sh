#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab3-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

#rm -rf *
cp $solution_path/lab3.go .
result=$(go run lab3.go 2>&1) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab3.go\" fails ; NO POINT"
else
  echo "\"go run lab4.go\" output: \"$result\""
  if [ $(python3 -c 'print("沒夾到喔！\n"*50)') != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0