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
	n, m := r.ReadInt(), r.ReadInt()
	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = r.ReadInt() - 1
	}

	var cur int
	zeroPositions := make([]int, m)
	for i := 0; i < m-1; i++ {
		if cur == a[i] {
			cur = a[i] + 1
		} else if cur == a[i]+1 {
			cur = a[i]
		}
		zeroPositions[i+1] = cur
	}

	stack := list.New()
	orderBotoms := make([]int, n)
	for i := 0; i < n; i++ {
		orderBotoms[i] = i
	}
	stack.PushFront(orderBotoms[zeroPositions[m-1]])
	for i := m - 1; i > 0; i-- {
		orderBotoms[a[i]], orderBotoms[a[i]+1] = orderBotoms[a[i]+1], orderBotoms[a[i]]
		stack.PushFront(orderBotoms[zeroPositions[i-1]])
	}

	for stack.Len() > 0 {
		p := stack.Remove(stack.Front()).(int)
		w.Int(p + 1)
		w.Cr()
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
