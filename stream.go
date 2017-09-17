package spigot

import (
	"fmt"
	"io"
	"math/big"
)

type LFT struct {
	q, r, s, t *big.Int
}

func extr(l LFT, x *big.Int) *big.Int {
	return new(big.Int).Quo(
		new(big.Int).Add(new(big.Int).Mul(l.q, x), l.r),
		new(big.Int).Add(new(big.Int).Mul(l.s, x), l.t),
	)
}

func comp(a, b LFT) LFT {
	return LFT{
		new(big.Int).Add(new(big.Int).Mul(a.q, b.q), new(big.Int).Mul(a.r, b.s)),
		new(big.Int).Add(new(big.Int).Mul(a.q, b.r), new(big.Int).Mul(a.r, b.t)),
		new(big.Int).Add(new(big.Int).Mul(a.s, b.q), new(big.Int).Mul(a.t, b.s)),
		new(big.Int).Add(new(big.Int).Mul(a.s, b.r), new(big.Int).Mul(a.t, b.t)),
	}
}

func PiStream(w io.Writer) {
	var (
		zero   = big.NewInt(0)
		one    = big.NewInt(1)
		two    = big.NewInt(2)
		three  = big.NewInt(3)
		four   = big.NewInt(4)
		ten    = big.NewInt(10)
		negten = big.NewInt(-10)
	)
	z := LFT{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)}
	k := big.NewInt(0)
	for {
		if y := extr(z, three); extr(z, four).Cmp(y) == 0 {
			fmt.Fprint(w, y)
			z = comp(LFT{
				new(big.Int).Set(ten),
				new(big.Int).Mul(negten, y),
				new(big.Int).Set(zero),
				new(big.Int).Set(one),
			}, z)
		}
		k.Add(k, one)
		z = comp(z, LFT{
			new(big.Int).Set(k),
			new(big.Int).Add(new(big.Int).Mul(four, k), two),
			new(big.Int).Set(zero),
			new(big.Int).Add(new(big.Int).Mul(two, k), one),
		})
	}
}
