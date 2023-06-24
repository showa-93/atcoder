package main

import (
	"bufio"
	"container/heap"
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
	n, m, k := reader.Int(), reader.Int(), reader.Int()
	g := new2dInt(n, 0, 0)
	for i := 0; i < m; i++ {
		a, b := reader.Int()-1, reader.Int()-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	// ダイクストラの類問
	// 警備員が警備できる頂点＝体力がゼロ以上割り当てられている頂点を確定させる
	// 最短経路問題と異なり各ノードが最大になるように割り当てる
	// 計算量がおおいため、優先度付きキューをつかって探索量をO(n)→O(log n)に減らす
	pq := NewPriorityQueue([]*Item{}, Desc)
	heap.Init(pq)
	d := new1dInt(n, -1)
	for i := 0; i < k; i++ {
		p, h := reader.Int()-1, reader.Int()
		d[p] = h
		heap.Push(pq, &Item{index: i, value: p, priority: h})
	}

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		for _, v := range g[item.value] {
			if d[item.value]-1 > d[v] {
				d[v] = d[item.value] - 1
				heap.Push(pq, &Item{value: v, priority: d[v]})
			}
		}
	}

	var count int
	ans := make([]int, 0, n)
	for v, h := range d {
		if h >= 0 {
			count++
			ans = append(ans, v+1)
		}
	}
	writer.Int(count).Cr()
	for _, v := range ans {
		writer.Int(v).Space()
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

func (r *reader) Float64() float64 {
	r.s.Scan()
	num, _ := strconv.ParseFloat(r.s.Text(), 64)

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

func (w *writer) Float64(v float64) *writer {
	w.w.WriteString(strconv.FormatFloat(v, 'f', 10, 64))
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

type Item struct {
	priority int
	value    int
	index    int
}

type PriorityQueue struct {
	count int
	items []*Item
	less  func([]*Item, int, int) bool
}

func NewPriorityQueue(items []*Item, less func([]*Item, int, int) bool) *PriorityQueue {
	return &PriorityQueue{
		count: len(items),
		items: items,
		less:  less,
	}
}

func (pq PriorityQueue) Len() int { return pq.count }

func (pq PriorityQueue) Empty() bool { return pq.count == 0 }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.less(pq.items, i, j)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.count++
	n := len(pq.items)
	item := x.(*Item)
	item.index = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	pq.count--
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	pq.items = old[0 : n-1]
	return item
}

func Desc(items []*Item, i, j int) bool {
	return items[i].priority > items[j].priority
}

func Asc(items []*Item, i, j int) bool {
	return items[i].priority < items[j].priority
}
