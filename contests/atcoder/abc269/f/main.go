package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

const BufferSize int = 1e9

const modValue int = 998244353

const (
	MinInt = -1 << (64 - 1)
	MaxInt = 1<<(64-1) - 1
)

func main() {
	solve(os.Stdin, os.Stdout)
}

var n, m int

func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()

	n, m = r.ReadInt(), r.ReadInt()
	q := r.ReadInt()

	for i := 0; i < q; i++ {
		a, b, c, d := r.ReadInt(), r.ReadInt(), r.ReadInt(), r.ReadInt()
		a--
		c--

		var ans int
		ans += sum(b, d)
		ans -= sum(b, c)
		ans -= sum(a, d)
		ans += sum(a, c)

		w.WriteInt(Mod(ans))
	}
}

func sumHalf(x, y, init int) int {
	ans := Mod(init + Mod(x-1)*m + (y - 1))
	return Mod(Mod(ans*x) * y)
}

func sum(x, y int) int {
	ans := sumHalf((x+1)/2, (y+1)/2, 1)
	ans += sumHalf(x/2, y/2, m+2)

	return Mod(ans)
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

func Mod(x int) int {
	v := x % modValue
	if v < 0 {
		v += modValue
	}
	return v
}
