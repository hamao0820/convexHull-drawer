package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hamao0820/convexHull-drawer/graham"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	plots      []*Plot
	convexHull []*Plot
}

func NewGame() *Game {
	plots := []*Plot{}
	return &Game{
		plots: plots,
	}
}

func (g *Game) Update() error {

	mouseX, mouseY := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.plots = append(g.plots, NewPlot(mouseX, mouseY))
	}

	convexHull := graham.Scan(g.plots)
	for i := range g.plots {
		g.plots[i].isConvex = false
	}

	for i := range convexHull {
		convexHull[i].isConvex = true
	}

	g.convexHull = convexHull

	for i := range g.plots {
		if err := g.plots[i].Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
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
