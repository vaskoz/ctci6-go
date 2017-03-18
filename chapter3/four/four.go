package four

type Queue interface {
	Push(v interface{})
	Pop() interface{}
	Size() int
}

type queue struct {
	pushes Stack
	pops   Stack
}

type Stack interface {
	Queue
}

type stack struct {
	data []interface{}
}

func (q *queue) Pop() interface{} {
	for q.pushes.Size() > 0 {
		q.pops.Push(q.pushes.Pop())
	}
	return q.pops.Pop()
}

func (q *queue) Push(v interface{}) {
	for q.pops.Size() > 0 {
		q.pushes.Push(q.pops.Pop())
	}
	q.pushes.Push(v)
}

func (q *queue) Size() int {
	return q.pushes.Size() + q.pops.Size()
}

func (s *stack) Pop() interface{} {
	var x interface{}
	x, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return x
}

func (s *stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *stack) Size() int {
	return len(s.data)
}

func New() Queue {
	return &queue{&stack{}, &stack{}}
}
