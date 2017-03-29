package four

type Listy []int

func (ls *Listy) elementAt(i int) int {
	d := []int(*ls)
	if i > len(d) {
		return -1
	} else {
		return d[i-1]
	}
}

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
