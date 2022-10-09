#!/bin/bash
. shell/contests_function.sh

# 雑に引数があったら削除しない
if [ $# = 1 ]; then
  remove=false
else
  remove=true
fi

problem_url=`get_problem_url`

problem_directory=`get_problem_directory`
rm -rf $problem_directory
mkdir -p $problem_directory

cp main.go ${problem_directory}/main.go
cp main_test.go ${problem_directory}/main_test.go
cp -r testdata ${problem_directory}/
echo ${problem_url} > ${problem_directory}/README.md

git add ${problem_directory}
git commit -m "`get_current site` `get_current title` `get_current problem`"

if "${remove}"; then
  rm main.go
  rm main_test.go
  rm -rf testdata
fi
