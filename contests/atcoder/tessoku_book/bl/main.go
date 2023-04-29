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
	n, m := reader.Int(), reader.Int()
	g := make([][][2]int, n)
	for i := 0; i < m; i++ {
		a, b, c := reader.Int()-1, reader.Int()-1, reader.Int()
		g[a] = append(g[a], [2]int{b, c})
		g[b] = append(g[b], [2]int{a, c})
	}

	// コストが小さいところから順に確定させていく
	que := NewPriorityQueue([]*Item{}, Asc)
	heap.Init(que)

	fixed := make([]bool, n)
	dists := new1dInt(n, MaxInt)
	dists[0] = 0
	heap.Push(que, &Item{priority: 0, key: 0})
	for que.Len() > 0 {
		item := heap.Pop(que).(*Item)
		if fixed[item.key] {
			continue
		}
		// キューから取り出された頂点を確定とする
		fixed[item.key] = true
		for _, vv := range g[item.key] {
			dist := dists[item.key] + vv[1]
			if dist < dists[vv[0]] {
				dists[vv[0]] = dist
				heap.Push(que, &Item{priority: dist, key: vv[0]})
			}
		}
	}

	for _, d := range dists {
		if d == MaxInt {
			writer.Int(-1).Cr()
		} else {
			writer.Int(d).Cr()
		}
	}
}

type Item struct {
	priority int
	key      int
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
