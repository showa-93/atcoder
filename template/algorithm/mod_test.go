package algorithm

import "testing"

func TestMint_Add(t *testing.T) {
	if Mint(1).Add(2) != Mint(3) {
		t.Fail()
	}
	if Mint(1000000006).Add(2) != Mint(1) {
		t.Fail()
	}
	if Mint(-2).Add(1) != Mint(-1) {
		t.Fail()
	}
}

func TestMint_Sub(t *testing.T) {
	if v := Mint(3).Sub(1); v != Mint(2) {
		t.Error(v)
	}
	if v := Mint(1).Sub(2); v != Mint(1000000006) {
		t.Error(v)
	}
	if v := Mint(-1).Sub(2); v != Mint(1000000004) {
		t.Error(v)
	}
	if v := Mint(0).Sub(2); v != Mint(1000000005) {
		t.Error(v)
	}
}

func TestMint_Mul(t *testing.T) {
	if v := Mint(3).Mul(2); v != Mint(6) {
		t.Error(v)
	}
	if v := Mint(500000004).Mul(2); v != Mint(1) {
		t.Error(v)
	}
}

func TestMint_Pow(t *testing.T) {
	if v := Mint(3).Pow(6); v != Mint(729) {
		t.Error(v)
	}
	modValue = 100
	defer func() { modValue = Mint(mod1000000007) }()
	if v := Mint(3).Pow(19); v != Mint(67) {
		t.Error(v)
	}
}

func TestMint_Div(t *testing.T) {
	if v := Mint(10).Div(2); v != Mint(5) {
		t.Error(v)
	}
	if v := Mint(9).Div(2); v != Mint(500000008) {
		t.Error(v)
	}
}
