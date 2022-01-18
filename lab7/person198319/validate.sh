#!/bin/bash

test_path="${BASH_SOURCE[0]}"
solution_path="$(realpath .)"
tmp_dir=$(mktemp -d -t lab7-XXXXXXXXXX)

echo "working directory: $tmp_dir"
cd $tmp_dir

#rm -rf *
cp $solution_path/app_url.txt .
ans=$(cat <<-END
[
    {
        "id": "1",
        "name": "Blue Bird",
        "pages": "500"
    }
]
END
)
curl -o result.txt `cat app_url.txt`bookshelf
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "id": "1",
    "name": "Blue Bird",
    "pages": "500"
}
END
)
curl -o result.txt `cat app_url.txt`bookshelf/1
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "message": "book not found"
}
END
)
curl -o result.txt `cat app_url.txt`bookshelf/2
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "id": "2",
    "name": "Pride and Prejudice",
    "pages": "600"
}
END
)
curl -X POST -H 'Content-Type: application/json' -d '{"ID":"2","NAME":"Pride and Prejudice","PAGES":"600"}' -o result.txt `cat app_url.txt`bookshelf
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "id": "2",
    "name": "Pride and Prejudice",
    "pages": "600"
}
END
)
curl -o result.txt `cat app_url.txt`bookshelf/2
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "id": "3",
    "name": "原子習慣：細微改變帶來巨大成就的實證法則",
    "pages": "33"
}
END
)
curl -X POST -H 'Content-Type: application/json' -d '{"ID":"3","NAME":"原子習慣：細微改變帶來巨大成就的實證法則","PAGES":"33"}' -o result.txt `cat app_url.txt`bookshelf
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "message": "duplicate book id"
}
END
)
curl -X POST -H 'Content-Type: application/json' -d '{"ID":"3","NAME":"原子習慣：細微改變帶來巨大成就的實證法則","PAGES":"33"}' -o result.txt `cat app_url.txt`bookshelf
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "id": "3",
    "name": "原子習慣：細微改變帶來巨大成就的實證法則",
    "pages": "600"
}
END
)

curl -X PUT -H 'Content-Type: application/json' -d '{"ID":"3","NAME":"原子習慣：細微改變帶來巨大成就的實證法則","PAGES":"600"}' -o result.txt `cat app_url.txt`bookshelf/3
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

ans=$(cat <<-END
{
    "id": "3",
    "name": "原子習慣：細微改變帶來巨大成就的實證法則",
    "pages": "600"
}
END
)
curl -X DELETE  -o result.txt `cat app_url.txt`bookshelf/3
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi



ans=$(cat <<-END
{
    "message": "book not found"
}
END
)
curl -X DELETE  -o result.txt `cat app_url.txt`bookshelf/3
echo $ans > ans.txt
DIFF=$(diff <(jq -S . result.txt) <(jq -S . ans.txt))
if [ "$DIFF" != "" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi

echo "deleting working directory $tmp_dir"
rm -rf $tmp_dir

exit 0
