package three

import "testing"

var testcases = []struct {
	l1, l2    Line
	intersect bool
	point     Point
}{
	{
		l1:        Line{start: Point{0.0, 0.0}, end: Point{4.0, 4.0}},
		l2:        Line{start: Point{0.0, 2.0}, end: Point{4.0, 2.0}},
		intersect: true,
		point:     Point{2.0, 2.0},
	},
	{
		l1:        Line{start: Point{0.0, 0.0}, end: Point{0.0, 4.0}},
		l2:        Line{start: Point{1.0, 0.0}, end: Point{1.0, 4.0}},
		intersect: false,
		point:     Point{},
	},
	{
		l1:        Line{start: Point{0.0, 0.0}, end: Point{5.0, 5.0}},
		l2:        Line{start: Point{5.0, 0.0}, end: Point{0.0, 5.0}},
		intersect: true,
		point:     Point{2.5, 2.5},
	},
	{
		l1:        Line{start: Point{0.0, 0.0}, end: Point{0.0, 4.0}},
		l2:        Line{start: Point{0.0, 4.0}, end: Point{0.0, 8.0}},
		intersect: false,
		point:     Point{},
	},
	{
		l1:        Line{start: Point{0.0, 0.0}, end: Point{0.0, 4.0}},
		l2:        Line{start: Point{0.0, 4.01}, end: Point{0.0, 8.0}},
		intersect: false,
		point:     Point{},
	},
}

func TestIntersection(t *testing.T) {
	for _, c := range testcases {
		if point, err := Intersection(c.l1, c.l2); point != c.point || (err == nil) != c.intersect {
			t.Errorf("expected %v and %t, got %v and %t", c.point, c.intersect, point, err == nil)
		}
	}
}

func BenchmarkIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testcases {
			Intersection(c.l1, c.l2)
		}
	}
}
