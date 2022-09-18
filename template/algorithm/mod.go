package algorithm

const (
	mod998244353  int = 998244353
	mod1000000007 int = 1000000007
)

var modValue Mint = Mint(mod1000000007)

func SetModValue(v int) {
	modValue = Mint(v)
}

type Mint int

func (m Mint) Mod() Mint {
	return m % modValue
}

func (m Mint) Inv() Mint {
	return m.Pow(Mint(0).Sub(2))
}

func (m Mint) Add(a Mint) Mint {
	return Mint(m + a).Mod()
}

func (m Mint) Sub(a Mint) Mint {
	v := Mint(m - a).Mod()
	if v < 0 {
		v += modValue
	}

	return v
}

func (m Mint) Mul(a Mint) Mint {
	return Mint(m * a).Mod()
}

func (m Mint) Div(a Mint) Mint {
	return m.Mul(a.Inv())
}

func (m Mint) Pow(n Mint) Mint {
	var p Mint = 1
	base := m
	for n > 0 {
		// nの2進数表記ごとに1のくらいの時だけかける
		if n&1 == 1 {
			p = p.Mul(base)
		}
		base = base.Mul(base)
		n >>= 1
	}

	return p
}
