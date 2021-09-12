#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab1-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

#rm -rf *
cp $solution_path/lab1.go .
action=1
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$((a+b))
result=$(printf '%d\n%d %d' $action $a $b | go run lab1.go 2>&1 | tail -1 ) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab1.go\" fails ; NO POINT"
else
  echo "\"go run lab1.go\" output: \"$result\""
  if [ "$(echo $ans)" != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi
action=2
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$((a-b))
result=$(printf '%d\n%d %d' $action $a $b | go run lab1.go 2>&1 | tail -1 ) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab1.go\" fails ; NO POINT"
else
  echo "\"go run lab1.go\" output: \"$result\""
  if [ "$(echo $ans)" != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi
action=3
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$((a*b))
result=$(printf '%d\n%d %d' $action $a $b | go run lab1.go 2>&1 | tail -1 ) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab1.go\" fails ; NO POINT"
else
  echo "\"go run lab1.go\" output: \"$result\""
  if [ "$(echo $ans)" != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi
action=4
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$((a/b))
result=$(printf '%d\n%d %d' $action $a $b | go run lab1.go 2>&1 | tail -1 ) ; ret=$?
if [ $ret -ne 0 ] ; then
  echo "\"go run lab1.go\" fails ; NO POINT"
else
  echo "\"go run lab1.go\" output: \"$result\""
  if [ "$(echo $ans)" != "$(echo $result)" ] ; then
    echo "wrong answer ; NO POINT"
  else
    echo "GET POINT 1"
  fi
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0