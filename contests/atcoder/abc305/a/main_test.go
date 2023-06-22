package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/showa-93/atcoder/testhelper"
)

func TestSolve_Case1(t *testing.T) {
	in, err := os.Open("testdata/case1/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/case1/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func TestSolve_Case2(t *testing.T) {
	in, err := os.Open("testdata/case2/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/case2/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func TestSolve_Case3(t *testing.T) {
	in, err := os.Open("testdata/case3/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/case3/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

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

func TestSolveSimple_Case1(t *testing.T) {
	in, err := os.Open("testdata/case1/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	SolveSimple(in, buf)

	want, err := os.ReadFile("testdata/case1/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func TestSolveSimple_Case2(t *testing.T) {
	in, err := os.Open("testdata/case2/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	SolveSimple(in, buf)

	want, err := os.ReadFile("testdata/case2/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func TestSolveSimple_Case3(t *testing.T) {
	in, err := os.Open("testdata/case3/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	SolveSimple(in, buf)

	want, err := os.ReadFile("testdata/case3/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}


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
}