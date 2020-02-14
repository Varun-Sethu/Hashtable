# Hashtable

Simple hashtable implementation in golang just for fun, two methods of resolving collisions
 - Chaining
 - Linear probing

## Usage
Just include the package, but honestly don't ever use this... just use map
```go
func main() {
    table := NewChainedTable(size)
    table.Insert(3, 2)

    table := NewProbedHashTable(size)
    table.Insert(3, 2)
}
```