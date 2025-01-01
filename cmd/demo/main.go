package main

import (
	"renderIdk/internal/camera"
	"renderIdk/internal/config"
	"renderIdk/internal/render"
	"renderIdk/internal/vector"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	config.LoadConfig()

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	defer sdl.Quit()

	width := int32(config.WinWidth)
	height := int32(config.WinHeight)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	cam := camera.Camera{
		NearClip: 1,
		FarClip:  1000,
		FovX:     60.0,
		Position: vector.Vector3{
			X: 0.0,
			Y: 0.0,
			Z: 50.0,
		},
	}

	scene := render.Scene{Camera: cam}
	scene = scene.AddCuboid(vector.Vector3{X: 10, Y: 10, Z: 5.5}, vector.Vector3{X: 0, Y: 0, Z: 20}, 255, 0, 255, 255)
	scene = scene.AddCuboid(vector.Vector3{X: 8, Y: 8, Z: 4}, vector.Vector3{X: 20, Y: 10, Z: 20}, 255, 255, 0, 255)
	scene = scene.AddCuboid(vector.Vector3{X: 12, Y: 12, Z: 6}, vector.Vector3{X: -20, Y: 0, Z: 20}, 0, 255, 0, 255)
	scene = scene.AddCuboid(vector.Vector3{X: 15, Y: 15, Z: 7}, vector.Vector3{X: -30, Y: -10, Z: 10}, 0, 0, 255, 255)
	scene = scene.AddCuboid(vector.Vector3{X: 5, Y: 5, Z: 3}, vector.Vector3{X: 40, Y: -5, Z: 15}, 255, 0, 0, 255)
	scene = scene.AddCuboid(vector.Vector3{X: 7, Y: 7, Z: 5}, vector.Vector3{X: 10, Y: 20, Z: 25}, 255, 165, 0, 255)

	var lastTime uint64
	running := true
	for running {
		currentTime := sdl.GetTicks64()
		deltaTime := float64(currentTime-lastTime) / 1000
		lastTime = currentTime

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		scene.Camera = scene.Camera.HandleInput(sdl.GetKeyboardState(), deltaTime)

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		scene.Render(renderer)

		renderer.Present()

		frameTime := float64(sdl.GetTicks64() - currentTime)
		remainingTime := config.TargetFrameTime - frameTime
		if remainingTime > 0 {
			sdl.Delay(uint32(remainingTime))
		}
	}
}
