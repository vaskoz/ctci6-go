package five

// OneAway checks if two strings are at most one edit away from each other.
func OneAway(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if diff := len(s1) - len(s2); diff > 1 || diff < -1 {
		return false
	}
	m := map[rune]int{}
	for _, r := range s1 {
		m[r]++
	}
	missing := 0
	for _, r := range s2 {
		if v, found := m[r]; v > 1 {
			m[r]--
		} else if found {
			delete(m, r)
		} else {
			missing++
		}
	}
	return missing+len(m) < 3
}
