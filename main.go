package delaunay

import (
	"errors"
	"fmt"

	"github.com/go-gl/mathgl/mgl64"
)

type (
	Point struct {
		X float64
		Y float64
	}

	Triangle struct {
		A Point
		B Point
		C Point
	}

	Edge struct {
		A Point
		B Point
	}
)

func Triangulate(points []Point) (_ []Triangle, err error) {
	if len(points) < 3 {
		return nil, errors.New("not enough points")
	}

	var left, top, right, bottom = getBounds(points)

	fmt.Println(left, top, right, bottom)

	var (
		leftBottom  = Point{X: left, Y: bottom}
		leftTop     = Point{X: left, Y: top}
		rightTop    = Point{X: right, Y: top}
		rightBottom = Point{X: right, Y: bottom}
	)

	var set = triangleSet{
		triangles: map[Triangle]struct{}{
			getTriangleOfPoints(rightBottom, leftBottom, leftTop): {},
			getTriangleOfPoints(leftTop, rightTop, rightBottom):   {},
		},
	}

	for _, p := range points {
		var badTriangles = make([]Triangle, 0)
		for t := range set.triangles {
			if isPointInTriangleCircumcircle(p, t) {
				set.Remove(t)
				badTriangles = append(badTriangles, t)
			}
		}
		var newTriangles = combineTrianglesAndPoint(badTriangles, p)
		for _, newTriangle := range newTriangles {
			set.Add(newTriangle)
		}
	}

	for t := range set.triangles {
		if triangleHasPoints(t, leftBottom, leftTop, rightTop, rightBottom) {
			set.Remove(t)
		}
	}

	return set.ToSlice(), nil
}

func getTriangleOfPoints(a, b, c Point) Triangle {
	var t = Triangle{A: a, B: b, C: c}
	if isClockwise(t) {
		return t
	}

	t = Triangle{A: a, B: c, C: b}
	return t
}

func isClockwise(t Triangle) bool {
	var (
		v1 = getVec3OfPoints(t.A, t.B)
		v2 = getVec3OfPoints(t.A, t.C)
	)
	return v1.Cross(v2).Z() < 0
}

func getVec3OfPoints(from, to Point) mgl64.Vec3 {
	return mgl64.Vec3{to.X - from.X, to.Y - from.Y, 0}
}

func getBounds(points []Point) (left, top, right, bottom float64) {
	left = points[0].X
	bottom = points[0].Y
	right = points[0].X
	top = points[0].Y
	for _, p := range points {
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
	return left - 1, top + 1, right + 1, bottom - 1
}

func isPointInTriangleCircumcircle(p Point, t Triangle) bool {
	// TODO разобраться и привести формулу в порядок
	var s = ((p.X-t.A.X)*(p.Y-t.C.Y)-(p.X-t.C.X)*(p.Y-t.A.Y))*((t.B.X-t.C.X)*(t.B.X-t.A.X)+(t.B.Y-t.C.Y)*(t.B.Y-t.A.Y)) +
		((p.X-t.A.X)*(p.X-t.C.X)+(p.Y-t.A.Y)*(p.Y-t.C.Y))*((t.B.X-t.C.X)*(t.B.Y-t.A.Y)-(t.B.X-t.A.X)*(t.B.Y-t.C.Y))
	return s >= 0
}

func combineTrianglesAndPoint(triangles []Triangle, p Point) []Triangle {
	var edges = make(map[Edge]struct{})
	for _, t := range triangles {
		addEdgePointsToMap(edges, t.A, t.B)
		addEdgePointsToMap(edges, t.B, t.C)
		addEdgePointsToMap(edges, t.C, t.A)
	}

	var res = make([]Triangle, 0, len(triangles))
	for e := range edges {
		res = append(res, getTriangleOfPoints(e.A, e.B, p))
	}

	return res
}

func addEdgePointsToMap(m map[Edge]struct{}, a, b Point) {
	var edge = getEdgeFromTwoPoints(a, b)
	if _, ok := m[edge]; ok {
		delete(m, edge)
	} else {
		m[edge] = struct{}{}
	}
}

func (t *Triangle) getEdges() [3]Edge {
	return [3]Edge{
		getEdgeFromTwoPoints(t.A, t.B),
		getEdgeFromTwoPoints(t.B, t.C),
		getEdgeFromTwoPoints(t.C, t.A),
	}
}

func getEdgeFromTwoPoints(a, b Point) Edge {
	if a.X <= b.X {
		return Edge{a, b}
	}
	return Edge{b, a}
}

func triangleHasPoints(t Triangle, points ...Point) bool {
	for _, p := range points {
		if t.A == p || t.B == p || t.C == p {
			return true
		}
	}
	return false
}
