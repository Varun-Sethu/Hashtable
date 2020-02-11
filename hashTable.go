package hashtable

// Two types of hash functions

// One is where collisions are resolved through bucketing

// ChainedHashTable is an implementation of a hashtable using chaining
type ChainedHashTable struct {
	data         []linkedList
	hashFunction func(HashType) uint64
}


// NewChainedTable generates a hashtable that uses chaining to resolve collisions
func NewChainedTable(size uint64) *ChainedHashTable {
	return &ChainedHashTable{
		data: make([]linkedList, size, size),
		hashFunction: UniversalHash(size),
	}
}

// Insert inserts data into a chained hashtable
func (table *ChainedHashTable) Insert(key HashType, value interface{}) {
	insertIndex := table.hashFunction(key)
	// Insert into table
	table.data[insertIndex].insert(keyValuePair{
		key: key,
		value: value,
	})
}

// Retreive method attains the value stored in a hash table
func (table *ChainedHashTable) Retreive(key HashType) (bool, keyValuePair) {
	// Determine which bucket to search
	insertIndex := table.hashFunction(key)
	// Generate the condition to determine if the key exists in the table
	condition := func(entry keyValuePair) bool {
		return entry.key == key
	}
	// Actually search it now
	return table.data[insertIndex].search(condition)
}

// The other is through quadratic probing
