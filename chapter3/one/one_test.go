package one

import "testing"

func TestBasicThreeStacks(t *testing.T) {
	threeStacks := New(3, 2)
	threeStacks.Push(0, "foo")
	threeStacks.Push(1, "baz")
	if size := threeStacks.Size(0); size != 1 {
		t.Errorf("incorrect stack size")
	}
	if size := threeStacks.Size(1); size != 1 {
		t.Errorf("incorrect stack size")
	}
	if size := threeStacks.Size(2); size != 0 {
		t.Errorf("incorrect stack size")
	}
}

func TestBasicFoundPop(t *testing.T) {
	s := New(3, 1)
	s.Push(2, "bar")
	if v, found := s.Pop(2); found == false || v != "bar" {
		t.Errorf("Pop returned %v and %v but expected bar and nil\n", v, found)
	}
}

func TestBasicNotFoundPop(t *testing.T) {
	s := New(3, 1)
	if v, found := s.Pop(2); found {
		t.Errorf("Pop returned %v and %v but expected not found\n", v, found)
	}
}

func TestExceedingCapacity(t *testing.T) {
	s := New(3, 2)
	s.Push(0, "foo")
	s.Push(1, "baz")
	if err := s.Push(0, "bar"); err == nil {
		t.Errorf("should return an error when capacity is reached")
	}
}

func TestBadStackCreation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when there should be")
		}
	}()
	New(0, 1)
}

func TestBadCapacityCreation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when there should be")
		}
	}()
	New(1, 0)
}

func TestBadIdSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when there should be")
		}
	}()
	s := New(1, 1)
	s.Size(1)
}

func TestBadIdPush(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when there should be")
		}
	}()
	s := New(1, 1)
	s.Push(1, "foo")
}

func TestBadIdPop(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("No panic when there should be")
		}
	}()
	s := New(1, 1)
	s.Pop(1)
}
