package eight

import "testing"

var testcases = []struct {
	n       int
	english string
}{
	{0, "zero"},
	{122, "one hundred twenty two"},
	{1012, "one thousand twelve"},
	{3522, "three thousand five hundred twenty two"},
	{1352255476, "one billion three hundred fifty two million two hundred fifty five thousand four hundred seventy six"},
}

func TestEnglishInt(t *testing.T) {
	for _, c := range testcases {
		if english := EnglishInt(c.n); english != c.english {
			t.Errorf("Expected '%v', got '%v'", c.english, english)
		}
	}
}

func BenchmarkEnglishInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			EnglishInt(c.n)
		}
	}
}
