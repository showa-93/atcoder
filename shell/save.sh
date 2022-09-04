#!/bin/bash

site=`cat .current/site 2> /dev/null`
url=`cat .current/url 2> /dev/null`
title=`cat .current/title 2> /dev/null`
problem=`cat .current/problem 2> /dev/null`

if [ -z $site -o -z $url -o -z $title -o -z $problem ]; then
  echo コンテストが開始されていません。init.shでコンテストを開始してください。
  exit 1
fi

if [ $site = atcoder ]; then
  problem_url="${url}/tasks/${title}_${problem}"
fi

problem_directory="contests/${site}/${title}/${problem}"
rm -rf $problem_directory
mkdir -p $problem_directory

cp main.go ${problem_directory}/main.go
cp -r testdata ${problem_directory}/
echo ${problem_url} > ${problem_directory}/README.md

rm main.go
rm -rf testdata
