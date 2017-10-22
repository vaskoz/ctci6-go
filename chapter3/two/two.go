package two

type MinStack interface {
	Push(v interface{})
	Pop() interface{}
	Min() interface{}
}

type minStack struct {
	data     []node
	lessFunc func(a, b interface{}) bool
}

func (ms *minStack) Pop() interface{} {
	var n node
	n, ms.data = ms.data[len(ms.data)-1], ms.data[:len(ms.data)-1]
	return n.value
}

func (ms *minStack) Min() interface{} {
	if len(ms.data) < 1 {
		return nil
	}
	n := ms.data[len(ms.data)-1]
	return ms.data[n.minIndex].value
}

func (ms *minStack) Push(v interface{}) {
	if min, index := ms.Min(), len(ms.data); min == nil {
		ms.data = append(ms.data, node{value: v, minIndex: index})
	} else if ms.lessFunc(v, min) {
		ms.data = append(ms.data, node{value: v, minIndex: index})
	} else {
		index := ms.data[len(ms.data)-1].minIndex
		ms.data = append(ms.data, node{value: v, minIndex: index})
	}
}

type node struct {
	minIndex int
	value    interface{}
}

func New(capacity int, lessFunc func(a, b interface{}) bool) MinStack {
	data := make([]node, 0, capacity)
	return &minStack{data: data, lessFunc: lessFunc}
}
