package four

// Listy is a slice of ints
type Listy []int

func (ls *Listy) elementAt(i int) int {
	d := []int(*ls)
	if i > len(d) {
		return -1
	}
	return d[i-1]
}

// SearchListy does a binary search to find the index of a search value.
// The returned index is NOT zero-based.
// -1 is returned if the search value does not exist in the list.
func SearchListy(ls *Listy, search int) int {
	index := 0
	for index = 1; ls.elementAt(index) != -1; index *= 2 {
	}
	l, r := 1, index
	for l <= r {
		mid := (l + r) / 2
		if val := ls.elementAt(mid); val == -1 || search < val {
			r = mid - 1
		} else if val == search {
			return mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
