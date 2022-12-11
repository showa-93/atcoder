package main

import (
	"bufio"
	"container/heap"
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
	n, m, k := r.ReadInt(), r.ReadInt(), r.ReadInt()
	a := r.ReadIntLine(n)
	sum := 0

	// 和に含まれないm-k個のキュー
	// popすると小さい値が返却される
	mQueue := &PriorityQueue{
		count: m,
		items: make([]*Item, m),
		less:  Asc,
	}
	// 和に含まれる小さいk個のキュー
	// popすると和から削除される可能性のある大きい値が返却される
	kQueue := &PriorityQueue{
		count: 0,
		items: make([]*Item, 0),
		less:  Desc,
	}
	// 和に含まれる小さいk個のキューに含まれるインデクスを保持
	used := make([]bool, n)
	// 削除する値を保持するキュー
	dQueue := list.New()
	for i := 0; i < m; i++ {
		mQueue.items[i] = &Item{
			priority: a[i], // 値
			value:    i,    // aのindex
			index:    i,    // 優先度付きキューの中のindex
		}
		dQueue.PushBack(mQueue.items[i])
	}
	heap.Init(mQueue)

	// m個の値から小さい方からk個取得する
	for i := 0; i < k; i++ {
		item := heap.Pop(mQueue).(*Item)
		sum += item.priority
		used[item.value] = true
		heap.Push(kQueue, item)
	}
	w.Int(sum)

	for i := m; i < n; i++ {
		// １つ削除する値を取得する
		rItem := dQueue.Remove(dQueue.Front()).(*Item)
		if used[rItem.value] {
			// k個のキューに含まれる場合、合計から引く
			sum -= rItem.priority
			heap.Remove(kQueue, rItem.index)
		} else {
			heap.Remove(mQueue, rItem.index)
		}

		// i番目のアイテムを追加する
		item := &Item{
			priority: a[i],
			value:    i,
			index:    i,
		}
		heap.Push(mQueue, item)
		dQueue.PushBack(item)

		if kQueue.Len() < k {
			// k個の和のキューの個数がk個以下になった場合、m-k個のキューから追加する
			for kQueue.Len() < k {
				item := heap.Pop(mQueue).(*Item)
				used[item.value] = true
				sum += item.priority
				heap.Push(kQueue, item)
			}
		} else {
			// １つだけ新しく追加したので、
			// k個のキューの最大とm-k個のキューの最小を比較して
			// 小さい方をk個のキューに追加する
			mitem := heap.Pop(mQueue).(*Item)
			kitem := heap.Pop(kQueue).(*Item)
			if mitem.priority < kitem.priority {
				// 忘れずに合計と和に含まれる値を記録する
				sum -= kitem.priority
				sum += mitem.priority
				used[kitem.value] = false
				used[mitem.value] = true
				heap.Push(mQueue, kitem)
				heap.Push(kQueue, mitem)
			} else {
				heap.Push(mQueue, mitem)
				heap.Push(kQueue, kitem)
			}
		}

		w.Int(sum)
	}
}

type Item struct {
	priority int
	value    int
	index    int
}

type PriorityQueue struct {
	count int
	less  func([]*Item, int, int) bool
	items []*Item
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
