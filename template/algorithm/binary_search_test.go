package algorithm

import "testing"

func TestLowerBound(t *testing.T) {
	tests := []struct {
		input []int
		value int
		want  int
	}{
		{
			input: []int{1, 2, 3, 4, 5},
			value: 3,
			want:  2,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			value: 0,
			want:  0,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			value: 1,
			want:  0,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			value: 5,
			want:  4,
		},
		{
			input: []int{},
			value: 2,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := LowerBound(tt.input, tt.value)
			if got != tt.want {
				t.Error(got)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	tests := []struct {
		input []int
		value int
		want  int
	}{
		{
			input: []int{1, 2, 3, 4, 5},
			value: 3,
			want:  3,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			value: 0,
			want:  0,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			value: 1,
			want:  1,
		},
		{
			input: []int{1, 2, 3, 4, 5},
			value: 5,
			want:  5,
		},
		{
			input: []int{},
			value: 2,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := UpperBound(tt.input, tt.value)
			if got != tt.want {
				t.Error(got)
			}
		})
	}
}
