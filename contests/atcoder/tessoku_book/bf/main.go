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
	reader := NewReader(in)
	writer := NewWriter(out)
	defer writer.Flush()
	n, q := reader.Int(), reader.Int()
	rmq := NewRMQ(n, 0, Max)
	for i := 0; i < q; i++ {
		switch reader.String() {
		case "1":
			pos, x := reader.Int()-1, reader.Int()
			rmq.Update(pos, x)
		case "2":
			l, r := reader.Int()-1, reader.Int()-1
			writer.Int(rmq.Query(l, r)).Cr()
		}
	}
}

type RMQ struct {
	n      int
	size   int
	dst    []int
	minmax func(int, int) int
}

func NewRMQ(n int, size int, minmax func(int, int) int) *RMQ {
	rmq := &RMQ{
		n:      1,
		size:   size,
		minmax: minmax,
	}

	// 完全２分木になるようにnの値を計算する
	for rmq.n < n {
		rmq.n *= 2
	}
	rmq.dst = make([]int, 2*rmq.n-1)
	for i := 0; i < len(rmq.dst); i++ {
		rmq.dst[i] = rmq.size
	}

	return rmq
}

func (rmq *RMQ) Update(k, a int) {
	// 葉にあたる値は後ろn個の値を更新する
	k += rmq.n - 1
	rmq.dst[k] = a
	// セグメント木の葉側から順番に頂点に向かって更新する
	for k > 0 {
		k = (k - 1) / 2
		rmq.dst[k] = rmq.minmax(rmq.dst[2*k+1], rmq.dst[2*k+2])
	}
}

func (rmq *RMQ) Query(a, b int) int {
	return rmq.query(a, b, 0, 0, rmq.n)
}

func (rmq *RMQ) query(a, b, k, l, r int) int {
	// a, bの完全に範囲外の場合、存在しない
	if r <= a || b <= l {
		return rmq.size
	}

	// 範囲内の場合現在地が最小なのでそれを返す
	if a <= l && r <= b {
		return rmq.dst[k]
	}

	// 左分木、右分木について探索する
	vl := rmq.query(a, b, k*2+1, l, (l+r)/2)
	vr := rmq.query(a, b, k*2+2, (l+r)/2, r)

	return rmq.minmax(vl, vr)
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

func (r *reader) String() string {
	r.s.Scan()
	return r.s.Text()
}

func (r *reader) Int() int {
	r.s.Scan()
	num, _ := strconv.Atoi(r.s.Text())

	return num
}

func (r *reader) StringList(n int) []string {
	line := make([]string, n)
	for i := 0; i < n; i++ {
		line[i] = r.String()
	}
	return line
}

func (r *reader) IntList(n int) []int {
	line := make([]int, n)
	for i := 0; i < n; i++ {
		line[i] = r.Int()
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

func (w *writer) String(s string) *writer {
	w.w.WriteString(s)
	return w
}

func (w *writer) Int(v int) *writer {
	w.w.WriteString(strconv.Itoa(v))
	return w
}

func (w *writer) Space() *writer {
	w.w.WriteString(" ")
	return w
}

func (w *writer) Cr() *writer {
	w.w.WriteRune('\n')
	return w
}

func new1dInt(n, v int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = v
	}

	return a
}

func reset1dInt(a []int, v int) {
	for i := 0; i < len(a); i++ {
		a[i] = v
	}
}

func new2dInt(m, n, v int) [][]int {
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
		for j := 0; j < n; j++ {
			table[i][j] = v
		}
	}

	return table
}

func reset2dInt(table [][]int, v int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			table[i][j] = v
		}
	}
}

func copy2dInt(dst, src [][]int) {
	for i := 0; i < len(src); i++ {
		copy(dst[i], src[i])
	}
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

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 == 1 {
			p *= a
		}
		a *= a
		b >>= 1
	}

	return p
}

func Factorial(a int) int {
	v := 1
	for i := 2; i <= a; i++ {
		v *= i
	}
	return v
}

func Permutation(a, b int) int {
	sum := 1
	for i := a - b + 1; i <= a; i++ {
		sum *= i
	}

	return sum
}

func Combination(a, b int) int {
	return Permutation(a, b) / Factorial(b)
}
