package main

import (
	"bufio"
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
	abc := make([][3]int, m)
	for i := 0; i < m; i++ {
		abc[i] = [3]int{r.ReadInt(), r.ReadInt(), r.ReadInt()}
	}
	e := r.ReadIntLine(k)

	min := MaxInt
	var f func(e []int, nn, v int)
	f = func(e []int, nn, v int) {
		if min < v {
			return
		}
		if nn == 1 {
			min = Min(min, v)
			return
		}
		p := make([]int, m)
		for i := len(e) - 1; i >= 0; i-- {
			if p[e[i]-1] == 1 {
				continue
			}
			p[e[i]-1] = 1
			a := abc[e[i]-1]
			if a[1] == nn {
				f(e[:i], a[0], v+a[2])
			}
		}
	}

	f(e, n, 0)
	if min == MaxInt {
		w.WriteInt(-1)
	} else {
		w.WriteInt(min)
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
