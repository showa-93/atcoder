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
	nodes := make(map[int][][2]int)
	n, m, k := r.ReadInt(), r.ReadInt(), r.ReadInt()
	for i := 0; i < m; i++ {
		u, v, a := r.ReadInt(), r.ReadInt(), r.ReadInt()
		// 無効な世界線を+nで表現する
		if a == 0 {
			u += n
			v += n
		}
		nodes[u] = append(nodes[u], [2]int{v, 1})
		nodes[v] = append(nodes[v], [2]int{u, 1})
	}

	for i := 0; i < k; i++ {
		s := r.ReadInt()
		// 世界線を切り替える
		nodes[s] = append(nodes[s], [2]int{s + n, 0})
		nodes[s+n] = append(nodes[s+n], [2]int{s, 0})
	}

	// 各頂点の最短を記録する
	dist := make([]int, 2*n+1)
	for i := 1; i < 2*n+1; i++ {
		dist[i] = MaxInt
	}
	dist[1] = 0

	que := list.New()
	que.PushBack(1)
	for que.Len() > 0 {
		v := que.Remove(que.Front()).(int)
		edges := nodes[v]
		for _, edge := range edges {
			u, c := edge[0], edge[1]
			if dist[u] > dist[v]+c {
				dist[u] = dist[v] + c
				// スイッチが押されたら、キューの最初におく
				if c == 0 {
					que.PushFront(u)
				} else {
					que.PushBack(u)
				}
			}
		}
	}

	ans := Min(dist[n], dist[n*2])
	if ans == MaxInt {
		w.WriteInt(-1)
	} else {
		w.WriteInt(ans)
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
