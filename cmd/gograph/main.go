package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kodlan/gograph/cmd/gograph/game"
	"github.com/kodlan/gograph/cmd/gograph/samples"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Go Graph demo")

	var stepDrawer = samples.NewSquareSpiral()

	var g = game.NewGame(windowWidth, windowHeight, stepDrawer)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
