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

type kStacks struct {
	stacks    int
	data      []interface{}
	positions [][]int
}

func (ks *kStacks) Push(id int, v interface{}) error {
	if id >= ks.stacks || id < 0 {
		log.Panicln("invalid stack id")
	}
	if index := len(ks.data); index == cap(ks.data) {
		return errors.New("stack is full")
	} else {
		ks.positions[id] = append(ks.positions[id], index)
		ks.data = append(ks.data, v)
	}
	return nil
}

func (ks *kStacks) Pop(id int) (interface{}, bool) {
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

func (ks *kStacks) Size(id int) int {
	if id >= ks.stacks || id < 0 {
		log.Panicln("invalid stack id")
	}
	return len(ks.positions[id])
}

func New(stacks, totalCapacity int) *kStacks {
	if stacks < 1 || totalCapacity < 1 {
		log.Panicln("number of stacks and total capacity must be greater than or equal to one.")
	}
	data := make([]interface{}, 0, totalCapacity)
	positions := make([][]int, stacks)
	for i := 0; i < stacks; i++ {
		positions[i] = make([]int, 0, stacks)
	}
	return &kStacks{stacks: stacks, data: data, positions: positions}
}
