
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
curl -X PUT -H 'Content-Type: application/json' -d '{"ID":"3","NAME":"原子習慣：細微改變帶來巨大成就的實證法則","PAGES":"600"}' -o result.txt `cat app_url.txt`bookshelf/2
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
curl -X DELETE  -o result.txt `cat app_url.txt`bookshelf/3
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
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
if [ "$(echo $ans)" != "$(cat result.txt)" ] ; then
  echo "right ans="$ans
  echo "your ans=$(cat result.txt)"
  echo "wrong answer ; NO POINT"
else
  echo "GET POINT 1"
fi
