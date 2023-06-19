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
	n, k, q := reader.Int(), reader.Int(), reader.Int()

	d := make(map[int]*Item)
	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		// initの前はindexをちゃんと設定すること
		d[i] = &Item{0, i, i}
		items[i] = d[i]
	}
	pq := NewPriorityQueue(items, Desc)
	heap.Init(pq)

	exists := make(map[int]struct{})
	bitems := make([]*Item, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(pq).(*Item)
		// initの前はindexをちゃんと設定すること
		item.index = i
		bitems[i] = item
		exists[item.position] = struct{}{}
	}
	bpq := NewPriorityQueue(bitems, Asc)
	heap.Init(bpq)

	var sum int
	for i := 0; i < q; i++ {
		x := reader.Int() - 1
		y := reader.Int()
		if _, ok := exists[x]; ok {
			item := d[x]
			sum += y - item.priority
			item.priority = y
			heap.Fix(bpq, item.index)
		} else {
			item := d[x]
			item.priority = y
			heap.Fix(pq, item.index)
		}

		bitem := heap.Pop(bpq).(*Item)
		sum -= bitem.priority
		heap.Push(pq, bitem)
		delete(exists, bitem.position)

		item := heap.Pop(pq).(*Item)
		sum += item.priority
		heap.Push(bpq, item)
		exists[item.position] = struct{}{}

		writer.Int(sum).Cr()
	}
}

type Item struct {
	priority int
	position int
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
