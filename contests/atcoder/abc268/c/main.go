package main

import (
	"bufio"
	"io"
	"os"
	"sort"
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

// とけませんでした。
// 0とn-1の取り扱いで詰んだ
func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()

	n := r.ReadInt()
	p := r.ReadIntLine(n)

	dp := make([]int, len(p)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = MaxInt
	}

	for i := 0; i < len(p); i++ {
		j, v := lowerBound(n, dp, p[i])
		dp[j] = v
	}

	var ans int
	for i := len(p) - 1; i >= 0; i-- {
		if dp[i] < MaxInt {
			ans = i + 1
			break
		}
	}

	w.WriteInt(ans)
}

func lowerBound(n int, dp []int, v int) (int, int) {
	vv := []int{
		v + 1,
		v,
		v - 1,
	}
	if vv[0] >= n {
		vv[0] = 0
	}
	if vv[2] < 0 {
		vv[2] = n - 1
	}
	sort.Slice(vv, func(i, j int) bool { return vv[i] < vv[j] })

	l, r := 0, len(dp)-1
	for r-l >= 0 {
		ok := false
		c := (l + r) / 2
		for i := 0; i < len(vv); i++ {
			if dp[c] < vv[i] {
				ok = true
				break
			}
		}

		if ok {
			l = c + 1
		} else {
			r = c - 1
		}
	}
	if l <= 0 {
		return l, vv[0]
	}

	var max int
	for i := 0; i < len(vv); i++ {
		if dp[l-1] < vv[i] {
			max = vv[i]
			break
		}
	}

	return l, max
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
