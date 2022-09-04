#!/bin/bash
base_dir=./

while (( $# > 0 ))
do
  case $1 in
    -b) # テストの対象となるディレクトリをルートディレクトリから変更する
      if [[ -z "$2" ]] || [[ "$2" =~ ^-+ ]]; then
        echo "'option' requires an argument." 1>&2
        exit 1
      else
        base_dir="$2"
        shift
      fi
      ;;
    -e) # 許容する誤差を指定する
      if [[ -z "$2" ]] || [[ "$2" =~ ^-+ ]]; then
        echo "'option' requires an argument." 1>&2
        exit 1
      else
        measurement_error=" -e ${2}"
        shift
      fi
      ;;
    -*)
      echo "Illegal option -- '$(echo $1 | sed 's/^-*//')'." 1>&2
      exit 1.
  esac
  shift
done

(cd $base_dir; oj t -D -c "go run main.go" -d testdata -f "%s/%e"${measurement_error})
