package main

import (
	"bufio"
	"io"
	"os"
	"sort"
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
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = r.ReadIntLine(n)
	}

	p := make([][][]int, n)
	q := make([][][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([][]int, n)
		q[i] = make([][]int, n)
	}
	p[0][0] = append(p[0][0], a[0][0])
	q[n-1][n-1] = append(q[n-1][n-1], a[n-1][n-1])

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n {
				continue
			}
			if i > 0 {
				for _, x := range p[i-1][j] {
					p[i][j] = append(p[i][j], x^a[i][j])
				}
			}
			if j > 0 {
				for _, x := range p[i][j-1] {
					p[i][j] = append(p[i][j], x^a[i][j])
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i+j < n-1 {
				continue
			}
			if i+1 < n {
				for _, x := range q[i+1][j] {
					q[i][j] = append(q[i][j], x^a[i][j])
				}
			}
			if j+1 < n {
				for _, x := range q[i][j+1] {
					q[i][j] = append(q[i][j], x^a[i][j])
				}
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		j := n - 1 - i
		sort.Slice(q[i][j], func(x, y int) bool { return q[i][j][x] < q[i][j][y] })
		for _, x := range p[i][j] {
			// a[i][j]とxのxorと同じ数値が存在する数を数える
			// このとき、長さを調べている
			// [1 2 2 2 3]で2を探すと、upperboundで4、lowerboudで1が返るので
			// 4-1で3個あることがわかる
			v := x ^ a[i][j]
			ans += UpperBound(q[i][j], v) - LowerBound(q[i][j], v)
		}
	}

	w.WriteInt(ans)
}

// ソート済みのスライスから指定された要素以上の値が現れる最初の位置を返す
func LowerBound(list []int, value int) int {
	l, r := 0, len(list)-1
	for r-l >= 0 {
		mid := (l + r) / 2
		if value <= list[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

// ソート済みのスライスから指定された要素より大きい値が現れる最初の位置を返す
// より大きなものがなければ、len(list)の値になる
func UpperBound(list []int, value int) int {
	l, r := 0, len(list)-1
	for r-l >= 0 {
		mid := (l + r) / 2
		if value < list[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
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
