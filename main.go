package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PalyerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Y += 2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	// ebitenutil.DebugPrint(screen, "Home")
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.X, g.Y)
	screen.DrawImage(g.PalyerImage.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	playerImage, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")

	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&Game{PalyerImage: playerImage, X: 100, Y: 200}); err != nil {
		log.Fatal(err)
	}
}
