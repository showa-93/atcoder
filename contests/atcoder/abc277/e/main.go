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

	type node struct {
		vertex int
		s      bool
		edges  map[int]int
	}
	nodes := make(map[int]*node)
	n, m, k := r.ReadInt(), r.ReadInt(), r.ReadInt()
	for i := 0; i < m; i++ {
		u, v, a := r.ReadInt(), r.ReadInt(), r.ReadInt()+1
		{
			n, ok := nodes[u]
			if !ok {
				nodes[u] = &node{
					vertex: u,
					edges:  make(map[int]int),
				}
				n = nodes[u]
			}
			n.edges[v] = n.edges[v] | a
		}
		{
			n, ok := nodes[v]
			if !ok {
				nodes[v] = &node{
					vertex: v,
					edges:  make(map[int]int),
				}
				n = nodes[v]
			}
			n.edges[u] = n.edges[u] | a
		}
	}

	for i := 0; i < k; i++ {
		s := r.ReadInt()
		if n, ok := nodes[s]; ok {
			n.s = true
		}
	}

	if _, ok := nodes[n]; !ok {
		w.WriteInt(-1)
		return
	}

	ans := MaxInt
	que := list.New()
	que.PushBack([3]int{1, 0, 2})

	for que.Len() > 0 {
		v := que.Remove(que.Front()).([3]int)
		nn := nodes[v[0]]
		count := v[1] + 1
		s := v[2]
		for u, c := range nn.edges {
			if c&s != s {
				continue
			}
			nn.edges[u] -= s
			nodes[u].edges[nn.vertex] -= s
			if u == n {
				ans = Min(ans, count)
				break
			}
			que.PushBack([3]int{u, count, s})
		}
		if nn.s {
			s ^= 3
			for u, c := range nn.edges {
				if c&s != s {
					continue
				}
				nn.edges[u] -= s
				nodes[u].edges[nn.vertex] -= s
				if u == n {
					ans = Min(ans, count)
					break
				}
				que.PushBack([3]int{u, count, s})
			}
		}
	}

	if ans == MaxInt {
		ans = -1
	}

	w.WriteInt(ans)
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
