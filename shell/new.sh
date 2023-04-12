#!/bin/bash
# 現在のコンテストをもとにテストケースを初期化する
# main.goのコピー、テストデータの取得を行う
. shell/contests_function.sh

case $# in
  1)
    problem_url=`get_problem_url ${1}`
    ;;
  2)
    problem_url=`get_problem_url ${1} ${2}`
    ;;
  *)
    echo 開始するコンテストの問題を指定してください
    exit 1
    ;;
esac

if [ -d testdata -o -e main.go ]; then
  # 間違ってさよならしないようにバックアップを取得する
  backup=.backup/program/`date "+%y%m/%d/%H%M%S"`
  mkdir -p $backup
  cp main.go ${backup}/main.go
  cp main_test.go ${backup}/main_test.go
  cp -r testdata ${backup}/testdata
  rm -rf testdata
  rm -f .current/problem
fi

oj d -d testdata -f "case%i/%e" $problem_url
cp template/main.go main.go
case $# in
  1)
    title=`get_current title`
    echo -n ${title}_${1} > .current/problem
    ;;
  2)
    echo -n ${1}_${2} > .current/problem
    ;;
esac

go run cmd/solve_test/main.go
