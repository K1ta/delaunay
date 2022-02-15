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
					A: Point{X: 1, Y: 0},
					B: Point{X: 0, Y: 0},
					C: Point{X: 0, Y: 1},
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
			if got := newTriangleFromPoints(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTriangleFromPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
