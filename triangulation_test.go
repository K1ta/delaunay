package delaunay

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestTriangulate(t *testing.T) {
	type args struct {
		points []Point
	}
	type testCase struct {
		name string
		args args
	}

	var tests = make([]testCase, 100)
	for i := range tests {
		tests[i] = testCase{
			name: strconv.Itoa(i),
			args: args{
				points: generatePoints(100, -100, 100),
			},
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triangles, err := Triangulate(tt.args.points)
			if err != nil {
				t.Errorf("Triangulate() error = %v", err)
				return
			}
			for _, triangle := range triangles {
				for _, point := range tt.args.points {
					if triangle.hasPoints(point) {
						continue
					}
					if point.isInTriangleCircumcircle(triangle) {
						t.Errorf("Point %s is in triangle %s", point, triangle)
					}
				}
			}
		})
	}
}

func generatePoints(n int, min, max float64) []Point {
	var points = make([]Point, n)
	for i := 0; i < n; i++ {
		x := min + rand.Float64()*(max-min)
		y := min + rand.Float64()*(max-min)
		points[i] = Point{X: x, Y: y}
	}
	return points
}

func Test_isClockwise(t *testing.T) {
	type args struct {
		t Triangle
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Clockwise",
			args: args{
				t: Triangle{
					A: Point{X: 0, Y: 0},
					B: Point{X: 0, Y: 1},
					C: Point{X: 1, Y: 0},
				},
			},
			want: true,
		},
		{
			name: "Clockwise with offset points",
			args: args{
				t: Triangle{
					A: Point{X: 1, Y: 0},
					B: Point{X: 0, Y: 0},
					C: Point{X: 0, Y: 1},
				},
			},
			want: true,
		},
		{
			name: "Not clockwise",
			args: args{
				t: Triangle{
					A: Point{X: 0, Y: 0},
					B: Point{X: 1, Y: 0},
					C: Point{X: 0, Y: 1},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.isClockwise(); got != tt.want {
				t.Errorf("isClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTriangleOfPoints(t *testing.T) {
	type args struct {
		a Point
		b Point
		c Point
	}
	tests := []struct {
		name string
		args args
		want Triangle
	}{
		{
			name: "Create as is",
			args: args{
				a: Point{X: 0, Y: 0},
				b: Point{X: 0, Y: 1},
				c: Point{X: 1, Y: 0},
			},
			want: Triangle{
				A: Point{X: 0, Y: 0},
				B: Point{X: 0, Y: 1},
				C: Point{X: 1, Y: 0},
			},
		},
		{
			name: "Make clockwise",
			args: args{
				a: Point{X: 0, Y: 0},
				b: Point{X: 1, Y: 0},
				c: Point{X: 0, Y: 1},
			},
			want: Triangle{
				A: Point{X: 0, Y: 0},
				B: Point{X: 0, Y: 1},
				C: Point{X: 1, Y: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTriangleFromPoints(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTriangleFromPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
