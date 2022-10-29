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
	SetModValue(mod998244353)
	n, m, k := r.ReadInt(), r.ReadInt(), r.ReadInt()
	cur := make([]int, n+1)
	next := make([]int, n+1)
	p := ModDiv(1, m)

	cur[0] = 1
	for i := 1; i <= k; i++ {
		for j := 0; j <= n; j++ {
			if j == n {
				next[j] = ModAdd(next[j], cur[j])
				continue
			}
			for a := 1; a <= m; a++ {
				k := j + a
				if k > n {
					k = 2*n - k
				}
				next[k] = ModAdd(next[k], ModMul(cur[j], p))
			}
		}
		copy(cur, next)
		for i := 0; i <= n; i++ {
			next[i] = 0
		}
	}

	w.WriteInt(cur[n])
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

func new2dInt(m, n int) [][]int {
	table := make([][]int, m)
	for i := 0; i < m; i++ {
		table[i] = make([]int, n)
	}

	return table
}

func reset2dInt(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			table[i][j] = 0
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

const (
	mod998244353  int = 998244353
	mod1000000007 int = 1000000007
)

func SetModValue(v int) {
	mod = v
}

var mod int = mod1000000007

func Mod(a int) int {
	a %= mod
	if a < 0 {
		a += mod
	}
	return a
}

func ModAdd(a, b int) int {
	return Mod(a + b)
}

func ModSub(a, b int) int {
	return ModAdd(a, -b)
}

func ModMul(a, b int) int {
	return Mod(a * b)
}

func ModPow(a, b int) int {
	p := 1
	for b > 0 {
		fmt.Printf("%b %v\n", b, b&1)
		if b&1 == 1 {
			p = ModMul(p, a)
		}
		a = ModMul(a, a)
		b >>= 1
	}

	return p
}

// 非再帰拡張Euclidの互除法
func ModInv(a int) int {
	b := mod
	x, y := 1, 0

	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		x -= t * y
		x, y = y, x
	}

	return Mod(x)
}

func ModDiv(a, b int) int {
	return ModMul(Mod(a), ModInv(b))
}
