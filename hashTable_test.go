package hashtable

import "testing"



// Types for serialisation
type hashableint struct {
	Value uint64
}
func (v hashableint) Serialize() uint64 {
	return v.Value
}


// Test the insertion of the hash table
func TestHashTableInsert(t *testing.T)  {
	table := NewChainedTable(100)
	table.Insert(
		hashableint{Value: 197}, 31)
	

	// Determine if the value was entered
	list := table.data[table.hashFunction(hashableint{Value: 197})]
	exists, _ := list.search(func (data keyValuePair) bool {
		return data.value.(int) == 31
	})

	if !exists {t.Error("Could not find the inserted value or the ey mapped to a different value")}
}


// Test the searching of the hashtable
func TestHashTableRetreival(t *testing.T) {
	table := NewChainedTable(100)
	table.Insert(
		hashableint{Value: 197}, 31)
	_, kv := table.Retreive(hashableint{Value: 197})	

	if kv.value.(int) != 31 {t.Error("Could not find the inserted value or the ey mapped to a different value")}
}