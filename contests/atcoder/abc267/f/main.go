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
	edgeList := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edgeList[i] = r.ReadIntLine(2)
	}

	// わーシャルフロイドで解こうとするがおおきすぎてTLE
	distanceList := make([][]int, n)
	for i := 0; i < n; i++ {
		distanceList[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				distanceList[i][j] = MaxInt
			}
		}
	}

	for k := 0; k < len(edgeList); k++ {
		distanceList[edgeList[k][0]-1][edgeList[k][1]-1] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			// 経由点への経路が存在しない場合スキップする
			if distanceList[i][k] == MaxInt {
				continue
			}
			for j := 0; j < n; j++ {
				// 経由点への経路が存在しない場合スキップする
				if distanceList[k][j] == MaxInt {
					continue
				}

				// 最短の方で「始点→終点」への距離を更新する
				distanceList[i][j] = Min(distanceList[i][j], distanceList[i][k]+distanceList[k][j])
				distanceList[j][i] = distanceList[i][j]
			}
		}
	}

	q := r.ReadInt()
	for i := 0; i < q; i++ {
		query := r.ReadIntLine(2)
		d := -1
		for i, v := range distanceList[query[0]-1] {
			if v == query[1] {
				d = i + 1
				break
			}
		}
		w.WriteInt(d)
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
