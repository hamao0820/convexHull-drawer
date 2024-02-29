package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hamao0820/convexHull-drawer/graham"
)

type Plot struct {
	p        graham.Point
	c        color.Color
	isConvex bool
	id       int
	r        int
}

var (
	plotID = 0
)

func NewPlot(x, y int) *Plot {
	plotID++
	return &Plot{
		p:        graham.Point{X: x, Y: y},
		c:        color.White,
		isConvex: false,
		id:       plotID,
		r:        4,
	}
}

func (p *Plot) Update() error {
	return nil
}

func (p *Plot) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(p.p.X), float32(p.p.Y), float32(p.r), p.c, true)
}
