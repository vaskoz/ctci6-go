package seven

type Elem struct {
	Value int
	next  *Elem
}

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

func Length(head *Elem) int {
	result := 0
	for ; head != nil; head = head.next {
		result++
	}
	return result
}
