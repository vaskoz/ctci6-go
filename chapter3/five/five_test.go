package five

import "testing"

func TestStackSort(t *testing.T) {
	s := New()
	s.Push(101)
	s.Push(107)
	s.Push(106)

	lessFunc := func(first, second interface{}) bool {
		fi := first.(int)
		si := second.(int)
		return fi < si
	}
	Sort(s, lessFunc)
	if v := s.Pop(); v != 101 {
		t.Errorf("Expected 101 got %v", v)
	}
	if v := s.Pop(); v != 106 {
		t.Errorf("Expected 106 got %v", v)
	}
	if v := s.Pop(); v != 107 {
		t.Errorf("Expected 107 got %v", v)
	}
}

func TestStackPeek(t *testing.T) {
	s := New()
	s.Push(103)
	if v, empty := s.Peek(), s.IsEmpty(); v != 103 || empty {
		t.Errorf("Peek should not empty a stack")
	}
	if v, empty := s.Pop(), s.IsEmpty(); v != 103 || !empty {
		t.Errorf("Pop should empty a stack")
	}
}
