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
	n, s := r.ReadInt(), r.ReadInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = r.ReadInt(), r.ReadInt()
	}

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, s+1)
	}

	// 基準をつくる
	dp[0][0] = true

	// jをそれまでのカードでつくれる(true)場合、
	// j+表、j+裏の番号がつくれる数になる
	// これをS以下の範囲で調べる
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			if dp[i][j] {
				if j+a[i] <= s {
					dp[i+1][j+a[i]] = true
				}
				if j+b[i] <= s {
					dp[i+1][j+b[i]] = true
				}
			}
		}
	}

	// すべてのカードを使ってSを作れた場合、復元をおこなう
	if dp[n][s] {
		w.WriteString("Yes")
		ans := ""
		for i := n - 1; i >= 0; i-- {
			if a[i] <= s && dp[i][s-a[i]] {
				s -= a[i]
				ans = "H" + ans
			} else {
				s -= b[i]
				ans = "T" + ans
			}
		}
		w.w.WriteString(ans)

		return
	}

	w.WriteString("No")
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
