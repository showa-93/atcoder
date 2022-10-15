package main

import (
	"bufio"
	"fmt"
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
	ww := NewWriter(out)
	defer ww.Flush()
	h, w := r.ReadInt(), r.ReadInt()
	s := [2]int{r.ReadInt() - 1, r.ReadInt() - 1}

	n := r.ReadInt()
	// メモリ食い過ぎ注意
	hwalls := make(map[int][]int)
	wwalls := make(map[int][]int)
	for i := 0; i < n; i++ {
		wall := [2]int{r.ReadInt() - 1, r.ReadInt() - 1}
		hwalls[wall[0]] = append(hwalls[wall[0]], wall[1])
		wwalls[wall[1]] = append(wwalls[wall[1]], wall[0])
	}
	for _, walls := range hwalls {
		sort.Slice(walls, func(i, j int) bool { return walls[i] < walls[j] })
	}
	for _, walls := range wwalls {
		sort.Slice(walls, func(i, j int) bool { return walls[i] < walls[j] })
	}

	q := r.ReadInt()
	for i := 0; i < q; i++ {
		sh, sw := s[0], s[1]
		d, l := r.Read(), r.ReadInt()
		switch d {
		case "L":
			sw -= l
			walls := hwalls[sh]
			if len(walls) == 0 {
				break
			}

			// 線形探索のときは、二分探索を使おう！！
			// 現在地より１つ大きい壁を探す
			// その１つ手前の壁があればそれが壁になる
			// 見つからなかった場合（len(walls)）、一番うしろの壁が対象
			ll := LowerBound(walls, s[1])
			if ll-1 >= 0 {
				sw = Max(sw, walls[ll-1]+1)
			}
		case "R":
			sw += l
			walls := hwalls[sh]
			if len(walls) == 0 {
				break
			}

			// 現在地より１つ大きい壁を探す
			// それがそのまま壁になる
			// 見つからなかった場合は対象がない
			ll := UpperBound(walls, s[1])
			if ll < len(walls) {
				sw = Min(sw, walls[ll]-1)
			}
		case "U":
			sh -= l
			walls := wwalls[sw]
			if len(walls) == 0 {
				break
			}

			ll := LowerBound(walls, s[0])
			if ll-1 >= 0 {
				sh = Max(sh, walls[ll-1]+1)
			}
		case "D":
			sh += l
			walls := wwalls[sw]
			if len(walls) == 0 {
				break
			}

			ll := UpperBound(walls, s[0])
			if ll < len(walls) {
				sh = Min(sh, walls[ll]-1)
			}
		}

		if sw < 0 {
			sw = 0
		} else if sw >= w {
			sw = w - 1
		}

		if sh < 0 {
			sh = 0
		} else if sh >= h {
			sh = h - 1
		}

		s[0], s[1] = sh, sw
		ww.WriteString(fmt.Sprintf("%d %d", s[0]+1, s[1]+1))
	}
}

func LowerBound(list []int, value int) int {
	l, r := 0, len(list)-1
	for r-l >= 0 {
		c := (l + r) / 2
		if value <= list[c] {
			r = c - 1
		} else {
			l = c + 1
		}
	}

	return l
}

func UpperBound(list []int, value int) int {
	l, r := 0, len(list)-1
	for r-l >= 0 {
		c := (l + r) / 2
		if value < list[c] {
			r = c - 1
		} else {
			l = c + 1
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
