package five

type Elem struct {
	Value int
	next  *Elem
}

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
