package hashtable

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)




// Test the insertion of the hash table
func TestHashTableInsert(t *testing.T)  {
	table := NewChainedTable(100)
	probedTable := NewProbedHashTable(100)
	table.Insert(
		197, 31)
	probedTable.Insert(
		197, 31)
	
	

	// Determine if the value was entered
	list := table.data[table.hashFunction(197)]
	insertedVal := probedTable.data[probedTable.hashFunction(197)]
	exists, _ := list.search(func (data keyValuePair) bool {
		return data.value == 31
	})

	if !exists || insertedVal.value != 31 {t.Error("Could not find the inserted value or the ey mapped to a different value")}
}


// Test the searching of the hashtable
func TestHashTableRetreival(t *testing.T) {
	table := NewChainedTable(100)
	table.Insert(
		197, 31)
	_, kv := table.Retreive(197)	

	if kv.value != 31 {t.Error("Could not find the inserted value or the ey mapped to a different value")}
}



// Tests terminating conditions of the functions
func TestBounds(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	table := NewProbedHashTable(100)
	
	for i := 0; i < 100; i++ {
		key, value := -1000 + rand.Intn(2000), -1000 + rand.Intn(2000)
		table.Insert(uint32(key), int32(value))
		exists, _ := table.Search(uint32(key))
		if !exists {t.Error("An insertion failed")}
	}


	table.Insert(6000, 170)
	exists, _ := table.Search(20000)
	if exists {t.Error("The overload test entry was inserted when it shouldn't have been")}
}













// Benchmarking, these just benchmark the hashtable
func BenchmarkChainedInsertion(b *testing.B) {
	fmt.Print(b.N, "\n")
	runtime.GC()
	table := NewProbedHashTable(uint32(b.N))
	b.ResetTimer()

	// Iterate n times
	for i := 0; i < b.N; i++ {
		table.Insert(uint32(i), 100)
	}
	fmt.Printf("\ncompleted test: ")
}