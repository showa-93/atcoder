package main

import (
	"os"
	"text/template"
)

var testFileTemplate string = `package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)
{{ range $i, $num := . -}}
{{ printf "" }}
func TestSolve_Case{{ $num }}(t *testing.T) {
	in, err := os.Open("testdata/case{{ $num }}/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/case{{ $num }}/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}
{{ printf "" }}
{{- end }}`

func main() {
	tmpl, err := template.New("").Parse(testFileTemplate)
	if err != nil {
		panic(err)
	}

	de, err := os.ReadDir("testdata")
	if err != nil {
		panic(err)
	}

	numList := make([]int, 0, len(de))
	for i := range de {
		numList = append(numList, i+1)
	}

	f, err := os.Create("main_test.go")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(f, numList)
}
