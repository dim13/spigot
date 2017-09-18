package spigot

import (
	"math/big"
)

type LFT [4]*big.Int

func extr(l LFT, x *big.Int) *big.Int {
	return new(big.Int).Quo(
		new(big.Int).Add(new(big.Int).Mul(l[0], x), l[1]),
		new(big.Int).Add(new(big.Int).Mul(l[2], x), l[3]),
	)
}

func comp(a, b LFT) LFT {
	return LFT{
		new(big.Int).Add(new(big.Int).Mul(a[0], b[0]), new(big.Int).Mul(a[1], b[2])),
		new(big.Int).Add(new(big.Int).Mul(a[0], b[1]), new(big.Int).Mul(a[1], b[3])),
		new(big.Int).Add(new(big.Int).Mul(a[2], b[0]), new(big.Int).Mul(a[3], b[2])),
		new(big.Int).Add(new(big.Int).Mul(a[2], b[1]), new(big.Int).Mul(a[3], b[3])),
	}
}

func PiStream(c chan<- int64) {
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
			c <- y.Int64()
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
