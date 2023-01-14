#!/bin/bash
# コンテストの回答を始める初期化をおこなう
if [ $# = 0 ]; then
  echo -n 開始するコンテストのURLを指定してください:
  read url
  echo -n 開始するコンテストの問題を指定してください:
  read problem
elif [ $# = 1 ]; then
  url=$1
  problem=a
else
  url=$1
  problem=$2
fi

readonly URL_ATCODER='^https://atcoder\.jp/contests/([a-z0-9\-]*)$'

if [[ $url =~ $URL_ATCODER ]] ; then
  contest_site=atcoder
  contest_title=$(echo ${BASH_REMATCH[1]} | sed -e s/-/_/)
else
  echo 対応していないURLです
  exit 1
fi

rm -f .current/problem

# 現在のコンテストの情報を.currentディレクトリに保持する
mkdir -p .current
echo -n $url > .current/url
echo -n $contest_site > .current/site
echo -n $contest_title > .current/title

bash shell/new.sh $problem
