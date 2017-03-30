package one

import "testing"

func TestMistake(t *testing.T) {
	if err := Mistake(); err == nil {
		t.Errorf("This loop should never end")
	}
}

func BenchmarkMistake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mistake()
	}
}
