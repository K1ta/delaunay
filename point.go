package delaunay

import "fmt"

type Point struct {
	X float64
	Y float64
}

func (p *Point) isInTriangleCircumcircle(t Triangle) bool {
	var s = ((p.X-t.A.X)*(p.Y-t.C.Y)-(p.X-t.C.X)*(p.Y-t.A.Y))*((t.B.X-t.C.X)*(t.B.X-t.A.X)+(t.B.Y-t.C.Y)*(t.B.Y-t.A.Y)) +
		((p.X-t.A.X)*(p.X-t.C.X)+(p.Y-t.A.Y)*(p.Y-t.C.Y))*((t.B.X-t.C.X)*(t.B.Y-t.A.Y)-(t.B.X-t.A.X)*(t.B.Y-t.C.Y))
	return s > 0
}

func (p Point) String() string {
	return fmt.Sprintf("[%f, %f]", p.X, p.Y)
}
