package three

func FindRotatedIndex(data []int, search int) int {
	l, r := 0, len(data)-1
	for l <= r {
		mid := (l + r) / 2
		if data[mid] == search {
			return mid
		} else if data[mid] < search {
			if search > data[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if data[mid] < data[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
