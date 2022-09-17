package main

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
)

const BufferSize int = 1e9

const (
	MinInt = -1 << (64 - 1)
	MaxInt = 1<<(64-1) - 1
)

func main() {
	// solve(os.Stdin, os.Stdout)
	solve2(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	n := r.ReadInt()
	base := 1000
	grid := [2001][2001]int{}

	for i := 0; i < n; i++ {
		x, y := r.ReadInt(), r.ReadInt()
		grid[x+base][y+base] = 1
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = 0

		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				x, y := i+dx, j+dy
				if math.Abs(float64(dx-dy)) == 2 {
					continue
				}
				if 0 <= x && x <= 2000 && 0 <= y && y <= 2000 && grid[x][y] == 1 {
					dfs(x, y)
				}
			}
		}
	}

	count := 0
	for i := 0; i < 2001; i++ {
		for j := 0; j < 2001; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				count++
			}
		}
	}

	w.WriteInt(count)
}

// 各頂点をグラフにみなして、UnionFindで素な集合の数を数える
func solve2(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	n := r.ReadInt()

	pairs := make([][2]int, n)
	mp := make(map[[2]int]int) // 頂点の存在チェック用
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{r.ReadInt(), r.ReadInt()}
		mp[pairs[i]] = i
	}

	dsf := NewDejointSetsForests(n)
	check := func(n int, p [2]int) {
		if i, ok := mp[p]; ok {
			dsf.Unite(n, i)
		}
	}

	// 各ペアで片側の隣り合うペアが存在するかチェックし、
	// 存在する場合、そのペアをdejoint setsに追加する
	// 片側の座標で問題ないのは、もし隣り合うペアがもう片側にあった場合、
	// もう片側にあるペアのほうで検出できるため片側で問題ない
	for i, pair := range pairs {
		x, y := pair[0], pair[1]
		check(i, [2]int{x + 1, y})
		check(i, [2]int{x, y + 1})
		check(i, [2]int{x + 1, y + 1})
	}

	count := 0
	for i := 0; i < n; i++ {
		// dsfの頂点になっているペアの数を数えることで対象を求める
		if dsf.FindSet(i) == i {
			count++
		}
	}
	w.WriteInt(count)
}

// 素な集合
type DejointSetsForests struct {
	parent []int
	rank   []int
}

func NewDejointSetsForests(n int) *DejointSetsForests {
	dsf := &DejointSetsForests{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := 1; i < n; i++ {
		dsf.MakeSet(i)
	}

	return dsf
}

func (dsf *DejointSetsForests) MakeSet(x int) {
	dsf.rank[x] = 0
	dsf.parent[x] = x
}

// ルートを返す
func (dsf *DejointSetsForests) FindSet(x int) int {
	// 対象がルート出ない場合、親の探索をおこない、経路圧縮をおこなう
	// ルートを直接参照させる
	if x != dsf.parent[x] {
		dsf.parent[x] = dsf.FindSet(dsf.parent[x])
	}
	return dsf.parent[x]
}

func (dsf *DejointSetsForests) Same(x, y int) bool {
	return dsf.FindSet(x) == dsf.FindSet(y)
}

func (dsf *DejointSetsForests) Unite(x, y int) {
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
