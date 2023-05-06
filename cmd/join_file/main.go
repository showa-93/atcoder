package main

import (
	"errors"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"

	"golang.org/x/tools/imports"
)

const mainFile = "main.go"

var (
	targetPaths = map[string]string{
		"mod":   "template/algorithm/mod.go",
		"bs":    "template/algorithm/binary_search.go",
		"prime": "template/algorithm/prime.go",
		"pq":    "template/algorithm/priority_queue.go",
		"next":  "template/algorithm/next.go",
		"flow":  "template/algorithm/maxflow.go",
		"per":   "template/algorithm/permuration.go",
		"ds":    "template/algorithm/dijoint_sets.go",
		"gcd":   "template/algorithm/gcd.go",
		"rmq":   "template/algorithm/segmenttree/rmq.go",
		"rsq":   "template/algorithm/segmenttree/rsq.go",
		"bit":   "template/algorithm/segmenttree/bit.go",
	}
)

func main() {
	flag.Parse()
	targets := flag.Args()
	if len(targets) == 0 {
		fmt.Println("引数にmain.goに追加するターゲットを指定してください。listで一覧を表示します。")
		os.Exit(1)
	}
	if targets[0] == "list" {
		fmt.Println("指定可能なターゲットの一覧")
		for target, path := range targetPaths {
			fmt.Println(target, "=", path)
		}
		return
	}

	fileMap, err := readFiles(append([]string{mainFile}, targets...))
	if err != nil {
		fmt.Printf("ファイル読み込みでエラーが発生しました。 %v\n", err)
		os.Exit(1)
	}

	for _, target := range targets {
		fileMap[mainFile], err = joinFile(fileMap[mainFile], fileMap[target])
		if err != nil {
			fmt.Printf("ファイル結合でエラーが発生しました。 %v\n", err)
			os.Exit(1)
		}
	}

	os.WriteFile("main.go", fileMap[mainFile], fs.FileMode(0644))
}

func readFiles(targets []string) (map[string][]byte, error) {
	m := make(map[string][]byte)
	for _, target := range targets {
		path, ok := targetPaths[target]
		if !ok {
			path = target
		}
		file, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		m[target] = file
	}

	return m, nil
}

func joinFile(target, src []byte) ([]byte, error) {
	fset := token.NewFileSet()
	srcAst, err := parser.ParseFile(fset, "src", src, parser.Mode(parser.ParseComments))
	if err != nil {
		return nil, err
	}

	targetAst, err := parser.ParseFile(fset, "target", target, parser.Mode(parser.ParseComments))
	if err != nil {
		return nil, err
	}

	var objectKeys []string
	for objectKey := range srcAst.Scope.Objects {
		if _, ok := targetAst.Scope.Objects[objectKey]; ok {
			objectKeys = append(objectKeys, objectKey)
		}
	}
	if len(objectKeys) > 0 {
		return nil, fmt.Errorf("すでに関数が存在します。 %v", objectKeys)
	}

	if len(srcAst.Decls) == 0 {
		return nil, errors.New("joinできる構文が存在しません")
	}

	var pos int = int(srcAst.Decls[0].Pos())
	if len(srcAst.Imports) > 0 {
		if len(srcAst.Decls) == 1 {
			return nil, errors.New("joinできる構文が存在しません")
		}
		pos = int(srcAst.Decls[1].Pos())
	}

	target = append(target, '\n')
	target = append(target, src[pos-1:]...)

	target, err = imports.Process(mainFile, target, &imports.Options{})
	if err != nil {
		return nil, err
	}

	return target, nil
}
