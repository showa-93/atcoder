package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
)

const BufferSize int = 1e9

const eps = 1e-9

const (
	MinInt = -1 << (64 - 1)
	MaxInt = 1<<(64-1) - 1
)

func main() {
	solve(os.Stdin, os.Stdout)
}

// ２問とけない
func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	n, a := r.ReadInt(), float64(r.ReadInt())
	fish := make([][3]int, n)
	for i := 0; i < n; i++ {
		fish[i] = [3]int{r.ReadInt(), r.ReadInt(), r.ReadInt()}
	}

	type time struct {
		t float64
		w int
	}

	var ans int

	for i := 0; i < n; i++ {
		times := make([]time, 0, 2*n)
		for j := 0; j < n; j++ {
			xDiff := float64(fish[j][1] - fish[i][1])
			vDiff := float64(fish[j][2] - fish[i][2])
			if vDiff == 0 {
				if 0 <= xDiff && xDiff <= a {
					times = append(times, time{0, fish[j][0]})
				}
			} else {
				t1, t2 := -xDiff/vDiff, (a-xDiff)/vDiff
				if t1 > t2 {
					t1, t2 = t2, t1
				}
				times = append(times, time{t1, fish[j][0]})
				times = append(times, time{t2 + eps, -fish[j][0]})
			}
		}

		sort.Slice(times, func(i, j int) bool { return times[i].t < times[j].t })

		var sum int
		for _, t := range times {
			sum += t.w
			ans = Max(ans, sum)
		}
	}

	w.WriteInt(ans)
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
