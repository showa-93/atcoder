package main

import (
	"bufio"
	"io"
	"math"
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
	base := 1000
	grid := [2001][2001]int{}

	for i := 0; i < n; i++ {
		x, y := r.ReadInt(), r.ReadInt()
		grid[x+base][y+base] = 1
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = 0

		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				x, y := i+dx, j+dy
				if math.Abs(float64(dx-dy)) == 2 {
					continue
				}
				if 0 <= x && x <= 2000 && 0 <= y && y <= 2000 && grid[x][y] == 1 {
					dfs(x, y)
				}
			}
		}
	}

	count := 0
	for i := 0; i < 2001; i++ {
		for j := 0; j < 2001; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				count++
			}
		}
	}

	w.WriteInt(count)
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
