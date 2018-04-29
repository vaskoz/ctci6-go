package three

// Elem represents a node in a singly linked list.
type Elem struct {
	Value interface{}
	next  *Elem
}

// DeleteMiddle shifts all the elements of the singly linked list back one position.
// Thereby deleting the value in the parameter elem.
func DeleteMiddle(elem *Elem) {
	for elem.next != nil {
		elem.Value = elem.next.Value
		elem = elem.next
	}
}
