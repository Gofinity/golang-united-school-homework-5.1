package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	side := int(s.a)
	return Point{x: s.start.x + side, y: s.start.y + side}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
