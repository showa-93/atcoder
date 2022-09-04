#!/bin/bash
# 現在のコンテストをもとにテストケースを初期化する
# main.goのコピー、テストデータの取得を行う
if [ $# != 1 ]; then
  echo 開始するコンテストの問題を指定してください
  exit 1
fi

site=`cat .current/site 2> /dev/null`
url=`cat .current/url 2> /dev/null`
title=`cat .current/title 2> /dev/null`

if [ -z $site -o -z $url -o -z $title ]; then
  echo コンテストが開始されていません。init.shでコンテストを開始してください。
  exit 1
fi

if [ $site = atcoder ]; then
  problem_url="${url}/tasks/${title}_${1}"
fi

if [ -d testdata -o -e main.go ]; then
  # 間違ってさよならしないようにバックアップを取得する
  backup=.backup/program/`date "+%y%m%d%H%M%S"`
  mkdir -p $backup
  cp main.go ${backup}/main.go
  cp -r testdata ${backup}/testdata
  rm -rf testdata
  rm -f .current/problem
fi

oj d -d testdata -f "case%i/%e" $problem_url
cp template/main.go main.go
echo -n $1 > .current/problem
