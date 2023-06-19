package main

import (
	"bytes"
	"io"
	"os"
	"sort"
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

// ランダムテスト実装
func SolveSimple(in io.Reader, out io.Writer) {
	reader := NewReader(in)
	writer := NewWriter(out)
	defer writer.Flush()
	n, k, q := reader.Int(), reader.Int(), reader.Int()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < q; i++ {
		x, y := reader.Int()-1, reader.Int()
		if x < 0 {
			print(x)
		}
		a[x] = y

		copy(b, a)
		sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

		ans := 0
		for _, v := range b[:k] {
			ans += v
		}
		writer.Int(ans).Cr()
	}
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

func BenchmarkSolve_Random(b *testing.B) {
	tb := testhelper.NewRandomTestBuilder()
	tb.AddIntKey("n", 1, 500)
	// b.AddIntKey("n", 1, 500_000)
	tb.AddIntKey("k", 1, "n")
	tb.AddIntKey("q", 1, 500)
	// b.AddIntKey("q", 1, 500_000)
	tb.AddIntKey("x", 1, "n")
	tb.AddIntKey("y", 0, 1_000)
	tb.AddBuildOrder(1, []string{"n", "k", "q"})
	tb.AddBuildOrder("q", []string{"x", "y"})

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
