#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab2-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir
cp $solution_path/../*.txt .
#rm -rf *
cp $solution_path/lab2.go .
result=$(cat input1.txt | go run lab2.go 2>&1) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab2.go\" fails ; NO POINT"
else
  echo "\"go run lab2.go\" output: \"$result\""
  if [ "$(cat output1.txt)" != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi

result=$(cat input2.txt | go run lab2.go > answer.txt) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab2.go\" fails ; NO POINT"
else
  echo "\"go run lab2.go\" output: \"$(cat answer.txt)\""
  if [ "$(cat output2.txt)" != "$(cat answer.txt)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi

result=$(cat input3.txt | go run lab2.go > answer.txt) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab2.go\" fails ; NO POINT"
else
  echo "\"go run lab2.go\" output: \"$(cat answer.txt)\""
  if [ "$(cat output3.txt)" != "$(cat answer.txt)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi


echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0