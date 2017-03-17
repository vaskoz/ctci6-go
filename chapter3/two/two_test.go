package two

import "testing"

var intLessFunc = func(a, b interface{}) bool {
	as := a.(int)
	bs := b.(int)
	if as < bs {
		return true
	} else {
		return false
	}
}

func TestMinStack(t *testing.T) {
	ms := New(10, intLessFunc)
	ms.Push(10)
	ms.Push(15)
	ms.Push(5)
	ms.Push(20)
	if min := ms.Min(); min != 5 {
		t.Errorf("wrong min")
	}
	ms.Pop() // 20
	ms.Pop() // 5
	if min := ms.Min(); min != 10 {
		t.Errorf("wrong min")
	}
}
