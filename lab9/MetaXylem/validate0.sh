#!/bin/bash
export PATH=$PATH:/usr/bin/go/bin
test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab9-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir
cp -r $solution_path/* .

count=0
echo "==========================================="
echo "Action1: go run lab9.go"
go run lab9.go &> result.txt
DIFF=$(diff result.txt ./ans/ans1.txt 2>&1)
diff -y result.txt ./ans/ans1.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi

echo "==========================================="
echo "Action2: go run lab9.go 999"
go run lab9.go 999&> result.txt
DIFF=$(diff result.txt ./ans/ans2.txt 2>&1)
diff -y result.txt ./ans/ans2.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi

echo "==========================================="
a='10'
b='10'
c='蔡英文'
d='韓國瑜'
e='宋楚瑜'
echo "Action3: go run lab9.go $a $b $c $d $e"
go run lab9.go $a $b $c $d $e &> result.txt
DIFF=$(diff result.txt ./ans/ans3.txt 2>&1)
diff -y result.txt ./ans/ans3.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi

echo "==========================================="
a='2'
b='7'
c='日本'
d='台灣'
e='中國'
f='香港'
echo "Action4: go run lab9.go $a $b $c $d $e $f"
go run lab9.go $a $b $c $d $e $f&> result.txt
DIFF=$(diff result.txt ./ans/ans4.txt 2>&1)
diff -y result.txt ./ans/ans4.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi

echo "==========================================="
a='3'
b='4'
c='柯文哲'
echo "Action5: go run lab9.go $a $b $c"
go run lab9.go $a $b $c &> result.txt
DIFF=$(diff result.txt ./ans/ans5.txt 2>&1)
diff -y result.txt ./ans/ans5.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi

echo "==========================================="
a='4'
b='5'
c='愛莉莎莎'
echo "Action6: go run lab9.go $a $b $c"
go run lab9.go $a $b $c &> result.txt
DIFF=$(diff result.txt ./ans/ans6.txt 2>&1)
diff -y result.txt ./ans/ans6.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi

echo "==========================================="
echo "Action7: go run lab9.go 1 1"
go run lab9.go 1 1 &> result.txt
DIFF=$(diff result.txt ./ans/ans7.txt 2>&1)
diff -y result.txt ./ans/ans7.txt > diff.txt
if [ "$DIFF" != ""  ] ; then
  echo "Diff="
  cat diff.txt
  echo "Wrong Answer ; NO POINT"
else
  echo "GET POINT 1"
  count=$(($count + 1))
fi


echo "==========================================="
echo "Pass: "$count"/7"


echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

read -n 1 -p "Press any key to continue..."

exit 0