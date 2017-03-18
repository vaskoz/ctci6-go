package six

import "testing"

func TestDequeueAny(t *testing.T) {
	as := New()
	as.Enqueue(&Dog{})
	as.Enqueue(&Cat{})
	if _, ok := as.DequeueAny().(*Dog); !ok {
		t.Errorf("Was supposed to be a Dog")
	}
	if _, ok := as.DequeueAny().(*Cat); !ok {
		t.Errorf("Was supposed to be a Cat")
	}
}

func TestDequeueDogCat(t *testing.T) {
	as := New()
	as.Enqueue(&Cat{"fluffy"})
	as.Enqueue(&Cat{"mittens"})
	as.Enqueue(&Dog{"fido"})
	if d := as.DequeueDog(); d == nil || d.name != "fido" {
		t.Errorf("Was supposed to be a Dog named %v", d)
	}
	if d := as.DequeueDog(); d != nil {
		t.Errorf("there were no more dogs")
	}
	if c := as.DequeueCat(); c == nil || c.name != "fluffy" {
		t.Errorf("Was supposed to be a Cat name %v", c)
	}
	if c := as.DequeueCat(); c == nil || c.name != "mittens" {
		t.Errorf("Was supposed to be a Cat name %v", c)
	}
	if c := as.DequeueCat(); c != nil {
		t.Errorf("there were no more cats")
	}
}
