package main

import (
	"bufio"
	"container/list"
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
	var (
		n, x, y  = r.ReadInt(), r.ReadInt() - 1, r.ReadInt() - 1
		edgeList = make([][]int, n)
		ansQue   = list.New()
	)

	for i := 0; i < n-1; i++ {
		u, v := r.ReadInt()-1, r.ReadInt()-1
		edgeList[u] = append(edgeList[u], v)
		edgeList[v] = append(edgeList[v], u)
	}

	var dfs func(v, p int) bool
	dfs = func(v, p int) bool {
		if v == x {
			ansQue.PushBack(v)
			return true
		}
		for _, u := range edgeList[v] {
			if u == p {
				continue
			}
			if dfs(u, v) {
				ansQue.PushBack(v)
				return true
			}
		}

		return false
	}

	dfs(y, -1)

	for ansQue.Len() > 0 {
		ans := strconv.Itoa(ansQue.Remove(ansQue.Front()).(int) + 1)
		if ansQue.Len() > 0 {
			ans += " "
		}
		w.w.WriteString(ans)
	}
	w.w.WriteRune('\n')
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
