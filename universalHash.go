package hashtable

import (
	"math/rand"
	"time"
)


// UniversalHash is a simple closure method that returns a hashfunction
func UniversalHash(size uint32) func(uint32) uint32 {
	rand.Seed(time.Now().UTC().UnixNano())

	// Implementation of universal hashing according to the Carter and Wegman method: https://www.cs.princeton.edu/courses/archive/fall09/cos521/Handouts/universalclasses.pdf
	mersenneIntegers := []uint8{19, 31, 61}
	// Compute a mersenne prime: 2^n - 1
	p := 1 << (mersenneIntegers[rand.Intn(2)]) - 1
	// Generate constants
	a := 1 + rand.Intn(p-1)
	b := rand.Intn(p)

	return func(data uint32) uint32 {
		baseVal := uint32(a)*data + uint32(b)

		return ((baseVal) % uint32(p)) % size
	}
}
