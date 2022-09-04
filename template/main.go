package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	NewReader(in)

	NewWriter(out)
}

type reader struct {
	s *bufio.Scanner
}

func NewReader(r io.Reader) *reader {
	s := bufio.NewScanner(r)
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
