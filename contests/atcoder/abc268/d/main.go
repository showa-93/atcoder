package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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
	n, m := r.ReadInt(), r.ReadInt()
	s := make([]string, n)
	t := make(map[string]struct{}, m)
	for i := 0; i < n; i++ {
		s[i] = r.Read()
	}
	for i := 0; i < m; i++ {
		t[r.Read()] = struct{}{}
	}

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	// 最小の組み合わせは順列+_になる
	// のこりの組み合わせは必要な文字列をひいたあまりの_をどうくみあわせるかなので
	// 余りを求めて再帰処理で利用する
	amari := 16 - len(strings.Join(s, "_"))
	var ans string
	for {
		if ans = dfs(s[1:], t, s[0], amari); ans != "" {
			w.WriteString(ans)
			return
		}

		// すべての順列に対しておこなう
		if !NextPermutation(s) {
			w.WriteInt(-1)
			return
		}
	}
}

func dfs(s []string, t map[string]struct{}, ans string, amari int) string {
	if len(ans) > 16 {
		return ""
	}

	// すべて使って組み合わせを作れたら、存在チェックする
	if len(s) == 0 {
		if _, ok := t[ans]; len(ans) >= 3 && !ok {
			return ans
		}
		return ""
	}

	join := "_"
	for i := 0; i <= amari; i++ {
		if ans2 := dfs(s[1:], t, ans+join+s[0], amari-i); ans2 != "" {
			return ans2
		}
		join += "_"
	}

	return ""
}

// 引数の配列をもとに次の順列を生成する
// 最初の呼び出しで渡す引数の配列は降順にソート済みであること
func NextPermutation(list []string) bool {
	if len(list) < 2 {
		return false
	}

	var i int
	for i = len(list) - 2; i >= 0; i-- {
		if list[i] < list[i+1] {
			break
		}
	}

	if i < 0 {
		return false
	}

	var j int
	for j = len(list) - 1; j >= i; j-- {
		if list[i] < list[j] {
			break
		}
	}

	list[i], list[j] = list[j], list[i]

	for p, q := i+1, len(list)-1; p < q; p, q = p+1, q-1 {
		list[p], list[q] = list[q], list[p]
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
