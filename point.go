package main

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
