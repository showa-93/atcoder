package algorithm

// a-zの範囲の文字列辞書順検索で使える前処理
// n文字目以降で文字iが最初に登場するindexをもつ配列
// この前処理でN文字目以降に存在する文字列iが登場する最初のインデクスの探索がO(1)になる
func CalcNextIndex(s string) [][26]int {
	n := len(s)
	nex := make([][26]int, n+1)
	for i := 0; i < 26; i++ {
		nex[n][i] = n
	}
	for i := n - 1; i >= 0; i-- {
		nex[i] = nex[i+1]
		nex[i][s[i]-'a'] = i
	}

	return nex
}
