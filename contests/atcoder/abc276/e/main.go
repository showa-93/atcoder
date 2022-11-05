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
	ww := NewWriter(out)
	defer ww.Flush()
	h, w := r.ReadInt(), r.ReadInt()
	var s []int
	board := make([][]rune, h)
	for i := 0; i < h; i++ {
		line := r.Read()
		board[i] = ([]rune)(line)
		if s == nil {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == 'S' {
					s = []int{i, j}
					break
				}
			}
		}
	}

	que := list.New()
	boardInt := new2dInt(h, w, -1)
	boardInt[s[0]][s[1]] = 0
	n := 1
	for _, p := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		x, y := s[0]+p[0], s[1]+p[1]
		if h > x && x >= 0 && w > y && y >= 0 && board[x][y] == '.' {
			que.PushBack([2]int{x, y})
			boardInt[x][y] = n
			n++
		}
	}

	for que.Len() > 0 {
		v := que.Remove(que.Front()).([2]int)
		num := boardInt[v[0]][v[1]]
		for _, p := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			x, y := v[0]+p[0], v[1]+p[1]
			if h > x && x >= 0 && w > y && y >= 0 && board[x][y] == '.' {
				if boardInt[x][y] == -1 {
					boardInt[x][y] = num
					que.PushBack([2]int{x, y})
				} else if boardInt[x][y] != num && boardInt[x][y] > 0 {
					ww.WriteString("Yes")
					return
				}
			}
		}
	}

	ww.WriteString("No")
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
