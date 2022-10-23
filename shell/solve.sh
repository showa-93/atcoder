#!/bin/bash
. shell/contests_function.sh

. shell/test.sh

if [ $? = 0 ]; then
  oj s `get_problem_url` main.go
fi
