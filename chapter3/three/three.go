package three

import "log"

// SetOfStacks represents a collection of stacks that combine into a larger stack.
type SetOfStacks interface {
	Push(v interface{})
	Pop() interface{}
	PopAt(id int) interface{}
	Size() int
}

type setOfStacks struct {
	capacity int
	sets     [][]interface{}
}

func (sos *setOfStacks) Push(v interface{}) {
	if id := len(sos.sets) - 1; id < 0 || len(sos.sets[id]) == sos.capacity {
		sos.sets = append(sos.sets, []interface{}{v})
	} else {
		sos.sets[id] = append(sos.sets[id], v)
	}
}

func (sos *setOfStacks) Pop() interface{} {
	id := len(sos.sets) - 1
	return sos.PopAt(id)
}

func (sos *setOfStacks) PopAt(id int) interface{} {
	if id < 0 || id >= len(sos.sets) {
		log.Panicln("no stack at this position")
	}
	popInd := len(sos.sets[id]) - 1
	result := sos.sets[id][popInd]
	sos.sets[id] = sos.sets[id][:popInd]
	if len(sos.sets[id]) == 0 {
		sos.sets = append(sos.sets[:id], sos.sets[id+1:]...)
	}
	return result
}

func (sos *setOfStacks) Size() int {
	size := 0
	for _, stack := range sos.sets {
		size += len(stack)
	}
	return size
}

// New returns a SetOfStacks where each stack is of a maximum specified capacity.
func New(capacity int) SetOfStacks {
	return &setOfStacks{capacity: capacity}
}
