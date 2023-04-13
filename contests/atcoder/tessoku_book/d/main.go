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
	reader := NewReader(in)
	writer := NewWriter(out)
	defer writer.Flush()
	SetModValue(mod1000000007)
	n, r := reader.ReadInt(), reader.ReadInt()
	ans := 1
	for i := 2; i <= n; i++ {
		ans = ModMul(ans, i)
	}
	for i := 2; i <= r; i++ {
		ans = ModDiv(ans, i)
	}
	for i := 2; i <= n-r; i++ {
		ans = ModDiv(ans, i)
	}
	writer.Int(ans)
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

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 == 1 {
			p *= a
		}
		a *= a
		b >>= 1
	}

	return p
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
		if b&1 == 1 {
			p = ModMul(p, a)
		}
		a = ModMul(a, a)
		b >>= 1
	}

	return p
}

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
