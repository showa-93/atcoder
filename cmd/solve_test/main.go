package main

import (
	"os"
	"text/template"
)

var testFileTemplate string = `package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/showa-93/atcoder/testhelper"
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
{{- end }}
func TestSolve_Random(t *testing.T) {
	in, err := os.Open("testdata/random/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/random/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func SolveSimple(in io.Reader, out io.Writer) {
	reader := NewReader(in)
	writer := NewWriter(out)
	defer writer.Flush()
	reader.Int()
}
{{ range $i, $num := . -}}
{{ printf "" }}
func TestSolveSimple_Case{{ $num }}(t *testing.T) {
	in, err := os.Open("testdata/case{{ $num }}/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	SolveSimple(in, buf)

	want, err := os.ReadFile("testdata/case{{ $num }}/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}
{{ printf "" }}
{{- end }}

func BenchmarkSolve_Random(b *testing.B) {
	tb := testhelper.NewRandomTestBuilder()

	for i := 0; i < b.N; i++ {
		input := tb.Build()
		out1 := bytes.NewBuffer(make([]byte, 0))
		SolveSimple(strings.NewReader(input), out1)
		result1, _ := io.ReadAll(out1)

		out2 := bytes.NewBuffer(make([]byte, 0))
		solve(strings.NewReader(input), out2)
		result2, _ := io.ReadAll(out2)
		if diff := cmp.Diff(string(result1), string(result2)); diff != "" {
			os.MkdirAll("testdata/random", 0777)
			os.WriteFile("testdata/random/in", []byte(input), 0644)
			os.WriteFile("testdata/random/out", result1, 0644)
			os.WriteFile("testdata/random_result.csv", result2, 0644)
			b.Fatal(diff)
		}
	}
}`

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
