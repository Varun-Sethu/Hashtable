package hashtable

import (
	"math/rand"
	"time"
)


// UniversalHash is a simple closure method that returns a hashfunction
func UniversalHash(size uint64) func(HashType) uint64 {
	rand.Seed(time.Now().UTC().UnixNano())

	// Implementation of universal hashing according to the Carter and Wegman method: https://www.cs.princeton.edu/courses/archive/fall09/cos521/Handouts/universalclasses.pdf
	mersenneIntegers := []uint{19, 31, 61}
	// Compute a mersenne prime: 2^n - 1
	p := uint64((1 << mersenneIntegers[rand.Intn(2)]) - 1)
	// Generate constants
	a := uint64(1 + rand.Intn(int(p-1)))
	b := uint64(rand.Intn(int(p)))

	return func(data HashType) uint64 {
		x := data.Serialize()
		baseVal := a*x + b

		return ((baseVal) % p) % size
	}
}
