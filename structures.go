package hashtable

// Common structures just used throughout the implementation

// HashType is any hashable type, must be implemented for universal hash to work effectively as otherwise there would be nothing to hash
type HashType interface {
	Serialize() uint64
}


// keyValuePair is an internal representation of a key and value
type keyValuePair struct {
	key 	HashType
	value 	interface{}
}

