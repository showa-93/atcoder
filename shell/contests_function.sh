#!/bin/bash

message="コンテストが開始されていません。init.shでコンテストを開始してください。"

function get_current() {
  value=`cat .current/${1} 2> /dev/null`
  echo $value
}

function get_problem_url() {
  site=`get_current site`
  url=`get_current url`
  title=`get_current title`

  case $# in
    1)
      problem="${title}_${1}"
      ;;
    2)
      problem="${1}_${2}"
      ;;
    *)
      problem=`get_current problem`
      ;;
  esac

  if [ -z $site -o -z $url -o -z $title -o -z $problem ]; then
    echo $message
    exit 1
  fi

  if [ $site = atcoder ]; then
    problem_url="${url}/tasks/${problem}"
  else
    echo 未定義のサイトが指定されました。
    exit 1
  fi

  echo ${problem_url}
}

function get_problem_directory() {
  site=`get_current site`
  problem=`get_current problem`
  title=$(echo "$problem" | sed -E 's/^(.*)_([a-z0-9]){1,2}$/\1/')
  problem_num=$(echo "$problem" | sed -E 's/^(.*)_([a-z0-9]{1,2})$/\2/')

  if [ -z $site -o -z $title -o -z $problem ]; then
    echo $message
    exit 1
  fi

  echo "contests/${site}/${title}/${problem_num}"
}
