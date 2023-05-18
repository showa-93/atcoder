package main

import (
	"bufio"
	"io"
	"math/bits"
	"os"
	"strconv"
)

const BufferSize int = 1e6

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
	h, w, t := reader.Int(), reader.Int(), reader.Int()
	a := make([]string, h)
	for i := 0; i < h; i++ {
		a[i] = reader.String()
	}

	// スタート、ゴール、お菓子の位置をノードとして取得
	nodes := make([][2]int, 0)
	var start, goal int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			switch a[i][j] {
			case 'S':
				start = len(nodes)
				nodes = append(nodes, [2]int{i, j})
			case 'G':
				goal = len(nodes)
				nodes = append(nodes, [2]int{i, j})
			case 'o':
				nodes = append(nodes, [2]int{i, j})

			}
		}
	}

	bfs := func(a []string, x, y int) [][]int {
		dist := new2dInt(h, w, -1)
		dist[x][y] = 0
		q := [][2]int{{x, y}}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			for _, d := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				x2, y2 := u[0]+d[0], u[1]+d[1]
				if x2 < 0 || x2 >= h || y2 < 0 || y2 >= w {
					continue
				}
				if a[x2][y2] == '#' {
					continue
				}
				if dist[x2][y2] == -1 {
					dist[x2][y2] = dist[u[0]][u[1]] + 1
					q = append(q, [2]int{x2, y2})
				}
			}
		}

		return dist
	}

	// 各頂点からの他の頂点までの距離をBFSで求める
	n := len(nodes)
	g := new2dInt(n, n, 1000000000) // 最大値は境界値を超えるように！
	for i := 0; i < n; i++ {
		u := nodes[i]
		dist := bfs(a, u[0], u[1])
		for j := 0; j < n; j++ {
			v := nodes[j]
			if d := dist[v[0]][v[1]]; d != -1 {
				g[i][j] = d
			}
		}
	}

	dp := new2dInt(1<<n, n, 1000000000)
	dp[1<<start][start] = 0 // スタートを開始位置とするのでstartにビットを立てたところが初期値

	for i := 0; i < 1<<n; i++ {
		// １つ前の頂点がu
		for u := 0; u < n; u++ {
			// 次に遷移する頂点がv
			for v := 0; v < n; v++ {
				// 現時点のdpに記録済みの現在値とu→vに遷移するときの値の小さい方をdpに記録
				j := i | 1<<v
				dp[j][v] = Min(dp[j][v], dp[i][u]+g[u][v])
			}
		}
	}

	ans := -1
	for s := 0; s < (1 << n); s++ {
		// T回以内でいけた経路だけを対象として、ビットのたっている数が多いものを探す
		if dp[s][goal] <= t {
			ans = Max(ans, bits.OnesCount(uint(s))-2)
		}
	}

	writer.Int(ans)
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
