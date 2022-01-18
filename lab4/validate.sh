#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab4-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

#rm -rf *
cp $solution_path/app_url.txt .
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$a" + "$b" = "$((a+b))
curl -o result.txt `cat app_url.txt`add/$a/$b
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$a" - "$b" = "$((a-b))
curl -o result.txt `cat app_url.txt`sub/$a/$b
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$a" * "$b" = "$((a*b))
curl -o result.txt `cat app_url.txt`mul/$a/$b
if [ "$(echo "$ans")" != "$(cat result.txt)" ] ; then
  echo "right ans=$(echo "$ans")"
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi
a=$((1 + RANDOM % 10))
b=$((1 + RANDOM % 10))
ans=$a" / "$b" = "$((a/b))", remainder = "$((a%b))
curl -o result.txt `cat app_url.txt`div/$a/$b
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0