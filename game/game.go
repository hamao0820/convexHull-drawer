package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hamao0820/convexHull-drawer/graham"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var textArea = ebiten.NewImage(ScreenWidth, 15)

func init() {
	textArea.Fill(color.Black)
}

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

func (g *Game) DrawConvexHull(screen *ebiten.Image) {
	if len(g.convexHull) < 2 {
		return
	}
	for i := 0; i < len(g.convexHull)-1; i++ {
		vector.StrokeLine(screen, float32(g.convexHull[i].X()), float32(g.convexHull[i].Y()), float32(g.convexHull[i+1].X()), float32(g.convexHull[i+1].Y()), 1, color.RGBA{0, 0, 0, 255}, true)
	}

	vector.StrokeLine(screen, float32(g.convexHull[len(g.convexHull)-1].X()), float32(g.convexHull[len(g.convexHull)-1].Y()), float32(g.convexHull[0].X()), float32(g.convexHull[0].Y()), 1, color.RGBA{0, 0, 0, 255}, true)
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

	g.DrawConvexHull(screen)
	for i := range g.plots {
		g.plots[i].Draw(screen)
	}

	screen.DrawImage(textArea, nil)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Num of vertices: %d/%d", len(g.convexHull), len(g.plots)))
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
