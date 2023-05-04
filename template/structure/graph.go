package structure

// 残余グラフの辺
type ResidualEdge struct {
	to  int // 次の頂点
	cap int // この辺で流せる容量
	rev int // 次の頂点の連結リストの位置
}

// 残余グラフ
type ResidualGraph struct {
	g       [][]ResidualEdge
	visited []bool
}

func NewResidualGraph(n int) *ResidualGraph {
	return &ResidualGraph{
		g:       make([][]ResidualEdge, n),
		visited: make([]bool, n),
	}
}

func (r *ResidualGraph) AddEdge(a, b, cap int) {
	r.g[a] = append(r.g[a], ResidualEdge{b, cap, len(r.g[b])})
	r.g[b] = append(r.g[b], ResidualEdge{a, 0, len(r.g[a]) - 1})
}

func (r *ResidualGraph) dfs(pos, goal, flow int) int {
	if pos == goal {
		return flow
	}

	r.visited[pos] = true
	for i, e := range r.g[pos] {
		if e.cap == 0 || r.visited[e.to] {
			continue
		}

		fixedFlow := r.dfs(e.to, goal, r.Min(flow, e.cap))
		if fixedFlow > 0 {
			r.g[pos][i].cap -= fixedFlow
			r.g[e.to][e.rev].cap += fixedFlow
			return fixedFlow
		}
	}

	return 0
}

// Ford-Fulkerson法で残余グラフから最大フローを計算する
func (r *ResidualGraph) MaxFlow(s, t int) int {
	var flow int
	for {
		for i := 0; i < len(r.visited); i++ {
			r.visited[i] = false
		}
		f := r.dfs(s, t, 1000000000)
		// すべての容量を使い切るまで流すとゴールに到達するフローがなくなる
		// その時点で最大流量になる
		if f == 0 {
			return flow
		}
		flow += f
	}
}

func (r *ResidualGraph) Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
