package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kodlan/gograph/cmd/gograph/samples"
	"github.com/kodlan/gograph/internal/turtle"
	"image/color"
	"log"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

type Game struct {
	windowWidth  int
	windowHeight int
	turtle       turtle.Mover
	step         turtle.Step
}

func (g *Game) Update() error {
	if g.turtle.Step() > g.turtle.MaxSteps() {
		return nil
	}

	g.step.Step(g.turtle)

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

func NewGame() *Game {
	return &Game{
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		turtle:       turtle.NewTurtle(windowWidth/2, windowHeight/2),
		step:         samples.NewSpiral(),
	}
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Go Graph demo")

	var game = NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
