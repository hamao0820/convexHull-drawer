package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hamao0820/convexHull-drawer/graham"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	plots []Plot
}

func NewGame() *Game {
	plots := []Plot{}
	for i := 0; i < 10; i++ {
		plots = append(plots, *NewPlot(i*30, i*30))
	}
	return &Game{
		plots: plots,
	}
}

func (g *Game) Update() error {
	for i := range g.plots {
		if err := g.plots[i].Update(); err != nil {
			return err
		}
	}

	mouseX, mouseY := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.plots = append(g.plots, *NewPlot(mouseX, mouseY))
	}

	points := []graham.Point{}
	for i := range g.plots {
		points = append(points, g.plots[i].p)
	}

	// convexHull := graham.Scan(points)
	// for i := range g.plots {
	// 	g.plots[i].convexHull = convexHull
	// }

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.plots {
		g.plots[i].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func RunGame() {
	g := NewGame()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Convex Hull Drawer")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
