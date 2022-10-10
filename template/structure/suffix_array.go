package structure

import (
	"sort"
	"strconv"
	"strings"
)

// バケットソートで愚直に解く
// 検算用
func BucketSort(s string) []int {
	s = s + "$"
	m := make(map[rune][]int, len(s))
	for i, r := range s {
		m[r] = append(m[r], i)
	}

	keys := make([]rune, 0, len(m))
	for key, v := range m {
		keys = append(keys, key)
		sort.Slice(v, func(i, j int) bool { return strings.Compare(s[v[i]:], s[v[j]:]) < 0 })
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	ans := make([]int, 0, len(s))
	for _, key := range keys {
		ans = append(ans, m[key]...)
	}

	return ans
}

// SA-IS法を用いて接尾辞配列を生成
// Suffix Array Included Sort
func SAIS(s string) []int {
	s = s + "$"
	var (
		lflags = make([]bool, len(s))
		bucket = make([]int, 256)
	)
	// バケットソートのために文字ごとの数を数える
	for i := 0; i < len(s); i++ {
		bucket[int(s[i])]++
	}

	// L S の順番に並ぶLeft most Sをさがす
	isL := false
	for i := len(s) - 1; i > 0; i-- {
		lflags[i] = isL
		if s[i] != s[i-1] {
			isL = s[i] < s[i-1]
		}
	}

	sa := make([]int, len(s))
	for i := 0; i < len(sa); i++ {
		sa[i] = -1
	}
	{ // LMS-Typeのソート
		bucket := rbucket(bucket)
		for i := len(s) - 1; i >= 0; i-- {
			j := i - 1
			if !lflags[i] && j >= 0 && lflags[j] {
				c := s[i]
				sa[bucket[c]] = i // 下から埋めていく
				bucket[c]--
			}
		}
	}
	includedSort(sa, s, lflags, bucket)
	lmsSubstrList := make([]int, 0, len(s)/2)
	for _, i := range sa {
		if !lflags[i] && i-1 >= 0 && lflags[i-1] {
			lmsSubstrList = append(lmsSubstrList, i)
		}
	}
	// 部分文字列が複数なければ終了する
	if len(lmsSubstrList) == 1 {
		return sa[1:]
	}

	// 順位付けをして、順位の文字列としてSA=ISでsuffix arrayを求める
	orderList := getOrderList(s, lflags, lmsSubstrList)
	os := ""
	for _, order := range orderList {
		os += strconv.Itoa(order[0])
	}

	ossa := SAIS(os)
	{
		for i := 0; i < len(sa); i++ {
			sa[i] = -1
		}
		bucket := rbucket(bucket)
		for i := len(ossa) - 1; i >= 0; i-- {
			j := lmsSubstrList[orderList[ossa[i]][1]]
			c := s[j]
			sa[bucket[c]] = j
			bucket[c]--
		}
	}
	includedSort(sa, s, lflags, bucket)

	return sa[1:]
}

// LMSを使って、L-Typeの位置を求め、L-Typeの位置からS-Typeの位置を決める
func includedSort(sa []int, s string, lflags []bool, src []int) {
	{ // L-Typeのソート
		bucket := lbucket(src)
		for i := 0; i < len(sa); i++ {
			index := sa[i] - 1
			if sa[i] >= 0 && index >= 0 && lflags[index] {
				c := int(s[index])
				sa[bucket[c]] = index
				bucket[c]++ // 左詰めのインデクスを持ってるので右にずらす
			}
		}
	}

	{ // S-Typeのソート
		bucket := rbucket(src)
		for i := len(sa) - 1; i >= 0; i-- {
			index := sa[i] - 1
			if index >= 0 && !lflags[index] {
				c := s[index]
				sa[bucket[c]] = index
				bucket[c]--
			}
		}
	}
}

func getOrderList(s string, lflags []bool, lmsSubstrList []int) [][2]int {
	order := 0
	tmpOrderList := make([][2]int, len(s)/2) // lmsSubstrListの長さは最大でLとSが交互に来る場合なので、もとの文字列の長さの半分の長さにする
	for i := 1; i < len(lmsSubstrList); i++ {
		// 同じときは同じ順序番号を採番するため変わった場合、インクリメント
		if substr(lflags, s, lmsSubstrList[i-1]) != substr(lflags, s, lmsSubstrList[i]) {
			order++
		}

		tmpOrderList[lmsSubstrList[i]/2] = [2]int{order, i} // 順番と部分文字列のインデクスを紐づけとく
	}

	orderList := make([][2]int, 0, len(s)/2)
	for i := 0; i < len(tmpOrderList)-1; i++ {
		if tmpOrderList[i][0] > 0 {
			orderList = append(orderList, tmpOrderList[i])
		}
	}
	orderList = append(orderList, tmpOrderList[len(tmpOrderList)-1])
	return orderList
}

func lbucket(src []int) []int {
	bucket := make([]int, len(src))
	copy(bucket, src)
	sum := 0
	for i := 0; i < len(bucket); i++ {
		b := bucket[i]
		bucket[i] = sum
		sum += b
	}
	return bucket
}

func rbucket(src []int) []int {
	bucket := make([]int, len(src))
	copy(bucket, src)
	sum := -1
	for i := 0; i < len(bucket); i++ {
		sum += bucket[i]
		bucket[i] = sum
	}
	return bucket
}

func substr(lflags []bool, s string, i int) string {
	for j := i + 1; j < len(lflags); j++ {
		if !lflags[j] && j >= 0 && lflags[j-1] {
			return string(s[i:j])
		}
	}
	return string(s[i])
}
