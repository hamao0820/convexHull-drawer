package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := Add(p1, p2)
	if p3.X != 4 || p3.Y != 6 {
		t.Errorf("Add failed. Expected (4, 6), got (%d, %d)", p3.X, p3.Y)
	}
}

func TestSub(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := Sub(p1, p2)
	if p3.X != -2 || p3.Y != -2 {
		t.Errorf("Sub failed. Expected (-2, -2), got (%d, %d)", p3.X, p3.Y)
	}
}

func TestLe(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"p1 < p2", args{Point{1, 2}, Point{3, 4}}, true},
		{"p1 > p2", args{Point{3, 4}, Point{1, 2}}, false},
		{"p1.Y == p2.Y, p1 < p2", args{Point{1, 3}, Point{3, 3}}, true},
		{"p1.Y == p2.Y, p1 > p2", args{Point{3, 3}, Point{1, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Le(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("Le() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCross(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Cross product", args{Point{1, 2}, Point{3, 4}}, -2},
		{"Cross product", args{Point{3, 4}, Point{1, 2}}, 2},
		{"Cross product", args{Point{1, 2}, Point{1, 2}}, 0},
		{"Cross product", args{Point{1, 2}, Point{2, 1}}, -3},
		{"Cross product", args{Point{2, 1}, Point{1, 2}}, 3},
		{"Cross product", args{Point{1, 2}, Point{0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cross(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsClockwise(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
		p3 Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"(1,1),(3,1),(2,2)", args{Point{1, 1}, Point{3, 1}, Point{2, 2}}, false},
		{"(1,1),(2,2),(3,1)", args{Point{1, 1}, Point{2, 2}, Point{3, 1}}, true},
		{"(3,1),(2,2),(1,1)", args{Point{3, 1}, Point{2, 2}, Point{1, 1}}, false},
		{"(1,1),(0,0),(3,1)", args{Point{1, 1}, Point{0, 0}, Point{3, 1}}, false},
		{"(2,2),(0,0),(3,1)", args{Point{2, 2}, Point{0, 0}, Point{3, 1}}, false},
		{"(1,1),(2,2),(3,3)", args{Point{1, 1}, Point{2, 2}, Point{3, 3}}, true}, // 直線上にある場合は時計回りとする
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsClockwise(tt.args.p1, tt.args.p2, tt.args.p3); got != tt.want {
				t.Errorf("IsClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElevationAngle(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"(1,2),(3,4)", args{Point{1, 2}, Point{3, 4}}, math.Pi / 4},
		{"(3,4),(1,2)", args{Point{3, 4}, Point{1, 2}}, math.Pi / 4},
		{"(1,2),(1,2)", args{Point{1, 2}, Point{1, 2}}, 0},
		{"(1,2),(3,2)", args{Point{1, 2}, Point{3, 2}}, 0},
		{"(0,0),(1,0)", args{Point{0, 0}, Point{1, 0}}, 0},
		{"(0,0),(0,1)", args{Point{0, 0}, Point{0, 1}}, math.Pi / 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElevationAngle(tt.args.p1, tt.args.p2); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("ElevationAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}
