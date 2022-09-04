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
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()

	n, m := r.ReadInt(), r.ReadInt()
	pq := make(PriorityQueue, n)
	itemMap := make(map[int]*Item)
	for i := 1; i <= n; i++ {
		itemMap[i] = &Item{
			value:    r.ReadInt(),
			priority: 0,
			index:    i - 1,
			vertex:   i,
		}
		pq[i-1] = itemMap[i]
	}

	for i := 0; i < m; i++ {
		u, v := r.ReadInt(), r.ReadInt()
		itemMap[u].vertexes = append(itemMap[u].vertexes, v)
		itemMap[u].priority += itemMap[v].value
		itemMap[v].vertexes = append(itemMap[v].vertexes, u)
		itemMap[v].priority += itemMap[u].value
	}

	heap.Init(&pq)

	var max int
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		delete(itemMap, item.vertex)
		max = Max(max, item.priority)

		for _, vv := range item.vertexes {
			item2, ok := itemMap[vv]
			if !ok {
				continue
			}
			item2.priority -= item.value
			pq.Update(item2)
		}
	}

	w.WriteInt(max)
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

type Item struct {
	value    int
	priority int
	index    int
	vertex   int
	vertexes []int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item) {
	heap.Fix(pq, item.index)
}
