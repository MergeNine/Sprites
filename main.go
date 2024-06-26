package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"gaming in go ep. 2",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Init window", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Init renderer", err)
		return
	}
	defer renderer.Destroy()

	img, err := sdl.LoadBMP("assets/sprites/player.bmp")
	if err != nil {
		fmt.Println("loading player sprite:", err)
		return
	}
	playerTex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("Creating player Texture", err)
		return
	}
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		renderer.Copy(playerTex,
			&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
			&sdl.Rect{X: 0, Y: 0, W: 105, H: 105})

		renderer.Present()
	}
}
