package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		m    int
		list []int
	)
	r := NewReader()
	line := r.ReadIntLine()
	m = line[1]
	list = r.ReadIntLine()

	a := make([]int, len(list))
	a[0] = list[0]
	for i := 1; i < len(list); i++ {
		a[i] = a[i-1] + list[i]
	}

	sum := make([]int, len(list)-m+1)
	for i := 0; i < m; i++ {
		sum[0] += list[i] * (i + 1)
	}
	max := sum[0]

	for i := 1; i < len(list)-m+1; i++ {
		sum[i] = sum[i-1] + m*list[i+m-1] - a[i+m-2]
		if i-2 >= 0 {
			sum[i] += a[i-2]
		}
		max = Max(
			max,
			sum[i],
		)
	}

	w := NewWriter()
	defer w.Close()
	w.WriteInt(max)
}

type reader struct {
	s *bufio.Scanner
}

func NewReader() *reader {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 1e8), 1e8)
	return &reader{
		s: s,
	}
}

func (r *reader) ReadLine() []string {
	r.s.Scan()
	return strings.Split(r.s.Text(), " ")
}

func (r *reader) ReadIntLine() []int {
	r.s.Scan()
	line := strings.Split(r.s.Text(), " ")
	numLine := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		numLine[i], _ = strconv.Atoi(line[i])
	}

	return numLine
}

type writer struct {
	w *bufio.Writer
}

func NewWriter() *writer {
	return &writer{
		w: bufio.NewWriter(os.Stdout),
	}
}

func (w *writer) Close() error {
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
