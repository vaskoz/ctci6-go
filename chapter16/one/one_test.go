package one

import "testing"

func generateTestcases() []struct{ x, y, ex, ey int } {
	return []struct{ x, y, ex, ey int }{
		{x: 1, y: 2, ex: 2, ey: 1},
		{x: -1, y: 2, ex: 2, ey: -1},
	}
}

func TestNumberSwap(t *testing.T) {
	t.Parallel()
	testcases := generateTestcases()
	for _, c := range testcases {
		if NumberSwap(&c.x, &c.y); c.x != c.ex && c.y != c.ey {
			t.Errorf("expected %v and %v, but got %v and %v for x & y respectively", c.ex, c.ey, c.x, c.y)
		}
	}
}

func TestNumberSwapGo(t *testing.T) {
	t.Parallel()
	testcases := generateTestcases()
	for _, c := range testcases {
		if NumberSwapGo(&c.x, &c.y); c.x != c.ex && c.y != c.ey {
			t.Errorf("expected %v and %v, but got %v and %v for x & y respectively", c.ex, c.ey, c.x, c.y)
		}
	}
}

func TestNumberSwapTemp(t *testing.T) {
	t.Parallel()
	testcases := generateTestcases()
	for _, c := range testcases {
		if NumberSwapTemp(&c.x, &c.y); c.x != c.ex && c.y != c.ey {
			t.Errorf("expected %v and %v, but got %v and %v for x & y respectively", c.ex, c.ey, c.x, c.y)
		}
	}
}

func BenchmarkNumberSwap(b *testing.B) {
	testcases := generateTestcases()
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			NumberSwap(&c.x, &c.y)
		}
	}
}

func BenchmarkNumberSwapGo(b *testing.B) {
	testcases := generateTestcases()
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			NumberSwapGo(&c.x, &c.y)
		}
	}
}

func BenchmarkNumberSwapTemp(b *testing.B) {
	testcases := generateTestcases()
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			NumberSwapTemp(&c.x, &c.y)
		}
	}
}
