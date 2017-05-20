package two

import (
	"log"
	"testing"
)

var testcases = []struct {
	in []int
}{
	{[]int{1, 2, 3, 4, 5}},
	{[]int{2, 4, 5, 6, 7, 8, 10}},
}

func TestCreateBinaryTree(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		head := CreateBinaryTree(c.in, 0, len(c.in))
		preOrder(head)
	}
}

func preOrder(head *Node) {
	if head == nil {
		return
	}
	log.Println(head)
	preOrder(head.left)
	preOrder(head.right)
}
