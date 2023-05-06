package algorithm

// ソート済みのスライスから指定された要素以上の値が現れる最初の位置を返す
func LowerBound(list []int, value int) int {
	l, r := 0, len(list)-1
	for r-l >= 0 {
		c := (l + r) / 2
		if value <= list[c] {
			r = c - 1
		} else {
			l = c + 1
		}
	}

	return l
}

// ソート済みのスライスから指定された要素より大きい値が現れる最初の位置を返す
// より大きなものがなければ、len(list)の値になる
func UpperBound(list []int, value int) int {
	l, r := 0, len(list)-1
	for r-l >= 0 {
		c := (l + r) / 2
		if value < list[c] {
			r = c - 1
		} else {
			l = c + 1
		}
	}

	return l
}
