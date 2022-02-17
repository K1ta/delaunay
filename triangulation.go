package delaunay

import (
	"errors"
)

func Triangulate(points []Point) (_ []Triangle, err error) {
	if len(points) < 3 {
		return nil, errors.New("not enough points")
	}

	var set = createTriangleSetWithSuperposition(points)

	for _, p := range points {
		addPointToTriangleSet(p, set)
	}

	removeSuperpositionFromTriangleSet(set)

	return set.ToSlice(), nil
}

func createTriangleSetWithSuperposition(points []Point) triangleSet {
	var leftBottom, leftTop, rightTop, rightBottom = getSuperpositionBounds(points)

	var set = triangleSet{
		triangles: map[Triangle]struct{}{
			newTriangleFromPoints(rightBottom, leftBottom, leftTop): {},
			newTriangleFromPoints(leftTop, rightTop, rightBottom):   {},
		},
		leftBottom:  leftBottom,
		leftTop:     leftTop,
		rightTop:    rightTop,
		rightBottom: leftBottom,
	}

	return set
}

func getSuperpositionBounds(points []Point) (leftBottom, leftTop, rightTop, rightBottom Point) {
	var (
		left   = points[0].X
		bottom = points[0].Y
		right  = points[0].X
		top    = points[0].Y
	)
	for _, p := range points[1:] {
		if p.X < left {
			left = p.X
		}
		if p.Y < bottom {
			bottom = p.Y
		}
		if p.X > right {
			right = p.X
		}
		if p.Y > top {
			top = p.Y
		}
	}
	left = left - 1
	top = top + 1
	right = right + 1
	bottom = bottom - 1

	leftBottom = Point{X: left, Y: bottom}
	leftTop = Point{X: left, Y: top}
	rightTop = Point{X: right, Y: top}
	rightBottom = Point{X: right, Y: bottom}
	return leftBottom, leftTop, rightTop, rightBottom
}

func addPointToTriangleSet(p Point, set triangleSet) {
	var edgesForNewTriangles = make(map[Edge]struct{})
	for t := range set.triangles {
		if p.isInTriangleCircumcircle(t) {
			set.Remove(t)
			addEdgeToMap(edgesForNewTriangles, newEdgeFromTwoPoints(t.A, t.B))
			addEdgeToMap(edgesForNewTriangles, newEdgeFromTwoPoints(t.B, t.C))
			addEdgeToMap(edgesForNewTriangles, newEdgeFromTwoPoints(t.C, t.A))
		}
	}

	var newTriangles = make([]Triangle, 0)
	for e := range edgesForNewTriangles {
		newTriangles = append(newTriangles, newTriangleFromPoints(e.A, e.B, p))
	}

	for _, newTriangle := range newTriangles {
		set.Add(newTriangle)
	}
}

func addEdgeToMap(m map[Edge]struct{}, e Edge) {
	if _, ok := m[e]; ok {
		delete(m, e)
	} else {
		m[e] = struct{}{}
	}
}

func removeSuperpositionFromTriangleSet(set triangleSet) {
	for t := range set.triangles {
		if t.hasPoints(set.leftBottom, set.leftTop, set.rightTop, set.rightBottom) {
			set.Remove(t)
		}
	}
}
