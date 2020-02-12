package hashtable

import (
	"testing"
)


func TestUniversalHashing(t *testing.T) {
	// Generate a hash function and if executing it doesnt fail then it passes the test
	function := UniversalHash(100)
	x := hashableint{Value: 10}
	function(x)
}
