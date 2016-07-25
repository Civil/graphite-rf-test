package rf2

import (
	"github.com/dchest/siphash"
	"github.com/dgryski/go-jump"
)

// RF2 - Replica Factor 2
type RF2 struct {
	hasher func([]byte) uint64
}

// New Replica Factor 1 with random
func New() *RF2 {
	res := RF2{}

	res.hasher = func(b []byte) uint64 { return siphash.Hash(0, 0, b) }

	return &res
}

func (data *RF2) Choose(key string, nodes int) []int {
	hkey := data.hasher([]byte(key))
	res := make([]int, 2)

	buckets := make([]int, nodes)
	for i := range buckets {
		buckets[i] = i
	}

	b := jump.Hash(hkey, len(buckets))
	res[0] = buckets[b]
	buckets[b], buckets = buckets[len(buckets)-1], buckets[:len(buckets)-1]
	b = jump.Hash(xorshiftMult64(hkey), len(buckets))
	res[1] = buckets[b]

	return res
}

// 64-bit xorshift multiply rng from http://vigna.di.unimi.it/ftp/papers/xorshift.pdf
func xorshiftMult64(x uint64) uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * 2685821657736338717
}
