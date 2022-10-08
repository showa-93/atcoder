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
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = -1
		}
	}
	t[0][0] = 0
	sm := make(map[int]int)
	mm := int(math.Sqrt(float64(m)))

	for i := 0; i <= m; i++ {
		sm[i] = -1
		v := int(math.Sqrt(float64(i)))
		if i == v*v {
			sm[i] = v
		}
	}

	que := list.New()
	que.PushBack([2]int{1, 1})

	for que.Len() > 0 {
		p := que.Remove(que.Front()).([2]int)
		for k := 1; k <= mm+p[0] && k <= n; k++ {
			ll := m - (k-p[0])*(k-p[0])
			if v, ok := sm[ll]; ok && v >= 0 {
				if l := p[1] + v; n >= l && l > 0 && t[k-1][l-1] < 0 {
					t[k-1][l-1] = t[p[0]-1][p[1]-1] + 1
					que.PushBack([2]int{k, l})
				}
				if l := p[1] - v; n >= l && l > 0 && t[k-1][l-1] < 0 {
					t[k-1][l-1] = t[p[0]-1][p[1]-1] + 1
					que.PushBack([2]int{k, l})
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		w.w.WriteString(strconv.Itoa(t[i][0]))
		for j := 1; j < n; j++ {
			w.w.WriteString(" " + strconv.Itoa(t[i][j]))
		}
		w.w.WriteRune('\n')
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
