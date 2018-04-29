package five

// Elem represents a node in a singly linked list.
type Elem struct {
	Value int
	next  *Elem
}

// SumReverseList takes two singly linked lists of ints.
// It sums their values assuming the first value of each list is the 1's value.
// And the 2nd value in each list is the 10's value and so on.
func SumReverseList(first, second *Elem) int {
	result := 0
	carrytheone := false
	place := 1
	for first != nil || second != nil {
		var total int
		if first != nil && second != nil {
			total = first.Value + second.Value
			first = first.next
			second = second.next
		} else if first != nil {
			total = first.Value
			first = first.next
		} else if second != nil {
			total = second.Value
			second = second.next
		}
		if carrytheone {
			total++
			carrytheone = false
		}
		if total >= 10 {
			total -= 10
			carrytheone = true
		}
		result += place * total
		place *= 10
	}
	if carrytheone {
		result += place * 1
	}
	return result
}
