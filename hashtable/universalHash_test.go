package hashtable

import (
	"fmt"
	"testing"
)


func TestUniversalHashing(t *testing.T) {
	// Generate a hash function and if executing it doesnt fail then it passes the test
	function := UniversalHash(100)
	fmt.Print(function(10))
}
