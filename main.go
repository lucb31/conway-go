package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lucb31/conway-go/logic"
)

const xSize, ySize = 320, 240

type Game struct {
	State  logic.State
	pixels []byte
}

func Init() Game {
	return Game{State: logic.InitRandomState(xSize, ySize)}
}

func (g *Game) Update() error {
	var err error
	err = g.State.Epoch()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, 4*screen.Bounds().Dy()*screen.Bounds().Dx())
	}
	for row := range screen.Bounds().Dy() {
		for col := range screen.Bounds().Dx() {
			if g.State.Vals[row][col] {
				g.pixels[4*row*col] = 0xFF
				g.pixels[4*row*col+1] = 0xFF
				g.pixels[4*row*col+2] = 0xFF
				g.pixels[4*row*col+3] = 0xFF
			} else {
				g.pixels[4*row*col] = 0
				g.pixels[4*row*col+1] = 0
				g.pixels[4*row*col+2] = 0
				g.pixels[4*row*col+3] = 0
			}
		}
	}
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return xSize, ySize
}

func main() {
	game := Init()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
