package eleven

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
