package rf1r

import (
	"github.com/dchest/siphash"
	"github.com/dgryski/go-jump"
)

// RF1R - Replica Factor 1 with random
type RF1R struct {
	hasher1 func([]byte) uint64
	hasher2 func([]byte) uint64
}

// New Replica Factor 1 with random
func New(randomized bool) *RF1R {
	res := RF1R{}

	if randomized {
		res.hasher1 = func(b []byte) uint64 { return siphash.Hash(0, 0, b) }
		res.hasher2 = func(b []byte) uint64 { return siphash.Hash(2, 3, b) }
	} else {
		res.hasher1 = func(b []byte) uint64 { return siphash.Hash(0, 0, b) }
		res.hasher2 = func(b []byte) uint64 { return siphash.Hash(0, 0, b) }
	}

	return &res
}

func (data *RF1R) Choose(key string, nodes int) []int {
	res := make([]int, 2)
	res[0] = int(jump.Hash(data.hasher1([]byte(key)), nodes/2))
	res[1] = nodes/2 + int(jump.Hash(data.hasher2([]byte(key)), nodes/2))

	return res
}
