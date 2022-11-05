package algorithm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextPermutation(t *testing.T) {
	input := []int{1, 2, 3}
	want := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}
	got := make([][]int, 0, len(want))
	for ok := true; ok; ok = NextPermutation(input) {
		dst := make([]int, len(input))
		copy(dst, input)
		got = append(got, dst)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}

func TestPrevPermutation(t *testing.T) {
	input := []int{3, 2, 1}
	want := [][]int{
		{3, 2, 1},
		{3, 1, 2},
		{2, 3, 1},
		{2, 1, 3},
		{1, 3, 2},
		{1, 2, 3},
	}
	got := make([][]int, 0, len(want))
	for ok := true; ok; ok = PrevPermutation(input) {
		dst := make([]int, len(input))
		copy(dst, input)
		got = append(got, dst)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Error(diff)
	}
}
