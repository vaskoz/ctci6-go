package two

// Node represents a node in a binary tree.
type Node struct {
	Value       int
	left, right *Node
}

// CreateBinaryTree builds a binary search tree given a sorted array of integers.
func CreateBinaryTree(data []int, start, end int) *Node {
	if start >= end {
		return nil
	}
	mid := start + ((end - start) / 2)
	result := &Node{Value: data[mid]}
	result.left = CreateBinaryTree(data, start, mid)
	result.right = CreateBinaryTree(data, mid+1, end)
	return result
}
