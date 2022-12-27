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
	ww := NewWriter(out)
	defer ww.Flush()
	n, m := r.ReadInt(), r.ReadInt()
	color := make(map[int]int)
	graph := make(map[int][]int)
	for i := 0; i < m; i++ {
		u, v := r.ReadInt(), r.ReadInt()
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs func(u, p int) (int, int)
	// ある頂点が含まれる連結成分について白と黒の数をそれぞれ数える
	dfs = func(u, p int) (black int, white int) {
		if c, ok := color[p]; ok {
			color[u] = -c
		} else {
			color[u] = 1
		}
		if color[u] == 1 {
			black++
		} else {
			white++
		}

		for _, v := range graph[u] {
			if v == p || color[v] == -color[u] {
				continue
			}
			if color[v] == color[u] {
				return -1, -1
			}
			b, w := dfs(v, u)
			if b == -1 {
				return -1, -1
			}
			black += b
			white += w
		}

		return
	}

	ans := n*(n-1)/2 - m
	for u := 1; u <= n; u++ {
		// 塗ったことない頂点について調べる
		if _, ok := color[u]; !ok {
			color[u] = 1
			b, w := dfs(u, -1)
			if b == -1 {
				ww.Int(0)
				return
			}

			// 結合部分でつくれる辺をを除いた辺の数が答えなので
			// 白同士、黒同士の連結部分でつくれる辺を答えから引く
			// ２部グラフである＝白同士、黒同士の辺はまだ存在しない
			ans -= b * (b - 1) / 2
			ans -= w * (w - 1) / 2
		}
	}

	ww.Int(ans)
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
