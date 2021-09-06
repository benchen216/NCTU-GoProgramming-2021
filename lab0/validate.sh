#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab0-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

#rm -rf *
cp $solution_path/helloworld.go .
result=$(go run helloworld.go 2>&1) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run helloworld.go\" fails ; NO POINT"
else
  echo "\"go run helloworld.go\" output: \"$result\""
  if [ "hello world" != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi

cp $solution_path/username.go .
result=$(go run username.go 2>&1) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run username.go\" fails ; NO POINT"
else
  echo "\"go run username.go\" output: \"$result\""
  if [ "$(basename $solution_path)" != "$(echo $result)" ] ; then
    echo "$(basename $solution_path) wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0