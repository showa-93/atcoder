package segmenttree

type RSQ struct {
	n      int
	dst    []int
	minmax func(int, int) int
}

func NewRSQ(n int) *RSQ {
	rsq := &RSQ{
		n: 1,
	}

	// 完全２分木になるようにnの値を計算する
	for rsq.n < n {
		rsq.n *= 2
	}
	rsq.dst = make([]int, 2*rsq.n-1)
	for i := 0; i < len(rsq.dst); i++ {
		rsq.dst[i] = 0
	}

	return rsq
}

func (rsq *RSQ) Update(k, a int) {
	// 葉にあたる値は後ろn個の値を更新する
	k += rsq.n - 1
	rsq.dst[k] = a
	// セグメント木の葉側から順番に頂点に向かって更新する
	for k > 0 {
		k = (k - 1) / 2
		rsq.dst[k] = rsq.dst[2*k+1] + rsq.dst[2*k+2]
	}
}

func (rsq *RSQ) Query(a, b int) int {
	return rsq.query(a, b, 0, 0, rsq.n)
}

func (rsq *RSQ) query(a, b, k, l, r int) int {
	// a, bの完全に範囲外の場合、存在しない
	if r <= a || b <= l {
		return 0
	}

	// 範囲内の場合現在地が最小なのでそれを返す
	if a <= l && r <= b {
		return rsq.dst[k]
	}

	// 左分木、右分木について探索する
	vl := rsq.query(a, b, k*2+1, l, (l+r)/2)
	vr := rsq.query(a, b, k*2+2, (l+r)/2, r)

	return vl + vr
}
