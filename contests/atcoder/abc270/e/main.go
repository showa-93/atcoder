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
	n, k := r.ReadInt(), r.ReadInt()
	que := list.New()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = r.ReadInt()
		if a[i] > 0 {
			que.PushBack(i)
		}
	}

	f := func(que *list.List) *list.List {
		que2 := list.New()
		m := k / que.Len()
		if m == 0 {
			m = 1
		}

		for que.Len() > 0 {
			if k <= 0 {
				break
			}
			i := que.Remove(que.Front()).(int)
			if a[i]-m < 0 {
				k -= a[i]
				a[i] = 0
				continue
			}

			a[i] -= m
			k -= m
			if a[i] > 0 {
				que2.PushBack(i)
			}
		}

		return que2
	}

	for k > 0 && que.Len() > 0 {
		que = f(que)
	}

	for i := 0; i < n; i++ {
		w.w.WriteString(strconv.Itoa(a[i]) + " ")
	}
	w.w.WriteRune('\n')
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
