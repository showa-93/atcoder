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
	n, k := r.ReadInt(), r.ReadInt()
	a := make([]int, k)
	for i := 0; i < k; i++ {
		a[i] = r.ReadInt()
	}

	dp := [2][]int{
		make([]int, n+1), // 先手から見て先手が取れる最大個数
		make([]int, n+1), // 後手から見て先手が取れる最小個数
	}
	for i := 1; i <= n; i++ {
		now := 0
		for _, ai := range a {
			if i-ai < 0 {
				break
			}
			{ // dp[1][i - a[i]](i - a[i]個選んだ場合の最大個数) + a[i]
				now = Max(now, dp[1][i-ai]+ai)
			}
			dp[0][i] = now
		}
		{
			now := MaxInt
			for _, ai := range a {
				if i-ai < 0 {
					break
				}
				now = Min(now, dp[0][i-ai])
			}
			dp[1][i] = now
		}
	}

	w.WriteInt(dp[0][n])
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
