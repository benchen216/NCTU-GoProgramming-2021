#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab5-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

#rm -rf *
cp $solution_path/app_url.txt .
a=$((1 + RANDOM % 100))
b=$((1 + RANDOM % 100))
ans=$a" + "$b" = "$((a+b))
curl -o result.txt `cat app_url.txt`"?op=add&num1=$a&num2=$b"
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

a=$((1 + RANDOM % 100))
b=$((1 + RANDOM % 100))
ans=$a" - "$b" = "$((a-b))
curl -o result.txt `cat app_url.txt`"?op=sub&num1=$a&num2=$b"
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi
a=$((1 + RANDOM % 100))
b=$((1 + RANDOM % 100))
ans=$a" * "$b" = "$((a*b))
curl -o result.txt `cat app_url.txt`"?op=mul&num1=$a&num2=$b"
if [ "$(echo "$ans")" != "$(cat result.txt)" ] ; then
  echo "right ans=$(echo "$ans")"
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi
a=$((1 + RANDOM % 100))
b=$((1 + RANDOM % 100))
ans=$a" / "$b" = "$((a/b))", remainder = "$((a%b))
curl -o result.txt `cat app_url.txt`"?op=div&num1=$a&num2=$b"
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

a=$((1 + RANDOM % 100))
b=$((1 + RANDOM % 100))
m=$a
if [ $b -lt $m ]
then
m=$b
fi
while [ $m -ne 0 ]
do
x=`expr $a % $m`
y=`expr $b % $m`
if [ $x -eq 0 -a $y -eq 0 ]
then
#echo gcd of $a and $b is $m
ans=$m
curl -o result.txt `cat app_url.txt`"?op=gcd&num1=$a&num2=$b"
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi
break
fi
m=`expr $m - 1`
done


echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0