package main

import (
	"bufio"
	"container/list"
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
	n, m := r.ReadInt(), r.ReadInt()
	t := make([][]int, n)
	print := func() {
		for i := 0; i < n; i++ {
			w.w.WriteString(strconv.Itoa(t[i][0]))
			for j := 1; j < n; j++ {
				w.w.WriteString(" " + strconv.Itoa(t[i][j]))
			}
			w.w.WriteRune('\n')
		}
	}
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = -1
		}
	}
	t[0][0] = 0

	x, y := func() (int, int) {
		m2 := int(math.Sqrt(float64(m)))
		for i := m2; m2 > 0; i-- {
			for j := 0; j <= i; j++ {
				if m == i*i+j*j {
					return i, j
				}
			}
		}
		return -1, -1
	}()
	if x < 0 {
		print()
		return
	}

	que := list.New()
	que.PushBack([2]int{0, 0})

	for que.Len() > 0 {
		v := que.Remove(que.Front()).([2]int)
		for i := 0; i < 2; i++ {
			for _, i := range [][2]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}} {
				xx := v[0] + i[0]*x
				yy := v[1] + i[1]*y
				if n > xx && xx >= 0 && n > yy && yy >= 0 && t[xx][yy] < 0 {
					t[xx][yy] = t[v[0]][v[1]] + 1
					que.PushBack([2]int{xx, yy})
				}
			}
			if x == y {
				break
			}
			x, y = y, x
		}
		x, y = y, x
	}
	print()
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
