package delaunay

import "github.com/go-gl/mathgl/mgl64"

type Triangle struct {
	A Point
	B Point
	C Point
}

func newTriangleFromPoints(a, b, c Point) Triangle {
	var t = Triangle{A: a, B: b, C: c}
	if t.isClockwise() {
		return t
	}

	t = Triangle{A: a, B: c, C: b}
	return t
}

func (t *Triangle) isClockwise() bool {
	var (
		v1 = getVec3OfPoints(t.A, t.B)
		v2 = getVec3OfPoints(t.A, t.C)
	)
	return v1.Cross(v2).Z() < 0
}

func getVec3OfPoints(from, to Point) mgl64.Vec3 {
	return mgl64.Vec3{to.X - from.X, to.Y - from.Y, 0}
}

func (t *Triangle) hasPoints(points ...Point) bool {
	for _, p := range points {
		if t.A == p || t.B == p || t.C == p {
			return true
		}
	}
	return false
}
