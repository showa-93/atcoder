package main

import (
	"bufio"
	"fmt"
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
	n := r.ReadInt()

	uu, dd := 1, n
	for uu < dd {
		i := (uu + dd) / 2
		v := Q(w, r, uu, i, 1, n)
		if v < i-uu+1 {
			dd = i
		} else {
			uu = i + 1
		}
	}

	ll, rr := 1, n
	for ll < rr {
		i := (ll + rr) / 2
		v := Q(w, r, 1, n, ll, i)
		if v < i-ll+1 {
			rr = i
		} else {
			ll = i + 1
		}
	}

	w.WriteString(fmt.Sprintf("! %d %d", uu, ll))
}

func Q(w *writer, r *reader, a, b, c, d int) int {
	w.WriteString(fmt.Sprintf("? %d %d %d %d", a, b, c, d))
	w.Flush()
	v := r.ReadInt()
	if v < 0 {
		os.Exit(0)
	}
	return v
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
