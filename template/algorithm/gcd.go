package algorithm

// 最大公約数
func GCD(x, y int) int {
	if x < y {
		x, y = y, x
	}

	for y > 0 {
		// もとのx,yの最大公約数と割ったあまりとyの最大公約数と一致する
		// yがゼロになるまで繰り返す
		r := x % y
		x, y = y, r
	}

	return x
}
