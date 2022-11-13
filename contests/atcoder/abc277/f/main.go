package main

import (
	"bufio"
	"container/list"
	"io"
	"os"
	"sort"
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
	h, w := r.ReadInt(), r.ReadInt()
	a := new2dInt(h, w, 0)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a[i][j] = r.ReadInt()
		}
	}

	// 行単位
	{
		rows := make([][2]int, h)
		for i := 0; i < h; i++ {
			rows[i] = [2]int{0, MaxInt}
			for j := 0; j < w; j++ {
				if v := a[i][j]; v > 0 {
					rows[i][0] = Max(rows[i][0], v)
					rows[i][1] = Min(rows[i][1], v)
				}
			}
		}

		sort.Slice(rows, func(i, j int) bool {
			if rows[i][0] == rows[j][0] {
				return rows[i][1] < rows[j][1]
			}
			return rows[i][0] < rows[j][0]
		})
		for i := 0; i < h-1; i++ {
			cur, next := rows[i], rows[i+1]
			if cur[0] > next[1] {
				ww.WriteString("No")
				return
			}
		}
	}

	// 列単位
	{
		n := w + h*w + 1
		id := w // グラフの簡略化のための追加の頂点番号。最大w + h*w
		g := make(map[int][]int, n)
		indegreeList := make([]int, n) // 入次数
		for i := 0; i < h; i++ {
			cols := make([][2]int, 0)
			for j := 0; j < w; j++ {
				if a[i][j] > 0 {
					cols = append(cols, [2]int{j, a[i][j]})
				}
			}

			sort.Slice(cols, func(i, j int) bool { return cols[i][1] < cols[j][1] })

			prev := -1
			for j, v := range cols {
				if j-1 >= 0 && v[1] != cols[j-1][1] {
					prev = id
					id++
				}
				if prev != -1 {
					// 最初の列以外の場合、間に頂点をおく
					g[prev] = append(g[prev], v[0])
					indegreeList[v[0]]++
				}

				g[v[0]] = append(g[v[0]], id)
				indegreeList[id]++
			}
			id++
		}

		que := list.New()
		for i, c := range indegreeList {
			if c == 0 {
				que.PushBack(i)
			}
		}

		// トポロジカルソート
		var ans int
		for que.Len() > 0 {
			v := que.Remove(que.Front()).(int)
			ans++
			for _, u := range g[v] {
				indegreeList[u]--
				if indegreeList[u] == 0 {
					que.PushBack(u)
				}
			}
		}

		// 頂点を通過した回数が頂点と異なる場合、閉路が存在するのでNo
		if ans != n {
			ww.WriteString("No")
		} else {
			ww.WriteString("Yes")
		}
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
