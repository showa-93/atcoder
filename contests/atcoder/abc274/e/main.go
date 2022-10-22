package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/bits"
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
	n, m := r.ReadInt(), r.ReadInt()
	a := make([][2]float64, n+m)
	for i := 0; i < n+m; i++ {
		a[i] = [2]float64{float64(r.ReadInt()), float64(r.ReadInt())}
	}

	dp := make([][]float64, 1<<(n+m))
	for i := 0; i < 1<<(n+m); i++ {
		dp[i] = make([]float64, n+m)
		for j := 0; j < n+m; j++ {
			dp[i][j] = MaxInt
		}
	}

	// 原点からある頂点に行くときの時間で初期化
	for i := 0; i < n+m; i++ {
		dp[1<<i][i] = calcDistance(a[i], [2]float64{0, 0})
	}

	for s := 1; s < 1<<(n+m); s++ {
		// [n,m]の順番でつまってるので、
		// mmmmnnnnnnnのようなbitのつまり方をしているため、
		// nだけ右にシフトしてmだけのビット列にする
		// その中でbitが立っている数がスピード
		sp := math.Pow(2, float64(bits.OnesCount(uint((s>>n)&(1<<m-1)))))
		for i := 0; i < n+m; i++ {
			if s&(1<<i) == 0 {
				continue
			}
			for j := 0; j < n+m; j++ {
				// 現在地「i」から次に向かう「j」が集合に含まれていないか
				if i == j || (s&(1<<j)) != 0 {
					continue
				}
				// jが集合に含まれた場合、i→jの時間を加算して小さい方にする
				dp[s|1<<j][j] = math.Min(dp[s|1<<j][j], dp[s][i]+calcDistance(a[i], a[j])/sp)
			}
		}
	}

	var ans float64 = MaxInt
	// 下からn桁のbitが1で満たされている必要があるので、1 << nずつ増やした集合のどれかが最小
	for s := 1<<n - 1; s < 1<<(n+m); s += 1 << n {
		sp := math.Pow(2, float64(bits.OnesCount(uint((s>>n)&(1<<m-1)))))
		for i := 0; i < n+m; i++ {
			// 原点までの距離を加算した最小値
			ans = math.Min(ans, dp[s][i]+calcDistance(a[i], [2]float64{0, 0})/sp)
		}
	}

	w.WriteString(fmt.Sprintf("%.10f", ans))
}

func calcDistance(p, q [2]float64) float64 {
	return math.Sqrt((p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1]))
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
