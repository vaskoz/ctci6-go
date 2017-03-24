package two

type Node struct {
	Value       int
	left, right *Node
}

func CreateBinaryTree(data []int, start, end int) *Node {
	if start >= end {
		return nil
	} else {
		mid := start + ((end - start) / 2)
		result := &Node{Value: data[mid]}
		result.left = CreateBinaryTree(data, start, mid)
		result.right = CreateBinaryTree(data, mid+1, end)
		return result
	}
}
