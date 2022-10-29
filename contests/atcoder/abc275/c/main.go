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
	pList := make([][2]int, 0, 81)
	for i := 0; i < 9; i++ {
		line := r.Read()
		for j, c := range line {
			if c == '#' {
				pList = append(pList, [2]int{i, j})
			}
		}
	}

	var ans int
	for i := 0; i < len(pList); i++ {
		for j := i + 1; j < len(pList); j++ {
			for k := j + 1; k < len(pList); k++ {
				for l := k + 1; l < len(pList); l++ {
					p1, p2, p3, p4 := pList[i], pList[j], pList[k], pList[l]
					if innerProduct(p1, p2, p3) && innerProduct(p4, p2, p3) && isSquare(p1, p2, p3, p4) {
						ans++
					}
				}
			}
		}
	}
	w.WriteInt(ans)
}

func innerProduct(a, b, c [2]int) bool {
	ab := [2]int{b[0] - a[0], b[1] - a[1]}
	ac := [2]int{c[0] - a[0], c[1] - a[1]}
	ip := ab[0]*ac[0] + ab[1]*ac[1]
	return ip == 0
}

func isSquare(p1, p2, p3, p4 [2]int) bool {
	p12 := calcEdge(p1, p2)
	p13 := calcEdge(p1, p3)
	p42 := calcEdge(p4, p2)
	p43 := calcEdge(p4, p3)
	return p12 == p13 && p13 == p42 && p42 == p43
}

func calcEdge(a, b [2]int) int {
	ab := [2]int{b[0] - a[0], b[1] - a[1]}
	return ab[0]*ab[0] + ab[1]*ab[1]
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
