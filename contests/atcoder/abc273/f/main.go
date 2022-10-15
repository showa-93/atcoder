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

const (
	Ham = iota
	Wall
	Goal
	Start
)

type Event struct {
	t   int
	num int
	pos int
}

func solve(in io.Reader, out io.Writer) {
	r := NewReader(in)
	w := NewWriter(out)
	defer w.Flush()
	n, x := r.ReadInt(), r.ReadInt()
	ll, rr := make([]Event, 0, n/2), make([]Event, 0, n/2)
	se := Event{
		t:   Start,
		pos: 0,
	}
	ll = append(ll, se)
	rr = append(rr, se)
	ge := Event{
		t:   Goal,
		pos: x,
	}
	isRight := true
	if ge.pos > 0 {
		rr = append(rr, ge)
	} else {
		isRight = false
		ge.pos *= -1
		ll = append(ll, ge)
	}
	for i := 0; i < n; i++ {
		e := Event{
			t:   Wall,
			num: i,
			pos: r.ReadInt(),
		}
		if e.pos > 0 {
			if !(isRight && ge.pos < e.pos) {
				rr = append(rr, e)
			}
		} else {
			e.pos *= -1
			if !(!isRight && ge.pos < e.pos) {
				ll = append(ll, e)
			}
		}
	}
	hm := make(map[int]int)
	for i := 0; i < n; i++ {
		e := Event{
			t:   Ham,
			num: i,
			pos: r.ReadInt(),
		}
		hm[e.num] = e.pos
		if e.pos > 0 {
			if !(isRight && ge.pos < e.pos) {
				rr = append(rr, e)
			}
		} else {
			e.pos *= -1
			if !(!isRight && ge.pos < e.pos) {
				ll = append(ll, e)
			}
		}
	}

	sort.Slice(ll, func(i, j int) bool { return ll[i].pos < ll[j].pos })
	sort.Slice(rr, func(i, j int) bool { return rr[i].pos < rr[j].pos })

	dp := make([][][2]int, len(ll))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, len(rr))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j][0] = MaxInt
			dp[i][j][1] = MaxInt
		}
	}
	dp[0][0][0] = 0
	dp[0][0][1] = 0

	for i := 0; i < len(ll); i++ {
		for j := 0; j < len(rr); j++ {
			if i == 0 && j == 0 {
				continue
			}

			// R
			if j > 0 {
				if canOpen(rr[j], hm, ll[i].pos, rr[j].pos) {
					if dp[i][j-1][0] < MaxInt {
						dp[i][j][1] = Min(dp[i][j][1], dp[i][j-1][0]+rr[j].pos+ll[i].pos)
					}
					if dp[i][j-1][1] < MaxInt {
						dp[i][j][1] = Min(dp[i][j][1], dp[i][j-1][1]+(rr[j].pos-rr[j-1].pos))
					}
				}
			}
			// L
			if i > 0 {
				if canOpen(ll[i], hm, ll[i].pos, rr[j].pos) {
					if dp[i-1][j][0] < MaxInt {
						dp[i][j][0] = Min(dp[i][j][0], dp[i-1][j][0]+(ll[i].pos-ll[i-1].pos))
					}
					if dp[i-1][j][1] < MaxInt {
						dp[i][j][0] = Min(dp[i][j][0], dp[i-1][j][1]+rr[j].pos+ll[i].pos)
					}
				}
			}

			if ll[i].t == Goal || rr[j].t == Goal {
				ans := Min(dp[i][j][0], dp[i][j][1])
				if ans == MaxInt {
					continue
				}
				w.WriteInt(ans)
				return
			}
		}
	}
	w.WriteInt(-1)
}

func canOpen(target Event, hm map[int]int, l, r int) bool {
	if target.t == Wall {
		pos := hm[target.num]
		if l*-1 <= pos && pos <= r {
			return true
		}
		return false
	}

	return true
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
