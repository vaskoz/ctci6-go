package twentyfive

import "testing"

func TestLRU(t *testing.T) {
	lru := New(3)
	lru.Set(2, "foo")
	lru.Set(4, "bar")
	lru.Set(6, "baz")
	lru.Set(1, "oops")
	if v := lru.Get(2); v != nil {
		t.Errorf("should have dropped least recently used, but got %v", v)
	}
	if v := lru.Get(4); v != "bar" {
		t.Errorf("should have dropped least recently used, but got %v", v)
	}
	// LRU now [4, 1, 6]
	lru.Set(6, "hi")
	lru.Set(9, "there")
	lru.Set(10, "code")
	// LRU now [10, 9, 6]
	if v := lru.Get(4); v != nil {
		t.Errorf("should have dropped least recently used, but got %v", v)
	}
	if v := lru.Get(2); v != nil {
		t.Errorf("should have dropped least recently used, but got %v", v)
	}
	if v := lru.Get(1); v != nil {
		t.Errorf("should have dropped least recently used, but got %v", v)
	}
	if v := lru.Get(10); v != "code" {
		t.Errorf("value should be in lru %v", v)
	}
	if v := lru.Get(9); v != "there" {
		t.Errorf("value should be in lru %v", v)
	}
	if v := lru.Get(6); v != "hi" {
		t.Errorf("value should be in lru %v", v)
	}
}

func BenchmarkLRU(b *testing.B) {
	lru := New(5)
	for i := 0; i < b.N; i++ {
		lru.Set(1, "a")
		lru.Set(2, "b")
		lru.Set(3, "c")
		lru.Set(4, "d")
		lru.Set(5, "e")
	}
}
