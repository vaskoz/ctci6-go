package twentyfive

import (
	"container/list"
)

// LRU is a least recently used cache algorithm.
type LRU struct {
	data     map[interface{}]*list.Element
	list     *list.List
	capacity int
}

type tuple struct {
	key, value interface{}
}

// New returns an LRU of the specified capacity.
func New(capacity int) *LRU {
	return &LRU{
		data:     make(map[interface{}]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

// Get retrieves a value by the key.
// By accessing the value it will persist longer than the other values in the cache.
func (l *LRU) Get(key interface{}) interface{} {
	if e, found := l.data[key]; found {
		l.list.MoveToFront(e)
		return e.Value.(*tuple).value
	}
	return nil
}

// Set associates a key and a value in the LRU.
func (l *LRU) Set(key, value interface{}) {
	if e, found := l.data[key]; !found && len(l.data) == l.capacity {
		target := l.list.Back()
		delete(l.data, target.Value.(*tuple).key)
		l.list.Remove(target)
		l.data[key] = l.list.PushFront(&tuple{key: key, value: value})
	} else if !found {
		l.data[key] = l.list.PushFront(&tuple{key: key, value: value})
	} else {
		l.list.MoveToFront(e)
		e.Value.(*tuple).value = value
	}
}
