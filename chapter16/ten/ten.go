package ten

// MostLiveYear returns the year with the most number of people alive.
// Input is a slice of births and a slice of deaths.
func MostLiveYear(births, deaths []int) int {
	years := make([]int, 1001)
	for i := 0; i < len(births); i++ {
		for y := births[i] - 1900; y <= deaths[i]-1900; y++ {
			years[y]++
		}
	}
	largestYear, largestCount := 0, 0
	for i := 0; i < len(years); i++ {
		if years[i] > largestCount {
			largestCount = years[i]
			largestYear = i + 1900
		}
	}
	return largestYear
}
