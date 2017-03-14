package six

import "container/list"

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
