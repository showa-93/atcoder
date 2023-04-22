package main

import (
	"bufio"
	"container/list"
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
	n, m := reader.Int(), reader.Int()
	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := reader.Int()-1, reader.Int()-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// 全体点の距離をすべて計算する
	distance := new2dInt(n, n, 0)
	for u, edge := range graph {
		que := list.New()
		for _, v := range edge {
			que.PushBack(v)
			distance[u][v] = 1
		}
		for que.Len() > 0 {
			v := que.Remove(que.Front()).(int)
			d := distance[u][v]
			for _, w := range graph[v] {
				if u != w && (distance[u][w] > d+1 || distance[u][w] == 0) {
					que.PushBack(w)
					distance[u][w] = d + 1
				}
			}
		}
	}

	// 白でないと駄目な場所を塗る
	k := reader.Int()
	c := new1dInt(n, 1)
	pd := make([][2]int, k)
	black := make([][]int, n)
	for i := 0; i < k; i++ {
		p, d := reader.Int()-1, reader.Int()
		pd[i] = [2]int{p, d}
		for u, du := range distance[p] {
			// 必ず白の箇所をすべて塗る
			if du < d {
				c[u] = 0
			} else if du == d {
				// 黒じゃないと駄目な場所をちぇっくしとく
				black[p] = append(black[p], u)
			}
		}
	}

	// 黒じゃないとだめなばしょが黒のままかチェク
	for _, pdi := range pd {
		var exists bool
		for _, v := range black[pdi[0]] {
			if c[v] == 1 {
				exists = true
				break
			}
		}
		if !exists {
			writer.String("No")
			return
		}
	}

	writer.String("Yes")
	writer.Cr()
	for _, i := range c {
		writer.Int(i)
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

func (w *writer) String(s string) {
	w.w.WriteString(s)
	w.Space()
}

func (w *writer) Int(v int) {
	w.w.WriteString(strconv.Itoa(v))
}

func (w *writer) Space() {
	w.w.WriteString(" ")
}

func (w *writer) Cr() {
	w.w.WriteRune('\n')
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
