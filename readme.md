# Hashtable

Simple hashtable implementation in golang just for fun, two methods of resolving collisions. The hashtable implements a universal class of hash functions in order to map data from the universal set to the table. More details at: https://www.cs.princeton.edu/courses/archive/fall09/cos521/Handouts/universalclasses.pdf
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
