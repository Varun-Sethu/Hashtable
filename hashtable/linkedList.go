package hashtable

// Linked list implementation

type node struct {
	data keyValuePair
	next *node
}

// Linkedlist data structure
type linkedList struct {
	head *node
	tail *node // simply for testing purposes
}
type searchCondition func(keyValuePair) bool




// Basic methods
// Insertion method
func (list *linkedList) insert(dataToInsert keyValuePair) {
	// Create the block in the linked list that will store the data
	data := &node{
		data: dataToInsert,
		next: nil,
	}

	// bit of a mess
	if list.head == nil {
		list.head = data
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		// Looks like a current node with a nil value for the pointer to the next chain in the list, set its next value to this new data block
		currentNode.next = data
	}
	list.tail = data // just set the tail for testing purposes
}




// search method
// iteratively search over the head node and its next ndoes to see if they satisfy a condition
func (list *linkedList) search(condition searchCondition) (bool, keyValuePair) {
	currentNode := list.head
	for currentNode != nil {
		if condition(currentNode.data) {
			return true, currentNode.data
		}
		currentNode = currentNode.next
	}

	return false, keyValuePair{}
}





// Size operation simply determines the number of nodes in the linked list
func (list *linkedList) size() int {
	currentNode := list.head
	var i int
	for i = 0; currentNode != nil; i++ {
		currentNode = currentNode.next
	}

	return i
}
