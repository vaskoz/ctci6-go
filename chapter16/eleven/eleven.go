package eleven

// DivingBoards returns all possible lengths of a diving board that can be constructed.
// You must use exactly k boards for each answer of either shorter or longer board lengths.
func DivingBoards(shorter, longer, k int) []int {
	if shorter == longer {
		return []int{shorter * k}
	}
	var result []int
	for i := 0; i <= k; i++ {
		result = append(result, shorter*(k-i)+longer*i)
	}
	return result
}
