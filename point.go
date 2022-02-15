package delaunay

type Point struct {
	X float64
	Y float64
}

func (p *Point) isInTriangleCircumcircle(t Triangle) bool {
	// TODO разобраться и привести формулу в порядок
	var s = ((p.X-t.A.X)*(p.Y-t.C.Y)-(p.X-t.C.X)*(p.Y-t.A.Y))*((t.B.X-t.C.X)*(t.B.X-t.A.X)+(t.B.Y-t.C.Y)*(t.B.Y-t.A.Y)) +
		((p.X-t.A.X)*(p.X-t.C.X)+(p.Y-t.A.Y)*(p.Y-t.C.Y))*((t.B.X-t.C.X)*(t.B.Y-t.A.Y)-(t.B.X-t.A.X)*(t.B.Y-t.C.Y))
	return s >= 0
}