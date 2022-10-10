package structure

import (
	"index/suffixarray"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSAIS(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			"aababcabddabcab",
			[]int{0, 13, 1, 10, 3, 6, 14, 2, 11, 4, 7, 12, 5, 9, 8},
		},
		{
			"abracadabra",
			[]int{10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2},
		},
		{
			"mmiissiissiippii",
			[]int{15, 14, 10, 6, 2, 11, 7, 3, 1, 0, 13, 12, 9, 5, 8, 4},
		},
		{
			"mmiissiippiissii",
			[]int{15, 14, 6, 10, 2, 7, 11, 3, 1, 0, 9, 8, 13, 5, 12, 4},
		},
		{
			"abracadabra0AbRa4Cad14abra",
			[]int{11, 20, 16, 21, 12, 17, 14, 25, 10, 15, 22, 7, 0, 3, 18, 5, 13, 23, 8, 1,
				4, 19, 6, 24, 9, 2},
		},
		{
			"zazazazaz",
			[]int{7, 5, 3, 1, 8, 6, 4, 2, 0},
		},
	}

	for _, test := range tests {
		got := SAIS(test.input)
		assert.Equal(t, test.want, got)
	}
}

func TestBucketSort(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			"aababcabddabcab",
			[]int{0, 13, 1, 10, 3, 6, 14, 2, 11, 4, 7, 12, 5, 9, 8},
		},
		{
			"abracadabra",
			[]int{10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2},
		},
		{
			"mmiissiissiippii",
			[]int{15, 14, 10, 6, 2, 11, 7, 3, 1, 0, 13, 12, 9, 5, 8, 4},
		},
		{
			"mmiissiippiissii",
			[]int{15, 14, 6, 10, 2, 7, 11, 3, 1, 0, 9, 8, 13, 5, 12, 4},
		},
		{
			"abracadabra0AbRa4Cad14abra",
			[]int{11, 20, 16, 21, 12, 17, 14, 25, 10, 15, 22, 7, 0, 3, 18, 5, 13, 23, 8, 1,
				4, 19, 6, 24, 9, 2},
		},
		{
			"zazazazaz",
			[]int{7, 5, 3, 1, 8, 6, 4, 2, 0},
		},
	}

	for _, test := range tests {
		got := SAIS(test.input)
		assert.Equal(t, test.want, got)
	}
}

func BenchmarkSAIS(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0x5a77a1))
	data := make([]byte, 1e6)
	for i := range data {
		data[i] = byte(r.Intn(256)%26 + 97)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		SAIS(string(data))
	}
}

func BenchmarkBucketSort(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0x5a77a1))
	data := make([]byte, 1e6)
	for i := range data {
		data[i] = byte(r.Intn(256)%26 + 97)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		BucketSort(string(data))
	}
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0x5a77a1))
	data := make([]byte, 1e6)
	for i := range data {
		data[i] = byte(r.Intn(256)%26 + 97)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		suffixarray.New(data)
	}
}
