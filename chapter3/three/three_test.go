package three

import "testing"

func TestSetOfStacks(t *testing.T) {
	sos := New(2)
	sos.Push(100)
	sos.Push(101)
	sos.Push(102)
	if size := sos.Size(); size != 3 {
		t.Errorf("SetOfStacks Size: expected 3 got %v", size)
	}
	if v, size := sos.Pop().(int), sos.Size(); v != 102 || size != 2 {
		t.Errorf("SetOfStacks Pop: expected 102 got %v and size 2 got %v", v, size)
	}
	if v, size := sos.Pop().(int), sos.Size(); v != 101 || size != 1 {
		t.Errorf("SetOfStacks Pop: expected 101 got %v and size 1 got %v", v, size)
	}
	if v, size := sos.Pop().(int), sos.Size(); v != 100 || size != 0 {
		t.Errorf("SetOfStacks Pop: expected 100 got %v and size 0 got %v", v, size)
	}
}

func TestSetOfStacksPopAt(t *testing.T) {
	sos := New(1)
	sos.Push(100)
	sos.Push(101)
	sos.Push(102)
	if v, size := sos.PopAt(1), sos.Size(); v != 101 || size != 2 {
		t.Errorf("SetOfStacks PopAt 1: expected 101 got %v and size 2 got %v", v, size)
	}
}

func TestSetOfStacksEmptyPop(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when there should be")
		}
	}()
	sos := New(1)
	sos.Pop()
}

func BenchmarkSetOfStacks(b *testing.B) {
	sos := New(2)
	for i := 0; i < b.N; i++ {
		sos.Push(101)
		sos.Push(102)
		sos.Push(103)
		sos.Push(104)
		sos.Push(105)
		for j := 0; j < 5; j++ {
			sos.Pop()
		}
	}
}
