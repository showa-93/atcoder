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

	n := r.ReadInt()
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		u, v := r.ReadInt(), r.ReadInt()
		u -= 1
		v -= 1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// クエリを頂点ごとに集計しとく
	q := r.ReadInt()
	query := make([][][2]int, n)
	for i := 0; i < q; i++ {
		u := r.ReadInt() - 1
		query[u] = append(query[u], [2]int{i, r.ReadInt()})
	}

	right := bfs(n, 0, graph)
	left := bfs(n, right, graph)

	ans := make([]int, q)
	for i := 0; i < q; i++ {
		ans[i] = -1
	}

	// 木の直系のどちらか端から数えた時の深さを調べる
	// 対象の頂点から祖先にかえったときのパス上の頂点を調べるので
	// 深さ優先探索
	var dfs func(u int, p int)
	for _, root := range []int{left, right} {
		path := make([]int, 0)
		dfs = func(u int, p int) {
			// 対象の頂点にクエリが存在する場合、
			// pathの配列の深さから対象の頂点を計算する
			for _, qq := range query[u] {
				// クエリの深さよりパスが長ければパス上に存在
				if qq[1] <= len(path) {
					// 頂点から順にpathに詰まるので、
					// 調べたい深さdの位置は長さ-d
					ans[qq[0]] = path[len(path)-qq[1]] + 1 // この+1は0始まりにするためにひいてるから帳尻合わせ
				}
			}

			// 子で続行するので、pathにappend
			path = append(path, u)
			for _, v := range graph[u] {
				// 親と一致しないときだけ
				if v != p {
					dfs(v, u)
				}
			}
			// おわったので、取り出す
			path = path[:len(path)-1]
		}
		dfs(root, -1)
	}

	for _, a := range ans {
		w.WriteInt(a)
	}
}

func bfs(n, root int, graph [][]int) int {
	var (
		que           = list.New()
		d             = make([]int, n)
		vmax, max int = root, 0
	)
	que.PushFront(root)
	for que.Len() != 0 {
		u := que.Remove(que.Back()).(int)
		for _, v := range graph[u] {
			if v != root && d[v] == 0 {
				d[v] = d[u] + 1
				if max < d[v] {
					max = d[v]
					vmax = v
				}
				que.PushFront(v)
			}
		}
	}

	return vmax
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
