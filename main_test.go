package delaunay

import (
	"reflect"
	"testing"
)

func TestTriangulate(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name    string
		args    args
		want    []Triangle
		wantErr bool
	}{
		{
			name: "",
			args: args{points: []Point{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 1, Y: 1},
				{X: 1, Y: 0},
			}},
			want: []Triangle{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Triangulate(tt.args.points)
			if (err != nil) != tt.wantErr {
				t.Errorf("Triangulate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Triangulate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCrossProductZ(t *testing.T) {
	type args struct {
		a Edge
		b Edge
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				a: Edge{A: Point{X: 1, Y: 1}, B: Point{X: 4, Y: 5}},
				b: Edge{A: Point{X: -1, Y: 0}, B: Point{X: 6, Y: 8}},
			},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCrossProductZ(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getCrossProductZ() = %v, want %v", got, tt.want)
			}
		})
	}
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
			name: "",
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
			name: "",
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
			if got := isClockwise(tt.args.t); got != tt.want {
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
			if got := getTriangleOfPoints(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTriangleOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
