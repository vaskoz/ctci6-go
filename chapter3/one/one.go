package one

import (
	"errors"
	"log"
)

type Stacks interface {
	Push(id int, v interface{}) error
	Pop(id int) (interface{}, bool)
	Size(id int) int
}

type stacks struct {
	stacks    int
	data      []interface{}
	positions [][]int
}

func (ks *stacks) Push(id int, v interface{}) error {
	if id >= ks.stacks || id < 0 {
		log.Panicln("invalid stack id")
	}
	if index := len(ks.data); index < cap(ks.data) {
		ks.positions[id] = append(ks.positions[id], index)
		ks.data = append(ks.data, v)
		return nil
	}
	return errors.New("stack is full")
}

func (ks *stacks) Pop(id int) (interface{}, bool) {
	if id >= ks.stacks || id < 0 {
		log.Panicln("invalid stack id")
	}
	a := ks.positions[id]
	if len(a) == 0 {
		return nil, false
	}
	index := a[len(a)-1]
	ks.positions[id] = a[:len(a)-1]
	return ks.data[index], true
}

func (ks *stacks) Size(id int) int {
	if id >= ks.stacks || id < 0 {
		log.Panicln("invalid stack id")
	}
	return len(ks.positions[id])
}

func New(numStacks, totalCapacity int) Stacks {
	if numStacks < 1 || totalCapacity < 1 {
		log.Panicln("number of stacks and total capacity must be greater than or equal to one.")
	}
	data := make([]interface{}, 0, totalCapacity)
	positions := make([][]int, numStacks)
	for i := 0; i < numStacks; i++ {
		positions[i] = make([]int, 0, numStacks)
	}
	return &stacks{stacks: numStacks, data: data, positions: positions}
}
