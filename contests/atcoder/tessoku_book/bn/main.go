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
	dsf := NewDijointSetsForests(n)

	for i := 0; i < q; i++ {
		query, u, v := reader.String(), reader.Int()-1, reader.Int()-1
		switch query {
		case "1":
			dsf.Unite(u, v)
		case "2":
			if dsf.Same(u, v) {
				writer.String("Yes").Cr()
			} else {
				writer.String("No").Cr()
			}
		}
	}
}

type DijointSetsForests struct {
	parent []int
	rank   []int
}

func NewDijointSetsForests(n int) *DijointSetsForests {
	dsf := &DijointSetsForests{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := 1; i < n; i++ {
		dsf.MakeSet(i)
	}

	return dsf
}

func (dsf *DijointSetsForests) MakeSet(x int) {
	dsf.rank[x] = 0
	dsf.parent[x] = x
}

// 引数の頂点が含まれる集合のルートを返す
func (dsf *DijointSetsForests) FindSet(x int) int {
	// 対象がルート出ない場合、親の探索をおこない、経路圧縮をおこなう
	// ルートを直接参照させる
	if x != dsf.parent[x] {
		dsf.parent[x] = dsf.FindSet(dsf.parent[x])
	}
	return dsf.parent[x]
}

func (dsf *DijointSetsForests) Same(x, y int) bool {
	return dsf.FindSet(x) == dsf.FindSet(y)
}

// ２つの頂点を結合する
func (dsf *DijointSetsForests) Unite(x, y int) {
	// ルートで結合するほうの比較をおこなう
	rootX, rootY := dsf.FindSet(x), dsf.FindSet(y)
	// 高さが高い方の木に低い木をつなげる
	if dsf.rank[rootX] > dsf.rank[rootY] {
		// 低い方のルートを変更する
		dsf.parent[rootY] = rootX
	} else {
		dsf.parent[rootX] = rootY
	}
	if dsf.rank[rootX] == dsf.rank[rootY] {
		// 同じ高さの場合、つなげると１つ高くなるので、加算する
		dsf.rank[rootY]++
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
