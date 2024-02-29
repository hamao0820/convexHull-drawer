package main

import "math"

type Point struct {
	X, Y int
}

func Add(p1, p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func Sub(p1, p2 Point) Point {
	return Point{p1.X - p2.X, p1.Y - p2.Y}
}

// p1はp2より左下にあるか
func Le(p1, p2 Point) bool {
	if p1.Y < p2.Y {
		return true
	}
	if p1.Y > p2.Y {
		return false
	}
	return p1.X < p2.X
}

func Cross(p1, p2 Point) int {
	return p1.X*p2.Y - p1.Y*p2.X
}

// p1->p2->p3が時計回りか
// p1->p2ベクトルとp1->p3ベクトルの外積が負なら時計回り
func IsClockwise(p1, p2, p3 Point) bool {
	return Cross(Sub(p2, p1), Sub(p3, p1)) < 0
}

// p1からp2(p1<p2)への角度
func ElevationAngle(p1, p2 Point) float64 {
	if Le(p2, p1) {
		p1, p2 = p2, p1
	}
	return math.Atan2(float64(p2.Y-p1.Y), float64(p2.X-p1.X))
}
