package hashtable

// keyValuePair is an internal representation of a key and value
// Keyvalue pair seems relatively alligned, basically 4x32 bit blocks
type keyValuePair struct {
	key 	uint32
	value 	int32
	set 	bool  // allignment :(
}





// Two types of hash functions

// One is where collisions are resolved through bucketing

// ChainedHashTable is an implementation of a hashtable using chaining
type ChainedHashTable struct {
	data         []linkedList
	hashFunction func(uint32) uint32
}
// NewChainedTable generates a hashtable that uses chaining to resolve collisions
func NewChainedTable(size uint32) *ChainedHashTable {
	return &ChainedHashTable{
		data: make([]linkedList, size, size),
		hashFunction: UniversalHash(size),
	}
}


// Insert inserts data into a chained hashtable
func (table *ChainedHashTable) Insert(key uint32, value int32) {
	insertIndex := table.hashFunction(key)
	// Insert into table
	table.data[insertIndex].insert(keyValuePair{
		key: key,
		value: value,
	})
}


// Retreive method attains the value stored in a hash table
func (table *ChainedHashTable) Retreive(key uint32) (bool, keyValuePair) {
	// Determine which bucket to search
	insertIndex := table.hashFunction(key)
	// Generate the condition to determine if the key exists in the table
	condition := func(entry keyValuePair) bool {
		return entry.key == key
	}
	// Actually search it now
	return table.data[insertIndex].search(condition)
}









// The other is through linear probing
type ProbedHashTable struct {
	data		 	[]keyValuePair
	size		 	uint32
	hashFunction 	func(uint32) uint32
}
// NewProbedHashTable builds a new linearly probed hashtable
func NewProbedHashTable(size uint32) *ProbedHashTable {
	return &ProbedHashTable{
		data: make([]keyValuePair, size, size),
		size: size,
		hashFunction: UniversalHash(size),
	}
}


// Insert inserts data into the table using linear probing
func (table *ProbedHashTable) Insert(key uint32, value int32) {
	searchIndex := table.hashFunction(key)
	initialIndex := searchIndex

	// Increment the search index until we find an empty spot
	for table.data[searchIndex].set {
		searchIndex = (searchIndex + 1) % table.size
		// Terminating condtion of th search condition
		if searchIndex == initialIndex {return}
	}

	table.data[searchIndex] = keyValuePair{
		key: key,
		value: value,
		set: true,
	}
}


// Search for an entry into a hashtable
func (table *ProbedHashTable) Search(key uint32) (bool, keyValuePair) {
	searchIndex := table.hashFunction(key)
	initialIndex := searchIndex

	// While we have not found the index continue incrementing
	for table.data[searchIndex].key != key {
		// Looks like there was an empty bucket before finding our index, we can then this conclude that no index exists
		if !table.data[searchIndex].set {
			return false, keyValuePair{}
		}
		searchIndex = (searchIndex + 1) % table.size
		
		// Terminating condtion of th search condition
		if searchIndex == initialIndex {return false, keyValuePair{}}
	}

	return true, table.data[searchIndex]
}
