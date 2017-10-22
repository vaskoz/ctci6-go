package two

type Elem struct {
	Value interface{}
	next  *Elem
}

func FindKth(head *Elem, k int) *Elem {
	size := ListSize(head)
	skip := size - k
	for i := 0; i < skip; i++ {
		head = head.next
	}
	return head
}

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

func ListSize(head *Elem) int {
	count := 0
	for head.next != nil {
		head = head.next
		count++
	}
	return count
}
