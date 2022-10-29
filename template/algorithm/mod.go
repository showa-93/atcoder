package algorithm

const (
	mod998244353  int = 998244353
	mod1000000007 int = 1000000007
)

func SetModValue(v int) {
	mod = v
}

var mod int = mod1000000007

func Mod(a int) int {
	a %= mod
	if a < 0 {
		a += mod
	}
	return a
}

func ModAdd(a, b int) int {
	return Mod(a + b)
}

func ModSub(a, b int) int {
	return ModAdd(a, -b)
}

func ModMul(a, b int) int {
	return Mod(a * b)
}

func ModPow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 == 1 {
			p = ModMul(p, a)
		}
		a = ModMul(a, a)
		b >>= 1
	}

	return p
}

// 非再帰拡張Euclidの互除法
func ModInv(a int) int {
	b := mod
	x, y := 1, 0

	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		x -= t * y
		x, y = y, x
	}

	return Mod(x)
}

func ModDiv(a, b int) int {
	return ModMul(Mod(a), ModInv(b))
}
