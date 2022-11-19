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
	solve2(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	ww := NewWriter(out)
	defer ww.Flush()
	H, W := r.ReadInt(), r.ReadInt()
	_ = r.ReadInt()
	h, w := r.ReadInt(), r.ReadInt()
	a := new2dInt(H, W, 0)
	nm := make(map[int][][2]int)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			v := r.ReadInt()
			a[i][j] = v
			if _, ok := nm[v]; !ok {
				nm[v] = [][2]int{{i, j}}
			} else {
				nm[v] = append(nm[v], [2]int{i, j})
			}
		}
	}

	table := make(map[int][4]int)
	for v, positions := range nm {
		hmax, hmin, wmax, wmin := MinInt, MaxInt, MinInt, MaxInt
		for _, p := range positions {
			hmax = Max(hmax, p[0])
			hmin = Min(hmin, p[0])
			wmax = Max(wmax, p[1])
			wmin = Min(wmin, p[1])
		}
		table[v] = [4]int{hmax, hmin, wmax, wmin}
	}

	for i := 0; i <= H-h; i++ {
		for j := 0; j <= W-w; j++ {
			var count int

			for _, p := range table {
				if !(i <= p[1] && p[0] < i+h && j <= p[3] && p[2] < j+w) {
					count++
				}
			}

			ww.Int(count)
		}
		ww.Cr()
	}
}

// 2次元累積和
func solve2(in io.Reader, out io.Writer) {
	r := NewReader(in)
	ww := NewWriter(out)
	defer ww.Flush()
	H, W := r.ReadInt(), r.ReadInt()
	n := r.ReadInt()
	h, w := r.ReadInt(), r.ReadInt()
	a := new2dInt(H, W, 0)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			a[i][j] = r.ReadInt() - 1
		}
	}

	table := make([][][]int, H)
	for i := 0; i < H; i++ {
		table[i] = make([][]int, W)
		for j := 0; j < W; j++ {
			table[i][j] = make([]int, n)
		}
	}
	for i := H - 1; i >= 0; i-- {
		for j := W - 1; j >= 0; j-- {
			for k := 0; k < n; k++ {
				var v int
				if i+1 < H {
					v += table[i+1][j][k]
				}
				if j+1 < W {
					v += table[i][j+1][k]
				}
				if i+1 < H && j+1 < W {
					v -= table[i+1][j+1][k]
				}
				table[i][j][k] = v
			}
			table[i][j][a[i][j]] += 1
		}
	}

	for i := 0; i < H-h+1; i++ {
		for j := 0; j < W-w+1; j++ {
			var count int

			for k := 0; k < n; k++ {
				mCount := table[i][j][k]
				if i+h < H {
					mCount -= table[i+h][j][k]
				}
				if j+w < W {
					mCount -= table[i][j+w][k]
				}
				if i+h < H && j+w < W {
					mCount += table[i+h][j+w][k]
				}
				if mCount != table[0][0][k] {
					count++
				}
			}

			ww.Int(count)
		}
		ww.Cr()
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
