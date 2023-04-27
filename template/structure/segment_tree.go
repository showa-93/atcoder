package structure

type RMQ struct {
	n      int
	size   int
	dst    []int
	minmax func(int, int) int
}

func NewRMQ(n int, size int, minmax func(int, int) int) *RMQ {
	rmq := &RMQ{
		n:      1,
		size:   size,
		minmax: minmax,
	}

	// 完全２分木になるようにnの値を計算する
	for rmq.n < n {
		rmq.n *= 2
	}
	rmq.dst = make([]int, 2*rmq.n-1)
	for i := 0; i < len(rmq.dst); i++ {
		rmq.dst[i] = rmq.size
	}

	return rmq
}

func (rmq *RMQ) Update(k, a int) {
	// 葉にあたる値は後ろn個の値を更新する
	k += rmq.n - 1
	rmq.dst[k] = a
	// セグメント木の葉側から順番に頂点に向かって更新する
	for k > 0 {
		k = (k - 1) / 2
		rmq.dst[k] = rmq.minmax(rmq.dst[2*k+1], rmq.dst[2*k+2])
	}
}

func (rmq *RMQ) Query(a, b int) int {
	return rmq.query(a, b, 0, 0, rmq.n)
}

func (rmq *RMQ) query(a, b, k, l, r int) int {
	// a, bの完全に範囲外の場合、存在しない
	if r <= a || b <= l {
		return rmq.size
	}

	// 範囲内の場合現在地が最小なのでそれを返す
	if a <= l && r <= b {
		return rmq.dst[k]
	}

	// 左分木、右分木について探索する
	vl := rmq.query(a, b, k*2+1, l, (l+r)/2)
	vr := rmq.query(a, b, k*2+2, (l+r)/2, r)

	return rmq.minmax(vl, vr)
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// binary indexed tree
// セグメント木の各ノードにある範囲の総和を記録する
// このとき、親の接点と左の子の接点があれば右の子の接点が求められるため
// 最小限の接点だけで構築した木をBinary indexed treeと呼ぶ
// この木は接点の番号を２進数表記したときの性質を利用することで
// 高速に和を求めることができる O(log n)
type BIT struct {
	n   int
	bit []int
}

func (b *BIT) Sum(i int) int {
	var sum int
	// 最後の1bitを減産しながら和をとることで
	// [1 i]の区間の総和を高速に取得できる
	for i > 0 {
		sum += b.bit[i]
		i -= i & -i
	}
	return sum
}

func (b *BIT) Add(i, x int) {
	for i <= b.n {
		b.bit[i] += x
		i += i & -i
	}
}
