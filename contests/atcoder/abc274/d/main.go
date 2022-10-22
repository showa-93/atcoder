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
	x, y := r.ReadInt(), r.ReadInt()
	a := r.ReadIntLine(n)

	var max int = 2e4 + 2
	dx := make([]bool, max)
	dy := make([]bool, max)
	dpbuf := make([]bool, max)

	dx[a[0]+1e4] = true
	dy[1e4] = true

	for i := 1; i < len(a); i++ {
		for j := 0; j < max; j++ {
			dpbuf[j] = false
		}
		if i%2 == 0 {
			for j := 0; j < max; j++ {
				if dx[j] {
					if sum := j + a[i]; sum < max {
						dpbuf[sum] = true
					}
					if sum := j - a[i]; 0 <= sum {
						dpbuf[sum] = true
					}
				}
			}
			copy(dx, dpbuf)
		} else {
			for j := 0; j < max; j++ {
				if dy[j] {
					if sum := j + a[i]; sum < max {
						dpbuf[sum] = true
					}
					if sum := j - a[i]; 0 <= sum {
						dpbuf[sum] = true
					}
				}
			}
			copy(dy, dpbuf)
		}
	}

	if dx[x+1e4] && dy[y+1e4] {
		w.WriteString("Yes")
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
