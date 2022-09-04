#!/bin/bash
base_dir=./

if [ $# == 1 ]; then
  base_dir=$1
fi

(cd $base_dir; oj t -D -c "go run main.go" -d testdata -f "%s/%e")
