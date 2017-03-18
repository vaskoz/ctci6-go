package six

import "container/list"

type Cat struct{ name string }
type Dog struct{ name string }

type AnimalShelter interface {
	Enqueue(a interface{})
	DequeueAny() interface{}
	DequeueDog() *Dog
	DequeueCat() *Cat
}

type animalShelter struct {
	arrivals *list.List
}

func (as *animalShelter) Enqueue(a interface{}) {
	as.arrivals.PushBack(a)
}

func (as *animalShelter) DequeueAny() interface{} {
	oldest := as.arrivals.Front()
	as.arrivals.Remove(oldest)
	return oldest.Value
}

func (as *animalShelter) DequeueDog() *Dog {
	for e := as.arrivals.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case *Dog:
			as.arrivals.Remove(e)
			return e.Value.(*Dog)
		}
	}
	return nil
}

func (as *animalShelter) DequeueCat() *Cat {
	for e := as.arrivals.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case *Cat:
			as.arrivals.Remove(e)
			return e.Value.(*Cat)
		}
	}
	return nil
}

func New() AnimalShelter {
	return &animalShelter{list.New()}
}
