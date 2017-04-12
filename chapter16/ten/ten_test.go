package ten

import "testing"

var testcases = []struct {
	births, deaths []int
	mostLiveYear   int
}{
	{
		births:       []int{1901, 1903, 1905, 1907, 1909, 1945, 1944, 1945, 1999},
		deaths:       []int{1902, 1904, 1906, 1908, 1910, 1945, 1945, 1946, 2000},
		mostLiveYear: 1945,
	},
}

func TestMostLiveYear(t *testing.T) {
	for _, c := range testcases {
		if mostLiveYear := MostLiveYear(c.births, c.deaths); mostLiveYear != c.mostLiveYear {
			t.Errorf("Expected %v, but got %v", c.mostLiveYear, mostLiveYear)
		}
	}
}

func BenchmarkMostLiveYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			MostLiveYear(c.births, c.deaths)
		}
	}
}
