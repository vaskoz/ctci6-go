package five

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}

type stack struct {
	data []interface{}
}

func (s *stack) Pop() interface{} {
	var x interface{}
	x, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return x
}

func (s *stack) Peek() interface{} {
	return s.data[len(s.data)-1]
}

func (s *stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

func New() Stack {
	return &stack{}
}

func Sort(s Stack, lessFunc func(first, second interface{}) bool) {
	var p interface{}
	t := New()
	t.Push(s.Pop())
	for !s.IsEmpty() {
		if lessFunc(s.Peek(), t.Peek()) {
			p = s.Pop()
			count := 0
			for !t.IsEmpty() && lessFunc(p, t.Peek()) {
				s.Push(t.Pop())
				count++
			}
			t.Push(p)
			for i := 0; i < count; i++ {
				t.Push(s.Pop())
			}
		} else {
			t.Push(s.Pop())
		}
	}
	for !t.IsEmpty() {
		s.Push(t.Pop())
	}
}
