// 互いに素な集合を管理するデータ構造
package structure

type DijointSetsForests struct {
	parent []int
	rank   []int
}

func NewDijointSetsForests(n int) *DijointSetsForests {
	dsf := &DijointSetsForests{
		parent: make([]int, n),
		rank:   make([]int, n),
	}

	for i := 1; i < n; i++ {
		dsf.MakeSet(i)
	}

	return dsf
}

func (dsf *DijointSetsForests) MakeSet(x int) {
	dsf.rank[x] = 0
	dsf.parent[x] = x
}

// 引数の頂点が含まれる集合のルートを返す
func (dsf *DijointSetsForests) FindSet(x int) int {
	// 対象がルート出ない場合、親の探索をおこない、経路圧縮をおこなう
	// ルートを直接参照させる
	if x != dsf.parent[x] {
		dsf.parent[x] = dsf.FindSet(dsf.parent[x])
	}
	return dsf.parent[x]
}

func (dsf *DijointSetsForests) Same(x, y int) bool {
	return dsf.FindSet(x) == dsf.FindSet(y)
}

// ２つの頂点を結合する
func (dsf *DijointSetsForests) Unite(x, y int) {
	// ルートで結合するほうの比較をおこなう
	rootX, rootY := dsf.FindSet(x), dsf.FindSet(y)
	// 高さが高い方の木に低い木をつなげる
	if dsf.rank[rootX] > dsf.rank[rootY] {
		// 低い方のルートを変更する
		dsf.parent[rootY] = rootX
	} else {
		dsf.parent[rootX] = rootY
	}
	if dsf.rank[rootX] == dsf.rank[rootY] {
		// 同じ高さの場合、つなげると１つ高くなるので、加算する
		dsf.rank[rootY]++
	}
}
