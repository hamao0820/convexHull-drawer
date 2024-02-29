package graham

type Pointer interface {
	X() int
	Y() int
	SetX(int)
	SetY(int)
	Add(Pointer, Pointer)
	Sub(Pointer, Pointer)
}

type Point struct {
	x, y int
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) SetY(y int) {
	p.y = y
}

func (p *Point) Add(p1, p2 Pointer) {
	p.x = p1.X() + p2.X()
	p.y = p1.Y() + p2.Y()
}

func (p *Point) Sub(p1, p2 Pointer) {
	p.x = p1.X() - p2.X()
	p.y = p1.Y() - p2.Y()
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func equal[T Pointer](p, p2 T) bool {
	return p.X() == p2.X() && p.Y() == p2.Y()
}

// p1はp2より左下にあるか
func Le[T Pointer](p1, p2 T) bool {
	if p1.Y() < p2.Y() {
		return true
	}
	if p1.Y() > p2.Y() {
		return false
	}
	return p1.X() < p2.X()
}

func Cross[T Pointer](p1, p2 T) int {
	return p1.X()*p2.Y() - p1.Y()*p2.X()
}

// p1->p2->p3が時計回りか
// p1->p2ベクトルとp1->p3ベクトルの外積が0以下なら時計回り
func IsClockwise[T Pointer](p1, p2, p3 T) bool {
	sub := NewPoint(0, 0)
	sub.Sub(p2, p1)
	sub2 := NewPoint(0, 0)
	sub2.Sub(p3, p1)
	return Cross(sub, sub2) <= 0
}
