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
	SetModValue(mod998244353)
	a, b, c := Mint(r.ReadInt()).Mod(), Mint(r.ReadInt()).Mod(), Mint(r.ReadInt()).Mod()
	d, e, f := Mint(r.ReadInt()).Mod(), Mint(r.ReadInt()).Mod(), Mint(r.ReadInt()).Mod()
	x := a.Mul(b.Mul(c))
	y := d.Mul(e.Mul(f))
	w.WriteInt(int(x.Sub(y)))
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

var modValue Mint = Mint(mod1000000007)

func SetModValue(v int) {
	modValue = Mint(v)
}

type Mint int

func (m Mint) Mod() Mint {
	return m % modValue
}

func (m Mint) Inv() Mint {
	return m.Pow(Mint(modValue - 2))
}

func (m Mint) Add(a Mint) Mint {
	return Mint(m + a).Mod()
}

func (m Mint) Sub(a Mint) Mint {
	v := Mint(m - a).Mod()
	if v < 0 {
		v += modValue
	}

	return v
}

func (m Mint) Mul(a Mint) Mint {
	return Mint(m * a).Mod()
}

func (m Mint) Div(a Mint) Mint {
	return m.Mul(a.Inv())
}

func (m Mint) Pow(n Mint) Mint {
	var p Mint = 1
	base := m
	for n > 0 {
		// nの2進数表記ごとに1のくらいの時だけかける
		if n&1 == 1 {
			p = p.Mul(base)
		}
		base = base.Mul(base)
		n >>= 1
	}

	return p
}
