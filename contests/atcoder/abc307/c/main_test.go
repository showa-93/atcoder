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

func TestSolve_Case4(t *testing.T) {
	in, err := os.Open("testdata/case4/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/case4/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func TestSolve_Case5(t *testing.T) {
	in, err := os.Open("testdata/case5/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	solve(in, buf)

	want, err := os.ReadFile("testdata/case5/out")
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

func TestSolveSimple_Case4(t *testing.T) {
	in, err := os.Open("testdata/case4/in")
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0))

	SolveSimple(in, buf)

	want, err := os.ReadFile("testdata/case4/out")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(buf.String(), string(want)); diff != "" {
		t.Error(diff)
	}
}

func BenchmarkSolve_Random(b *testing.B) {
	tb := testhelper.NewRandomTestBuilder()
	tb.AddIntKey("Ha", 1, 11)
	tb.AddIntKey("Wa", 1, 11)
	tb.AddString("a", "Wa", []rune{'.', '#'})
	tb.AddIntKey("Hb", 1, 11)
	tb.AddIntKey("Wb", 1, 11)
	tb.AddString("b", "Wb", []rune{'.', '#'})
	tb.AddIntKey("Hx", 1, 11)
	tb.AddIntKey("Wx", 1, 11)
	tb.AddString("x", "Wx", []rune{'.', '#'})
	tb.AddBuildOrder(1, []string{"Ha", "Wa"})
	tb.AddBuildOrder("Ha", []string{"a"})
	tb.AddBuildOrder(1, []string{"Hb", "Wb"})
	tb.AddBuildOrder("Hb", []string{"b"})
	tb.AddBuildOrder(1, []string{"Hx", "Wx"})
	tb.AddBuildOrder("Hx", []string{"x"})

	for i := 0; i < b.N; i++ {
		input := tb.Build()
		out2 := bytes.NewBuffer(make([]byte, 0))
		os.WriteFile("testdata/random_input.csv", []byte(input), 0644)
		solve(strings.NewReader(input), out2)
	}
}
