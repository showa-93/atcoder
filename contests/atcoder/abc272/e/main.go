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
	n, m := r.ReadInt(), r.ReadInt()
	a := r.ReadIntLine(n)
	a = append([]int{0}, a...)
	results := make([][]int, m+1)
	for i := 1; i < n+1; i++ {
		// 最小の非負整数をNにするには、0~N-1まで埋めル必要がある
		// nより大きいものが１つある場合は、0~N-1の中で埋まらない数値がでるため、それが非負整数になる
		// よってN以上の整数は無視できる
		if a[i] >= n {
			continue
		}

		// i番目のaが正になるMの回数
		l := 1
		if a[i] < 0 {
			l = (a[i] * (-1)) / i
		}
		// i番目のaがN以上になるMの回数
		r := Min(m+1, (m-a[i]+i)/i)

		for j := l; j < r; j++ {
			results[j] = append(results[j], a[i]+i*j)
		}
	}

	for i := 1; i < m+1; i++ {
		size := len(results[i])
		m := make(map[int]struct{}, size)
		for _, r := range results[i] {
			m[r] = struct{}{}
		}
		res := 0
		for _, ok := m[res]; ok; _, ok = m[res] {
			res++
		}
		w.WriteInt(res)
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
