package graham

import "testing"

// 値が同じかどうかを比較する
func TestScan(t *testing.T) {
	type args struct {
		points []*Point
	}
	tests := []struct {
		name string
		args args
		want []*Point
	}{
		{"(0,0),(3,0),(1,1)", args{[]*Point{NewPoint(0, 0), NewPoint(3, 0), NewPoint(1, 1)}}, []*Point{NewPoint(0, 0), NewPoint(3, 0), NewPoint(1, 1)}},
		{"(0,0),(3,0),(1,1),(1,3)", args{[]*Point{NewPoint(0, 0), NewPoint(3, 0), NewPoint(1, 1), NewPoint(1, 3)}}, []*Point{NewPoint(0, 0), NewPoint(3, 0), NewPoint(1, 3)}},
		{"(3,0),(0,0),(1,3),(1,1)", args{[]*Point{NewPoint(3, 0), NewPoint(0, 0), NewPoint(1, 3), NewPoint(1, 1)}}, []*Point{NewPoint(0, 0), NewPoint(3, 0), NewPoint(1, 3)}},
		{"(0,0),(1,1),(2,3),(3,1),(2,2)", args{[]*Point{NewPoint(0, 0), NewPoint(1, 1), NewPoint(2, 3), NewPoint(3, 1), NewPoint(2, 2)}}, []*Point{NewPoint(0, 0), NewPoint(3, 1), NewPoint(2, 3)}},
		{"(0,0),(1,2),(1,1),(2,4),(3,1),(2,2)", args{[]*Point{NewPoint(0, 0), NewPoint(1, 2), NewPoint(1, 1), NewPoint(2, 4), NewPoint(3, 1), NewPoint(2, 2)}}, []*Point{NewPoint(0, 0), NewPoint(3, 1), NewPoint(2, 4), NewPoint(1, 2)}},
	}
	equal := func(a, b []*Point) bool {
		if len(a) != len(b) {
			return false
		}
		m := make(map[Point]int)
		for _, p := range a {
			m[*p]++
		}
		for _, p := range b {
			m[*p]--
		}
		for _, v := range m {
			if v != 0 {
				return false
			}
		}
		return true
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scan(tt.args.points); !equal(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ポインタが同じかどうかを比較する
func TestScanPointer(t *testing.T) {
	// {"(0,0),(3,0),(1,1),(1,3)", args{[]*Point{newPoint(0, 0), newPoint(3, 0), newPoint(1, 1), newPoint(1, 3)}}, []*Point{newPoint(0, 0), newPoint(3, 0), newPoint(1, 3)}}
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 0)
	p3 := NewPoint(1, 1)
	p4 := NewPoint(1, 3)
	points := []*Point{p1, p2, p3, p4}
	convex := Scan(points)
	if len(convex) != 3 {
		t.Errorf("len(convex) = %v, want %v", len(convex), 3)
	}
	equal := func(a, b []*Point) bool {
		if len(a) != len(b) {
			return false
		}
		m := make(map[*Point]int)
		for _, p := range a {
			m[p]++
		}
		for _, p := range b {
			m[p]--
		}
		for _, v := range m {
			if v != 0 {
				return false
			}
		}
		return true
	}
	if !equal(convex, []*Point{p1, p2, p4}) {
		t.Errorf("convex = %v, want %v", convex, []*Point{p1, p2, p4})
	}
}
