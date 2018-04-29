package six

import "container/list"

// IsPalindrome checks if a linked list is a palindrome.
func IsPalindrome(head *list.List) bool {
	front := head.Front()
	back := head.Back()
	for front != nil && back != nil {
		if front.Value != back.Value {
			return false
		}
		front = front.Next()
		back = back.Prev()
	}
	return true
}
