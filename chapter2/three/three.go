package three

type Elem struct {
	Value interface{}
	next  *Elem
}

func DeleteMiddle(elem *Elem) {
	for elem.next != nil {
		elem.Value = elem.next.Value
		elem = elem.next
	}
}
