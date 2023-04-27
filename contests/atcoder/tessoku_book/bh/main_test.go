package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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
