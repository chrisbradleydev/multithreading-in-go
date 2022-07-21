package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount = 750
	viewRadius = 10
	adjRate = 0.025
)

var (
	pink   = color.RGBA{244, 114, 182, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth+1][screenHeight+1]int
	rWlock = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), pink)
		screen.Set(int(boid.position.x-1), int(boid.position.y), pink)
		screen.Set(int(boid.position.x), int(boid.position.y-1), pink)
		screen.Set(int(boid.position.x), int(boid.position.y+1), pink)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
