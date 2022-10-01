package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const BufferSize int = 1e9

const (
	MinInt = -1 << (64 - 1)
	MaxInt = 1<<(64-1) - 1
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	n, m, k := r.ReadInt(), r.ReadInt(), r.ReadInt()
	a, b, c := make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = r.ReadInt() - 1
		b[i] = r.ReadInt() - 1
		c[i] = r.ReadInt()
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = MaxInt
	}
	// 最初の街１を０で初期化
	// こうすることで開始位置と一致する道から始まる
	dp[0] = 0

	for i := 0; i < k; i++ {
		e := r.ReadInt() - 1
		// すでに訪れた場合とeの経路とそれまでのコストの和と比較して
		// 小さい方をセットする
		if dp[a[e]] < MaxInt && dp[a[e]]+c[e] < dp[b[e]] {
			dp[b[e]] = dp[a[e]] + c[e]
		}
	}

	if dp[n-1] < MaxInt {
		w.WriteInt(dp[n-1])
	} else {
		w.WriteInt(-1)
	}
}

type reader struct {
	s *bufio.Scanner
}

func NewReader(r io.Reader) *reader {
	s := bufio.NewScanner(r)
	s.Buffer(make([]byte, BufferSize), BufferSize)
	s.Split(bufio.ScanWords)
	return &reader{
		s: s,
	}
}

func (r *reader) Read() string {
	r.s.Scan()
	return r.s.Text()
}

func (r *reader) ReadInt() int {
	r.s.Scan()
	num, _ := strconv.Atoi(r.s.Text())

	return num
}

func (r *reader) ReadLine(n int) []string {
	line := make([]string, n)
	for i := 0; i < n; i++ {
		line[i] = r.Read()
	}
	return line
}

func (r *reader) ReadIntLine(n int) []int {
	line := make([]int, n)
	for i := 0; i < n; i++ {
		line[i] = r.ReadInt()
	}
	return line
}

type writer struct {
	w *bufio.Writer
}

func NewWriter(w io.Writer) *writer {
	return &writer{
		w: bufio.NewWriter(w),
	}
}

func (w *writer) Flush() error {
	return w.w.Flush()
}

func (w *writer) WriteString(s string) {
	w.w.WriteString(s)
	w.w.WriteRune('\n')
}

func (w *writer) WriteInt(v int) {
	w.w.WriteString(strconv.Itoa(v))
	w.w.WriteRune('\n')
}

func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
