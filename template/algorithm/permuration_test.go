package algorithm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextPermutation(t *testing.T) {
	input := []string{"1", "2", "3"}
	want := [][]string{
		{"1", "2", "3"},
		{"1", "3", "2"},
		{"2", "1", "3"},
		{"2", "3", "1"},
		{"3", "1", "2"},
		{"3", "2", "1"},
	}
	got := make([][]string, 0, len(want))
	for ok := true; ok; ok = NextPermutation(input) {
		dst := make([]string, len(input))
		copy(dst, input)
		got = append(got, dst)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}
