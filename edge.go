package delaunay

type Edge struct {
	A Point
	B Point
}

func newEdgeFromTwoPoints(a, b Point) Edge {
	if a.X <= b.X {
		return Edge{a, b}
	}
	return Edge{b, a}
}
