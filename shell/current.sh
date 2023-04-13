#!/bin/bash
# 現在解いている問題を表示します

site=`cat .current/site 2> /dev/null`
url=`cat .current/url 2> /dev/null`
title=`cat .current/title 2> /dev/null`
problem=`cat .current/problem 2> /dev/null`

if [ -z $site -o -z $url -o -z $title -o -z $problem ]; then
  echo コンテストが開始されていません。init.shでコンテストを開始してください。
  exit 1
fi

if [ $site = atcoder ]; then
  problem_url="${url}/tasks/${problem}"
fi

echo 現在の問題：${problem_url}
