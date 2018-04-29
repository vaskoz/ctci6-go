package two

// Elem represents a node in a singly linked list.
type Elem struct {
	Value interface{}
	next  *Elem
}

// FindKth finds the k-th element from the end of the singly linked list.
// This uses the ListSize function.
func FindKth(head *Elem, k int) *Elem {
	size := ListSize(head)
	skip := size - k
	for i := 0; i < skip; i++ {
		head = head.next
	}
	return head
}

// FindKth2Ptr finds the k-th element from the end of the singly linked list too.
// However, it uses another pointer to run ahead k positions and then moves them both
// in lock-step until the front pointer reaches the end. Then the following pointer
// is pointing to the answer.
func FindKth2Ptr(head *Elem, k int) *Elem {
	ptr := head
	for i := 0; i < k; i++ {
		head = head.next
	}
	for head.next != nil {
		head = head.next
		ptr = ptr.next
	}
	return ptr
}

// ListSize returns the size of the given singly linked list.
// Runtime O(N) since it has to traverse the entire list.
func ListSize(head *Elem) int {
	count := 0
	for head.next != nil {
		head = head.next
		count++
	}
	return count
}
