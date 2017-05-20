package nine

import "testing"

var testcases = []struct {
	a, b                            int
	add, subtract, multiply, divide int
}{
	{a: 100, b: 100, add: 200, subtract: 0, multiply: 10000, divide: 1},
	{a: 1, b: 100000, add: 100001, subtract: -99999, multiply: 100000, divide: 0},
	{a: 1000000, b: 2, add: 1000002, subtract: 999998, multiply: 2000000, divide: 500000},
	{a: -1, b: 2, add: 1, subtract: -3, multiply: -2, divide: 0},
	{a: -1, b: -2, add: -3, subtract: 1, multiply: 2, divide: 0},
	{a: -4, b: -2, add: -6, subtract: -2, multiply: 8, divide: 2},
	{a: -4, b: 2, add: -2, subtract: -6, multiply: -8, divide: -2},
}

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if add := Add(c.a, c.b); c.add != add {
			t.Errorf("Add: expected %v, but got %v", c.add, add)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if subtract := Subtract(c.a, c.b); c.subtract != subtract {
			t.Errorf("Subtract: expected %v, but got %v", c.subtract, subtract)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if multiply := Multiply(c.a, c.b); c.multiply != multiply {
			t.Errorf("Multiply: expected %v, but got %v", c.multiply, multiply)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if divide := Divide(c.a, c.b); c.divide != divide {
			t.Errorf("Divide: expected %v, but got %v", c.divide, divide)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Add(c.a, c.b)
		}
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Subtract(c.a, c.b)
		}
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Multiply(c.a, c.b)
		}
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Divide(c.a, c.b)
		}
	}
}
