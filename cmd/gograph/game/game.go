package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kodlan/gograph/internal/turtle"
	"image/color"
)

type Game struct {
	windowWidth  int
	windowHeight int
	turtle       turtle.Drawer
}

func (g *Game) Update() error {
	g.turtle.NextStep()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	for _, v := range g.turtle.Lines() {
		vector.StrokeLine(screen, v[0], v[1], v[2], v[3], 1, g.turtle.Color(), false)
	}
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return g.windowWidth, g.windowHeight
}

func NewGame(windowWidth, windowHeight int, stepDrawer turtle.StepDrawer) *Game {
	var centerX, centerY = float32(windowWidth / 2), float32(windowHeight / 2)
	return &Game{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		turtle:       turtle.NewTurtle(centerX, centerY, stepDrawer),
	}
}
