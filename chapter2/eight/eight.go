package eight

type Elem struct {
	Value interface{}
	next  *Elem
}

func DetectLoop(head *Elem) *Elem {
	seen := map[*Elem]struct{}{}
	for head != nil {
		if _, found := seen[head]; found {
			return head
		} else {
			seen[head] = struct{}{}
		}
		head = head.next
	}
	return nil
}
