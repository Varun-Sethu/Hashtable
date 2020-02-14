package hashtable

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Test that insertions are working effectively for linked list
func TestInsertion(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	list := &linkedList{}

	for i := 0; i < 100; i++ {
		generatedTestData := int32(-1000 + rand.Intn(2000))

		list.insert(keyValuePair{
			key: 0,
			value: generatedTestData,
		})

		// Check that data was actually inserted
		if list.tail.data.value != generatedTestData {
			t.Error(fmt.Sprintf("Failed to insert: %d, iteration: %d", generatedTestData, i))
		}
	}
}




// Assumes insertion is working perfectly fine
// Tests the search function
func TestSearch(t *testing.T) {
	// Generate a list of data to insert
	rand.Seed(time.Now().UTC().UnixNano())
	list := &linkedList{}

	randArray := make([]byte, 100)
	rand.Read(randArray)

	for i := range randArray {
		list.insert(keyValuePair{
			key: 0,
			value: int32(i),
		})

		if exists, _ := list.search(func(data keyValuePair) bool { 
			return int32(i) == data.value 
		}); !exists {
			t.Error(fmt.Sprintf("Failed to find: %d after insertion", i))
		}
	}
}




// Assumes insertion is working properly
// Tests that the size function is working
func TestSize(t *testing.T) {
	list := &linkedList{}

	for i := 0; i < 100; i++ {
		list.insert(keyValuePair{
			key: 0,
			value: 1,
		})
	}

	if list.size() != 100 {
		t.Error("Size function failed to return correct size of 100")
	}
}
