package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModAdd(t *testing.T) {
	tests := map[string]struct {
		a, b int
		want int
	}{
		"1 + 2 ≡ 3":           {1, 2, 3},
		"1000000006 + 2 ≡ 1":  {1000000006, 2, 1},
		"-2 + 1 ≡ 1000000006": {-2, 1, 1000000006},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ModAdd(tt.a, tt.b))
		})
	}
}

func TestModSub(t *testing.T) {
	tests := map[string]struct {
		a, b int
		want int
	}{
		"3 - 2 ≡ 1":           {3, 2, 1},
		"1000000009 - 1 ≡ 1":  {1000000009, 1, 1},
		"-2 - 1 ≡ 1000000004": {-2, 1, 1000000004},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ModSub(tt.a, tt.b))
		})
	}
}

func TestModMul(t *testing.T) {
	tests := map[string]struct {
		a, b int
		want int
	}{
		"3 * 2 ≡ 6":           {3, 2, 6},
		"500000004 * 2 ≡ 1":   {500000004, 2, 1},
		"3 * -2 ≡ 1000000001": {3, -2, 1000000001},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ModMul(tt.a, tt.b))
		})
	}
}

func TestModPow(t *testing.T) {
	tests := map[string]struct {
		a, b int
		want int
	}{
		"3 ^ 6 ≡ 729":        {3, 6, 729},
		"22 ^ 7 ≡ 494357874": {22, 7, 494357874},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ModPow(tt.a, tt.b))
			t.Error()
		})
	}
}

func TestModInv(t *testing.T) {
	tests := map[string]struct {
		a    int
		want int
	}{
		"12345678900000 ^ -1": {12345678900000, 237800188},
		"213134656876 ^ -1":   {213134656876, 50269932},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ModInv(tt.a))
		})
	}
}

func TestModDiv(t *testing.T) {
	tests := map[string]struct {
		a, b int
		want int
	}{
		"10 / 2 ≡ 5":                          {10, 2, 5},
		"9 / 2 ≡ 500000008":                   {9, 2, 500000008},
		"12345678900000 / 100000 ≡ 123456789": {12345678900000, 100000, 123456789},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, ModDiv(tt.a, tt.b))
		})
	}
}
