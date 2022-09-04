#!/bin/bash
. shell/contests_function.sh

problem_url=`get_problem_url`

problem_directory=`get_problem_directory`
rm -rf $problem_directory
mkdir -p $problem_directory

cp main.go ${problem_directory}/main.go
cp -r testdata ${problem_directory}/
echo ${problem_url} > ${problem_directory}/README.md

rm main.go
rm -rf testdata
