package segmenttree

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
