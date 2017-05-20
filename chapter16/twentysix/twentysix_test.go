package twentysix

import "testing"

var testcases = []struct {
	input  string
	output float64
}{
	{"2*3+5/6*3+15", 23.5},
	{"231.2*3.1+5/6*3+15", 734.22},
	{"1+2+3*6+7", 28.0},
	{"256/2/2/8", 8.0},
	{"256-256", 0.0},
}

func TestCalculate(t *testing.T) {
	t.Parallel()
	for _, c := range testcases {
		if output := Calculate(c.input); output != c.output {
			t.Errorf("Expected %v, got %v", c.output, output)
		}
	}
}

func TestBadCalculateInput(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic on bad input")
		}
	}()
	Calculate("256a-1xf")
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Calculate(c.input)
		}
	}
}
