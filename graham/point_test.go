package graham

import (
	"testing"
)

func TestAdd(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	p3 := Add(p1, p2)
	if !equal(p3, NewPoint(4, 6)) {
		t.Errorf("Add failed. Expected (4, 6), got (%d, %d)", p3.X, p3.Y)
	}
}

func TestSub(t *testing.T) {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	p3 := Sub(p1, p2)
	if !equal(p3, NewPoint(-2, -2)) {
		t.Errorf("Sub failed. Expected (-2, -2), got (%d, %d)", p3.X, p3.Y)
	}
}

func TestLe(t *testing.T) {
	type args struct {
		p1 *Point
		p2 *Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"p1 < p2", args{NewPoint(1, 2), NewPoint(3, 4)}, true},
		{"p1 > p2", args{NewPoint(3, 4), NewPoint(1, 2)}, false},
		{"p1.Y == p2.Y, p1 < p2", args{NewPoint(1, 3), NewPoint(3, 3)}, true},
		{"p1.Y == p2.Y, p1 > p2", args{NewPoint(3, 3), NewPoint(1, 3)}, false},
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
		p1 *Point
		p2 *Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"(1,2),(3,4)", args{NewPoint(1, 2), NewPoint(3, 4)}, -2},
		{"(3,4),(1,2)", args{NewPoint(3, 4), NewPoint(1, 2)}, 2},
		{"(1,2),(1,2)", args{NewPoint(1, 2), NewPoint(1, 2)}, 0},
		{"(1,2),(2,1)", args{NewPoint(1, 2), NewPoint(2, 1)}, -3},
		{"(2,1),(1,2)", args{NewPoint(2, 1), NewPoint(1, 2)}, 3},
		{"(1,2),(0,0)", args{NewPoint(1, 2), NewPoint(0, 0)}, 0},
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
		p1 *Point
		p2 *Point
		p3 *Point
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"(1,1),(3,1),(2,2)", args{NewPoint(1, 1), NewPoint(3, 1), NewPoint(2, 2)}, false},
		{"(1,1),(2,2),(3,1)", args{NewPoint(1, 1), NewPoint(2, 2), NewPoint(3, 1)}, true},
		{"(3,1),(2,2),(1,1)", args{NewPoint(3, 1), NewPoint(2, 2), NewPoint(1, 1)}, false},
		{"(1,1),(0,0),(3,1)", args{NewPoint(1, 1), NewPoint(0, 0), NewPoint(3, 1)}, false},
		{"(2,2),(0,0),(3,1)", args{NewPoint(2, 2), NewPoint(0, 0), NewPoint(3, 1)}, false},
		{"(1,1),(2,2),(3,3)", args{NewPoint(1, 1), NewPoint(2, 2), NewPoint(3, 3)}, true}, // 直線上にある場合は時計回りとする
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsClockwise(tt.args.p1, tt.args.p2, tt.args.p3); got != tt.want {
				t.Errorf("IsClockwise() = %v, want %v", got, tt.want)
			}
		})
	}
}
