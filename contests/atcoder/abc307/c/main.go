package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const BufferSize int = 1e6

const (
	MinInt = -1 << (64 - 1)
	MaxInt = 1<<(64-1) - 1
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(in io.Reader, out io.Writer) {
	reader := NewReader(in)
	writer := NewWriter(out)
	defer writer.Flush()
	// 簡単にするために黒点のある範囲に２次元配列を絞り込む
	area := func(ss []string) [][]rune {
		lw, rw := len(ss[0]), 0
		for _, s := range ss {
			i := strings.Index(s, "#")
			if i != -1 {
				lw = Min(lw, i)
			}
			j := strings.LastIndex(s, "#")
			if j != -1 {
				rw = Max(rw, j)
			}
		}
		th, bh := 0, len(ss)-1
		for i := 0; i < len(ss); i++ {
			j := strings.Index(ss[i], "#")
			if j != -1 {
				break
			}
			th = i + 1
		}
		for i := len(ss) - 1; i >= 0; i-- {
			j := strings.Index(ss[i], "#")
			if j != -1 {
				break
			}
			bh = i - 1
		}

		boards := make([][]rune, bh-th+1)
		for i := th; i <= bh; i++ {
			row := i - th
			boards[row] = make([]rune, rw-lw+1)
			for j := lw; j <= rw; j++ {
				boards[row][j-lw] = rune(ss[i][j])
			}
		}

		return boards
	}

	ha, _ := reader.Int(), reader.Int()
	aBoards := area(reader.StringList(ha))
	hb, _ := reader.Int(), reader.Int()
	bBoards := area(reader.StringList(hb))
	hx, _ := reader.Int(), reader.Int()
	xBoards := area(reader.StringList(hx))
	// マスの数に制限がないので、ちゃんとAとBの黒点がXの範囲内だよねチェックをする
	if len(xBoards) < len(aBoards) || len(xBoards[0]) < len(aBoards[0]) || len(xBoards) < len(bBoards) || len(xBoards[0]) < len(bBoards[0]) {
		writer.String("No")
		return
	}

	mathced := func(boards, boards2 [][]rune) bool {
		// AまたはBのどちらかがXの先頭行と一致するのでそこの探査だけで片方の位置が決まる
		for i := 0; i < len(xBoards[0])-len(boards[0])+1; i++ {
			tmp := new2dRune(len(xBoards), len(xBoards[0]), '.')
			matched := true
			for h := 0; h < len(boards) && matched; h++ {
				for w := 0; w < len(boards[h]) && matched; w++ {
					if boards[h][w] == '#' && xBoards[h][i+w] != boards[h][w] {
						matched = false
					}
					tmp[h][i+w] = boards[h][w]
				}
			}
			if matched {
				// もう一方のシートを全探索する
				for ih := 0; ih < len(xBoards)-len(boards2)+1; ih++ {
					for iw := 0; iw < len(xBoards[0])-len(boards2[0])+1; iw++ {
						matched = true
						tmp2 := new2dRune(len(xBoards), len(xBoards[0]), '.')
						for i := 0; i < len(tmp2); i++ {
							for j := 0; j < len(tmp2[i]); j++ {
								tmp2[i][j] = tmp[i][j]
							}
						}
						for h := 0; h < len(boards2); h++ {
							for w := 0; w < len(boards2[h]); w++ {
								if boards2[h][w] == '#' {
									tmp2[ih+h][iw+w] = boards2[h][w]
								}
							}
						}
						for i := 0; i < len(xBoards) && matched; i++ {
							for j := 0; j < len(xBoards[i]) && matched; j++ {
								if xBoards[i][j] != tmp2[i][j] {
									matched = false
								}
							}
						}
						if matched {
							return true
						}
					}
				}
			}
		}

		return false
	}

	if mathced(aBoards, bBoards) || mathced(bBoards, aBoards) {
		writer.String("Yes")
	} else {
		writer.String("No")
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

func (r *reader) String() string {
	r.s.Scan()
	return r.s.Text()
}

func (r *reader) Int() int {
	r.s.Scan()
	num, _ := strconv.Atoi(r.s.Text())

	return num
}

func (r *reader) Float64() float64 {
	r.s.Scan()
	num, _ := strconv.ParseFloat(r.s.Text(), 64)

	return num
}

func (r *reader) StringList(n int) []string {
	line := make([]string, n)
	for i := 0; i < n; i++ {
		line[i] = r.String()
	}
	return line
}

func (r *reader) IntList(n int) []int {
	line := make([]int, n)
	for i := 0; i < n; i++ {
		line[i] = r.Int()
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

func (w *writer) String(s string) *writer {
	w.w.WriteString(s)
	return w
}

func (w *writer) Int(v int) *writer {
	w.w.WriteString(strconv.Itoa(v))
	return w
}

func (w *writer) Float64(v float64) *writer {
	w.w.WriteString(strconv.FormatFloat(v, 'f', 10, 64))
	return w
}

func (w *writer) Space() *writer {
	w.w.WriteString(" ")
	return w
}

func (w *writer) Cr() *writer {
	w.w.WriteRune('\n')
	return w
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

func new2dRune(m, n int, v rune) [][]rune {
	table := make([][]rune, m)
	for i := 0; i < m; i++ {
		table[i] = make([]rune, n)
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

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 == 1 {
			p *= a
		}
		a *= a
		b >>= 1
	}

	return p
}

func Factorial(a int) int {
	v := 1
	for i := 2; i <= a; i++ {
		v *= i
	}
	return v
}

func Permutation(a, b int) int {
	sum := 1
	for i := a - b + 1; i <= a; i++ {
		sum *= i
	}

	return sum
}

func Combination(a, b int) int {
	return Permutation(a, b) / Factorial(b)
}
