package eight

// Elem represents a node in a singly linked list.
type Elem struct {
	Value interface{}
	next  *Elem
}

// DetectLoop returns a pointer to the element that causes the loop in the singly linked list.
// Runtime is O(N) as it need to traverse the entire list.
// Space is also O(N) to store seen nodes.
func DetectLoop(head *Elem) *Elem {
	seen := map[*Elem]struct{}{}
	for head != nil {
		if _, found := seen[head]; found {
			return head
		}
		seen[head] = struct{}{}
		head = head.next
	}
	return nil
}
