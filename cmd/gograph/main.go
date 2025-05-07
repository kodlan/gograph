package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

type Game struct {
	windowWidth  int
	windowHeight int
}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
}

func (g Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return g.windowWidth, g.windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Go Graph demo")

	var game = &Game{windowHeight: windowHeight, windowWidth: windowWidth}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
