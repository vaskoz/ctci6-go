package five

// SparseStringSearch does a binary search for a search value.
// The data is expected to be sorted lexically (natural ordering).
func SparseStringSearch(data []string, search string) int {
	l, r := 0, len(data)-1
	ll, lr := l, r
	for l <= r {
		mid := (l + r) / 2
		if val := data[mid]; val == search {
			return mid
		} else if val == "" {
			for ; data[mid] == ""; mid++ {
			}
			if nextval := data[mid]; nextval == search {
				return mid
			} else if search < nextval {
				r = mid - 1
			} else {
				l = mid + 1
			}
			if r == lr && l == ll {
				return -1
			}
		} else if search < val {
			r = mid - 1
		} else {
			l = mid + 1
		}
		ll = l
		lr = r
	}
	return -1
}
