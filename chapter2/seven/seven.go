package seven

// Elem represents a node in a singly linked list.
type Elem struct {
	Value int
	next  *Elem
}

// Intersection returns a list that represents the common elements to both input lists.
func Intersection(first, second *Elem) *Elem {
	fl, sl := Length(first), Length(second)
	if fl < sl {
		for i := 0; i < sl-fl; i++ {
			second = second.next
		}
	} else {
		for i := 0; i < fl-sl; i++ {
			first = first.next
		}
	}
	for first != nil && second != nil {
		if first.Value == second.Value {
			return first
		}
		first = first.next
		second = second.next
	}
	return nil
}

// Length calculates the length of the singly linked list by traversing it.
// This operation is runtime O(N).
func Length(head *Elem) int {
	result := 0
	for ; head != nil; head = head.next {
		result++
	}
	return result
}
