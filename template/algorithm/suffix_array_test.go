package algorithm

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
		t.Run(test.input, func(t *testing.T) {
			got := NewSAIS([]byte(test.input))
			assert.Equal(t, test.want, got.sa)
		})
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
		got := BucketSort(test.input)
		assert.Equal(t, test.want, got[1:])
	}
}

// 検証結果
// メモリアロケート大量に発生しているが、それでもバケットソートよりは倍以上の性能はでている
// しかし、標準ライブラリと比べると半分以下の性能しか出ていない
// あと、標準ライブラリはメモリの使い回しがテクニカルすぎる
// TODO:仕組みはなんとなくわかったので、いつか作り直す。いつか
//
// > go test -benchmem -run=^$ -benchtime 2s -bench . github.com/showa-93/atcoder/template/structure
// goos: linux
// goarch: amd64
// pkg: github.com/showa-93/atcoder/template/structure
// cpu: Intel(R) Core(TM) i7-9700 CPU @ 3.00GHz
// BenchmarkSAIS-8                       22         104793993 ns/op        57184818 B/op     325552 allocs/op
// BenchmarkBucketSort-8                  8         265481594 ns/op        115157676 B/op     24324 allocs/op
// BenchmarkIndex-8                      49          43160793 ns/op         4005971 B/op          2 allocs/op
// PASS
// ok      github.com/showa-93/atcoder/template/structure  7.019s
func BenchmarkSAIS(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0x5a77a1))
	data := make([]byte, 1e6)
	for i := range data {
		data[i] = byte(r.Intn(255) + 1)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// NewSAIS([]byte("aababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcab"))
		NewSAIS(data)
	}
}

func BenchmarkBucketSort(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0x5a77a1))
	data := make([]byte, 1e6)
	for i := range data {
		data[i] = byte(r.Intn(255) + 1)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// BucketSort(string("aababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcab"))
		BucketSort(string((data)))
	}
}

func BenchmarkIndex(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(0x5a77a1))
	data := make([]byte, 1e6)
	for i := range data {
		data[i] = byte(r.Intn(255) + 1)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// suffixarray.New([]byte("aababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcabaababcabddabcab"))
		suffixarray.New(data)
	}
}
