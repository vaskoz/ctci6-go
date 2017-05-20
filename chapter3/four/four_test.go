package four

import "testing"

func TestQueue(t *testing.T) {
	t.Parallel()
	q := New()
	q.Push(100)
	q.Push(101)
	q.Push(103)
	if v, size := q.Pop().(int), q.Size(); v != 100 || size != 2 {
		t.Errorf("Queue expected 100 got %v and size should be 1 but got %v", v, size)
	}
	if v, size := q.Pop().(int), q.Size(); v != 101 || size != 1 {
		t.Errorf("Queue expected 101 got %v and size should be 0 but got %v", v, size)
	}
	q.Push(200)
	q.Push(201)
	if v, size := q.Pop().(int), q.Size(); v != 103 || size != 2 {
		t.Errorf("Queue expected 103 got %v and size should be 2 but got %v", v, size)
	}
	if v, size := q.Pop().(int), q.Size(); v != 200 || size != 1 {
		t.Errorf("Queue expected 200 got %v and size should be 1 but got %v", v, size)
	}
	if v, size := q.Pop().(int), q.Size(); v != 201 || size != 0 {
		t.Errorf("Queue expected 201 got %v and size should be 0 but got %v", v, size)
	}
}

func BenchmarkQueue(b *testing.B) {
	q := New()
	for i := 0; i < b.N; i++ {
		q.Push(100)
		q.Push(101)
		q.Pop()
		q.Push(102)
		q.Pop()
		q.Pop()
	}
}
