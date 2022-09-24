package main

import (
	"bufio"
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
	w := NewWriter(out)
	defer w.Flush()
	n, m := r.ReadInt(), r.ReadInt()
	x := r.ReadIntLine(n) // 空港コスト
	y := r.ReadIntLine(n) // 港コスト

	edgeList := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		edgeList = append(edgeList, []int{r.ReadInt() - 1, r.ReadInt() - 1, r.ReadInt()})
	}

	ans := MaxInt
	// 港、空港のありなしの組み合わせ(4通り)で考える
	// 最小シュタイナー木
	for xi := 0; xi < 2; xi++ {
		for yi := 0; yi < 2; yi++ {
			edgeList2 := make([][]int, len(edgeList))
			copy(edgeList2, edgeList)
			// グラフを変形して空港用のハブになる点を考える
			// ハブ用の点からの辺の重みで考える
			nn := n
			if xi == 1 {
				for i := 0; i < n; i++ {
					edgeList2 = append(edgeList2, []int{nn, i, x[i]})
				}
				nn++
			}
			// 同様に港用のハブになる点を考える
			if yi == 1 {
				for i := 0; i < n; i++ {
					edgeList2 = append(edgeList2, []int{nn, i, y[i]})
				}
				nn++
			}

			sort.Slice(edgeList2, func(i, j int) bool { return edgeList2[i][2] < edgeList2[j][2] })
			dset := NewDijointSetsForests(nn)
			var totalCost int
			nm := 0
			for _, e := range edgeList2 {
				if !dset.Same(e[0], e[1]) {
					totalCost += e[2]
					dset.Unite(e[0], e[1])
					nm++
				}
			}
			// すべての頂点が連結した木になっているか確認する
			if nm == nn-1 {
				ans = Min(ans, totalCost)
			}
		}
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
