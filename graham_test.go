package main

import "testing"

func TestScan(t *testing.T) {
	type args struct {
		points []Point
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{"(0, 0),(3, 0),(1, 1)", args{[]Point{{0, 0}, {3, 0}, {1, 1}}}, []Point{{0, 0}, {3, 0}, {1, 1}}},
		{"(0,0),(3,0),(1,1),(1,3)", args{[]Point{{0, 0}, {3, 0}, {1, 1}, {1, 3}}}, []Point{{0, 0}, {3, 0}, {1, 3}}},
		{"(3,0),(0,0),(1,3),(1,1)", args{[]Point{{3, 0}, {0, 0}, {1, 3}, {1, 1}}}, []Point{{0, 0}, {3, 0}, {1, 3}}},
		{"(0,0),(1,1),(2,3),(3,1),(2,2)", args{[]Point{{0, 0}, {1, 1}, {2, 3}, {3, 1}, {2, 2}}}, []Point{{0, 0}, {3, 1}, {2, 3}}},
		{"(0,0),(1,2),(1,1),(2,4),(3,1),(2,2)", args{[]Point{{0, 0}, {1, 2}, {1, 1}, {2, 4}, {3, 1}, {2, 2}}}, []Point{{0, 0}, {3, 1}, {2, 4}, {1, 2}}},
	}
	equal := func(a, b []Point) bool {
		if len(a) != len(b) {
			return false
		}
		m := make(map[Point]int)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scan(tt.args.points); !equal(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
