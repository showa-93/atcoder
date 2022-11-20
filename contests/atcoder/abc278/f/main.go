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
	n := r.ReadInt()
	ss := make([][2]int, n)
	for i := 0; i < n; i++ {
		s := ([]rune)(r.Read())
		ss[i] = [2]int{int(s[0]), int(s[len(s)-1])}
	}

	var f func(pre, t int) bool
	memo := make(map[[2]int]bool)
	f = func(pre, t int) bool {
		if ret, ok := memo[[2]int{pre, t}]; ok {
			return ret
		}

		ret := false
		for i := 0; i < n; i++ {
			if t>>i&1 == 1 && (pre == -1 || ss[pre][1] == ss[i][0]) {
				ret = ret || !f(i, t^(1<<i))
			}
		}

		memo[[2]int{pre, t}] = ret
		return ret
	}

	if f(-1, 1<<n-1) {
		w.WriteString("First")
	} else {
		w.WriteString("Second")
	}
}

// DPで解法
func solve2(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	n := r.ReadInt()
	ss := make([][2]int, n)
	for i := 0; i < n; i++ {
		s := ([]rune)(r.Read())
		ss[i] = [2]int{int(s[0]), int(s[len(s)-1])}
	}

	dp := make([][]bool, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]bool, n)
	}

	for s := 1<<n - 1; s >= 0; s-- {
		for i := 0; i < n; i++ {
			// iがまだ未使用かチェック
			if s>>i&1 == 1 {
				continue
			}
			res := false
			// 次に選ぶ文字j
			for j := 0; j < n; j++ {
				// jがまだ未使用かチェック
				if i == j || s>>j&1 == 1 {
					continue
				}
				// しりとりが"i→j"で成立
				if ss[i][1] == ss[j][0] {
					// iが使用済みの集合でjを選んだときの結果の反対の結果を格納
					// 前回と逆の人物が勝つから不等号つける
					res = res || !dp[s|1<<i][j]
				}
			}
			dp[s][i] = res
		}
	}

	for i := 0; i < n; i++ {
		if !dp[0][i] {
			w.String("First")
			return
		}
	}

	w.String("Second")
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
