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
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	N, m := r.ReadInt(), r.ReadInt()
	type node struct {
		n     int
		c     int8
		nodes []*node
	}
	mm := make(map[int]*node)
	if m == 0 {
		var ans int
		for i := 1; i < N; i++ {
			ans += i
		}
		w.Int(ans)
		return
	}

	var f int
	for i := 0; i < m; i++ {
		u, v := r.ReadInt(), r.ReadInt()
		f = u
		if _, ok := mm[u]; !ok {
			mm[u] = &node{n: u, nodes: make([]*node, 0)}
		}
		if _, ok := mm[v]; !ok {
			mm[v] = &node{n: v, nodes: make([]*node, 0)}
		}
		mm[u].nodes = append(mm[u].nodes, mm[v])
		mm[v].nodes = append(mm[v].nodes, mm[u])
	}

	black := 0
	que := list.New()
	que2 := list.New()
	mm[f].c = 1
	que2.PushBack(f)
	for _, n := range mm[f].nodes {
		que.PushBack(n.n)
		n.c = mm[f].c * -1
		black++
	}
	for que.Len() > 0 {
		n := que.Remove(que.Back()).(int)
		for _, nn := range mm[n].nodes {
			if nn.c != 0 {
				if mm[n].c == nn.c {
					w.Int(0)
					return
				}
			} else {
				nn.c = mm[n].c * -1
				if nn.c < 0 {
					black++
				} else {
					que2.PushBack(nn.n)
				}
				que.PushBack(nn.n)
			}
		}
	}

	for i := 1; i <= N; i++ {
		if _, ok := mm[i]; !ok {
			black++
		}
	}
	var ans int
	for i := 1; i <= N; i++ {
		if _, ok := mm[i]; !ok {
			ans += black - 1
		}
	}

	for que2.Len() > 0 {
		n := que2.Remove(que2.Back()).(int)
		ans += black - len(mm[n].nodes)
	}

	w.Int(ans)
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

func (w *writer) String(s string) {
	w.w.WriteString(s)
	w.Space()
}

func (w *writer) Int(v int) {
	w.w.WriteString(strconv.Itoa(v))
	w.Space()
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
