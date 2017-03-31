package one

import "testing"

func TestMistake(t *testing.T) {
	if err := Mistake(false); err != nil {
		t.Errorf("This won't overflow so no error")
	}
	if err := Mistake(true); err == nil {
		t.Errorf("This should have overflowed so it's an error")
	}
}

func BenchmarkMistake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mistake(true)
	}
}
