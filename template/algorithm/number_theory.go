// 整数論
package algorithm

import "math"

// エラトステネスの篩
// 0から整数までの中で素数を列挙する
func Eratos(n int) []int {
	isPrimeMap := make(map[int]bool, n+1)
	for i := 0; i <= n; i++ {
		isPrimeMap[i] = true
	}
	isPrimeMap[0], isPrimeMap[1] = false, false

	for i := 2; i < n; i++ {
		if isPrimeMap[i] {
			// 素数iの倍数をnの範囲ですべて消す
			j := i * 2
			for j <= n {
				isPrimeMap[j] = false
				j += i
			}
		}
	}

	primeList := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		if isPrimeMap[i] {
			primeList = append(primeList, i)
		}
	}

	return primeList
}

// 対象の値が素数であるかチェックする
func IsPrime(x int) bool {
	if x == 2 {
		return true
	}
	if x < 2 || x%2 == 0 {
		return false
	}

	// 合成数xはp<=√xを満たす素因数を持つ
	// 上記性質を利用し、√xまでの値で素因数をもつかを確認することで
	// 判定をおこなう
	s := int(math.Sqrt(float64(x)))
	i := 3
	for i <= s {
		if x%i == 0 {
			return false
		}
		i += 2
	}

	return true
}

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

// Mで割ったあまりを求めるような場合に利用する
// 大きな値を扱うときに計算途中で余りを使える
const M = 1000000007

// べき乗
func Pow(x, n int) int {
	if n == 0 {
		return 1
	}

	ans := Pow(x*x%M, n/2)

	// 奇数のとき１回掛ける
	if n%2 == 1 {
		ans *= x % M
	}

	return ans
}
