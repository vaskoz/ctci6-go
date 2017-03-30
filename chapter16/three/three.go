package three

import "errors"

type Point struct {
	x, y float64
}

type Line struct {
	start, end Point
}

func Intersection(l1, l2 Line) (Point, error) {
	x1 := l1.start.x
	x2 := l1.end.x
	y1 := l1.start.y
	y2 := l1.end.y
	x3 := l2.start.x
	x4 := l2.end.x
	y3 := l2.start.y
	y4 := l2.end.y
	det := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)
	if det == 0 {
		return Point{}, errors.New("parallel lines")
	} else {
		x := ((x1*y2-y1*x2)*(x3-x4) - (x1-x2)*(x3*y4-y3*x4)) / det
		y := ((x1*y2-y1*x2)*(y3-y4) - (y1-y2)*(x3*y4-y3*x4)) / det
		return Point{x, y}, nil
	}
}
