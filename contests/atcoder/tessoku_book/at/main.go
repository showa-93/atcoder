package main

import (
	"bufio"
	"io"
	"math"
	"math/rand"
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
	n := reader.Int()
	p := make([][2]int, n)
	for i := 0; i < n; i++ {
		p[i] = [2]int{reader.Int(), reader.Int()}
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, i)
	}
	ans = append(ans, 0)

	calcLength := func(a [2]int, b [2]int) float64 {
		return math.Sqrt(float64(Pow(a[0]-b[0], 2) + Pow(a[1]-b[1], 2)))
	}

	calcScore := func() float64 {
		var ret float64
		for i := 1; i < len(ans); i++ {
			ret += calcLength(p[ans[i-1]], p[ans[i]])
		}

		return ret
	}

	randInt := func(a, b int) int {
		return a + rand.Intn(b-a+1)
	}

	// 2-opt法
	score := calcScore()
	for t := 1; t <= 200000; t++ {
		// ランダムに反転する場所を探す
		l, r := randInt(1, n-1), randInt(1, n-1)
		if l > r {
			l, r = r, l
		}

		for i := 0; i <= (r-l)/2; i++ {
			ans[l+i], ans[r-i] = ans[r-i], ans[l+i]
		}

		newScore := calcScore()
		// 焼きなまし法による確率による変化を起こす
		var T float64 = 30.0 - 28.0*float64(t)/200000.0
		pro := math.Exp(math.Min(0.0, (score-newScore)/T))
		if rand.Float64() < pro {
			score = newScore
		} else {
			for i := 0; i <= (r-l)/2; i++ {
				ans[l+i], ans[r-i] = ans[r-i], ans[l+i]
			}
		}
	}

	for _, p := range ans {
		writer.Int(p + 1)
		writer.Cr()
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

func (r *reader) String() string {
	r.s.Scan()
	return r.s.Text()
}

func (r *reader) Int() int {
	r.s.Scan()
	num, _ := strconv.Atoi(r.s.Text())

	return num
}

func (r *reader) StringList(n int) []string {
	line := make([]string, n)
	for i := 0; i < n; i++ {
		line[i] = r.String()
	}
	return line
}

func (r *reader) IntList(n int) []int {
	line := make([]int, n)
	for i := 0; i < n; i++ {
		line[i] = r.Int()
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

func Factorial(a int) int {
	v := 1
	for i := 2; i <= a; i++ {
		v *= i
	}
	return v
}

func Permutation(a, b int) int {
	sum := 1
	for i := a - b + 1; i <= a; i++ {
		sum *= i
	}

	return sum
}

func Combination(a, b int) int {
	return Permutation(a, b) / Factorial(b)
}
