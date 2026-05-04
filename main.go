package main

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH int32 = 300
const HEIGHT int32 = 600

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.SetTargetFPS(60)
	
    rl.InitWindow(WIDTH, HEIGHT, "Tetris Go")
    defer rl.CloseWindow()


    for !rl.WindowShouldClose() {
        rl.BeginDrawing()

        
        rl.ClearBackground(rl.DarkBlue)
        
        rl.EndDrawing()
    }
}
