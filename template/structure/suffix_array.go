package structure

import (
	"bytes"
	"reflect"
	"sort"
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

type SuffixArray struct {
	data []byte
	sa   []int
}

// SA-IS法を用いて接尾辞配列を生成
// Suffix Array Included Sort
func NewSAIS(text []byte) *SuffixArray {
	sa := &SuffixArray{
		data: text,
	}
	satmp := make([]int, len(text)+1)
	SAIS(text, satmp)
	sa.sa = satmp[1:]

	return sa
}

func SAIS(text []byte, sa []int) {
	var (
		bucket = make([]int, 256)
	)
	sais(text, sa, bucket)
}

func sais(text []byte, sa, bucket []int) {
	if len(text) == 0 {
		return
	}
	if len(text) == 1 {
		sa[1] = 0
		return
	}
	text = append(text, 0)
	var (
		lflags = make([]bool, len(text))
		freq   = make([]int, 256)
	)

	for _, b := range text {
		freq[b]++
	}

	// L S の順番に並ぶLeft most Sをさがす
	isL := false
	for i := len(text) - 1; i > 0; i-- {
		lflags[i] = isL
		if text[i] != text[i-1] {
			isL = text[i] < text[i-1]
		}
	}
	var lmsSubstrCount int
	for i := 1; i < len(text); i++ {
		if !lflags[i] && lflags[i-1] {
			lmsSubstrCount++
		}
	}

	for i := 0; i < len(sa); i++ {
		sa[i] = -1
	}
	{ // LMS-Typeのソート
		rbucket(bucket, freq)
		for i := len(text) - 1; i >= 0; i-- {
			j := i - 1
			if !lflags[i] && j >= 0 && lflags[j] {
				c := text[i]
				sa[bucket[c]] = i // 下から埋めていく
				bucket[c]--
			}
		}
	}
	includedSort(sa, text, lflags, bucket, freq)

	// 部分文字列が複数なければ終了する
	if lmsSubstrCount <= 1 {
		return
	}

	// 順位付けをして、順位の文字列としてSA=ISでsuffix arrayを求める
	orderList, maxOrder := getOrderList(text, lflags, sa)
	os := make([]int, len(orderList), len(orderList)+1)
	for i, order := range orderList {
		os[i] = order[0]
	}

	ossa := make([]int, len(os)+1)
	sais_int(os, ossa, make([]int, maxOrder+1))
	ossa = ossa[1:]
	{
		for i := 0; i < len(sa); i++ {
			sa[i] = -1
		}
		rbucket(bucket, freq)
		for i := len(ossa) - 1; i >= 0; i-- {
			j := orderList[ossa[i]][1]
			c := text[j]
			sa[bucket[c]] = j
			bucket[c]--
		}
	}

	includedSort(sa, text, lflags, bucket, freq)

	return
}

// LMSを使って、L-Typeの位置を求め、L-Typeの位置からS-Typeの位置を決める
func includedSort(sa []int, text []byte, lflags []bool, bucket, freq []int) {
	{ // L-Typeのソート
		lbucket(bucket, freq)
		for i := 0; i < len(sa); i++ {
			index := sa[i] - 1
			if sa[i] >= 0 && index >= 0 && lflags[index] {
				c := text[index]
				sa[bucket[c]] = index
				bucket[c]++ // 左詰めのインデクスを持ってるので右にずらす
			}
		}
	}

	{ // S-Typeのソート
		rbucket(bucket, freq)
		for i := len(sa) - 1; i >= 0; i-- {
			index := sa[i] - 1
			if index >= 0 && !lflags[index] {
				c := text[index]
				sa[bucket[c]] = index
				bucket[c]--
			}
		}
	}
}

func getOrderList(text []byte, lflags []bool, sa []int) ([][2]int, int) {
	textLen := len(text)
	order := 1
	orderList := make([][2]int, textLen)
	j := sa[0]
	for _, i := range sa[1:] {
		if i-1 >= 0 && !lflags[i] && lflags[i-1] {
			// 同じときは同じ順序番号を採番するため変わった場合、インクリメント
			if bytes.Compare(substr(lflags, text, j), substr(lflags, text, i)) != 0 {
				order++
			}

			orderList[i/2] = [2]int{order, i} // 順番と部分文字列のインデクスを紐づけとく
			j = i
		}
	}

	j = textLen / 2
	for i := 0; i < textLen/2; i++ {
		if orderList[i][0] > 0 {
			orderList[j] = orderList[i]
			j++
		}
	}
	orderList[j] = [2]int{order, sa[0]}
	j++

	return orderList[textLen/2 : j], order
}

func lbucket(bucket, freq []int) {
	for i := 0; i < len(bucket); i++ {
		bucket[i] = 0
	}
	sum := 0
	for i := 0; i < len(bucket); i++ {
		b := freq[i]
		bucket[i] = sum
		sum += b
	}
}

func rbucket(bucket, freq []int) {
	for i := 0; i < len(bucket); i++ {
		bucket[i] = 0
	}
	sum := -1
	for i := 0; i < len(bucket); i++ {
		sum += freq[i]
		bucket[i] = sum
	}
}

func substr(lflags []bool, text []byte, i int) []byte {
	for j := i + 1; j < len(lflags); j++ {
		if !lflags[j] && j >= 0 && lflags[j-1] {
			return text[i:j]
		}
	}
	return text[i : i+1]
}

// 順序の配列のsuffix arrayを構築するようの巻数
func sais_int(ints []int, sa, bucket []int) {
	if len(ints) == 0 {
		return
	}
	if len(ints) == 1 {
		sa[0] = 0
		return
	}
	ints = append(ints, 0)
	var (
		lflags = make([]bool, len(ints))
		freq   = make([]int, len(bucket))
	)

	for _, b := range ints {
		freq[b]++
	}

	// L S の順番に並ぶLeft most Sをさがす
	isL := false
	for i := len(ints) - 1; i > 0; i-- {
		lflags[i] = isL
		if ints[i] != ints[i-1] {
			isL = ints[i] < ints[i-1]
		}
	}
	var lmsSubstrCount int
	for i := 1; i < len(ints); i++ {
		if !lflags[i] && lflags[i-1] {
			lmsSubstrCount++
		}
	}

	for i := 0; i < len(sa); i++ {
		sa[i] = -1
	}
	{ // LMS-Typeのソート
		rbucket_int(bucket, freq)
		for i := len(ints) - 1; i >= 0; i-- {
			j := i - 1
			if !lflags[i] && j >= 0 && lflags[j] {
				c := ints[i]
				sa[bucket[c]] = i // 下から埋めていく
				bucket[c]--
			}
		}
	}
	includedSort_int(sa, ints, lflags, bucket, freq)

	// 部分文字列が複数なければ終了する
	if lmsSubstrCount <= 1 {
		return
	}

	// 順位付けをして、順位の文字列としてSA=ISでsuffix arrayを求める
	orderList, maxOrder := getOrderList_int(ints, lflags, sa)
	os := make([]int, len(orderList), len(orderList)+1)
	for i, order := range orderList {
		os[i] = order[0]
	}

	ossa := make([]int, len(os)+1)
	sais_int(os, ossa, make([]int, maxOrder+1))
	ossa = ossa[1:]
	{
		for i := 0; i < len(sa); i++ {
			sa[i] = -1
		}
		rbucket_int(bucket, freq)
		for i := len(ossa) - 1; i >= 0; i-- {
			j := orderList[ossa[i]][1]
			c := ints[j]
			sa[bucket[c]] = j
			bucket[c]--
		}
	}

	includedSort_int(sa, ints, lflags, bucket, freq)

	return
}

func includedSort_int(sa []int, ints []int, lflags []bool, bucket, freq []int) {
	{ // L-Typeのソート
		lbucket_int(bucket, freq)
		for i := 0; i < len(sa); i++ {
			index := sa[i] - 1
			if sa[i] >= 0 && index >= 0 && lflags[index] {
				c := ints[index]
				sa[bucket[c]] = index
				bucket[c]++ // 左詰めのインデクスを持ってるので右にずらす
			}
		}
	}

	{ // S-Typeのソート
		rbucket_int(bucket, freq)
		for i := len(sa) - 1; i >= 0; i-- {
			index := sa[i] - 1
			if index >= 0 && !lflags[index] {
				c := ints[index]
				sa[bucket[c]] = index
				bucket[c]--
			}
		}
	}
}

func getOrderList_int(ints []int, lflags []bool, sa []int) ([][2]int, int) {
	textLen := len(ints)
	order := 1
	orderList := make([][2]int, textLen)
	j := sa[0]
	for _, i := range sa[1:] {
		if i-1 >= 0 && !lflags[i] && lflags[i-1] {
			// 同じときは同じ順序番号を採番するため変わった場合、インクリメント
			if !reflect.DeepEqual(substr_int(lflags, ints, j), substr_int(lflags, ints, i)) {
				order++
			}

			orderList[i/2] = [2]int{order, i} // 順番と部分文字列のインデクスを紐づけとく
			j = i
		}
	}

	j = textLen / 2
	for i := 0; i < textLen/2; i++ {
		if orderList[i][0] > 0 {
			orderList[j] = orderList[i]
			j++
		}
	}
	orderList[j] = [2]int{order, sa[0]}
	j++

	return orderList[textLen/2 : j], order
}

func lbucket_int(bucket, freq []int) {
	for i := 0; i < len(bucket); i++ {
		bucket[i] = 0
	}
	sum := 0
	for i := 0; i < len(bucket); i++ {
		b := freq[i]
		bucket[i] = sum
		sum += b
	}
}

func rbucket_int(bucket, freq []int) {
	for i := 0; i < len(bucket); i++ {
		bucket[i] = 0
	}
	sum := -1
	for i := 0; i < len(bucket); i++ {
		sum += freq[i]
		bucket[i] = sum
	}
}

func substr_int(lflags []bool, ints []int, i int) []int {
	for j := i + 1; j < len(lflags); j++ {
		if !lflags[j] && j > 0 && lflags[j-1] {
			return ints[i:j]
		}
	}
	return ints[i : i+1]
}
