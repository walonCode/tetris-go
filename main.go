package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/walonCode/tetris-go/internals/game"
)

const WIDTH int32 = 500
const HEIGHT int32 = 620
var LastUpdateTime float64 = 0

func main() {
	//raylib setup
	rl.SetConfigFlags(rl.FlagWindowResizable | rl.FlagWindowMaximized)
	rl.SetTargetFPS(60)
    rl.InitWindow(WIDTH, HEIGHT, "Tetris Go")
    defer rl.CloseWindow()
    font := rl.LoadFont("fonts/monogram.ttf")

    //game struct
    game := game.New()


    for !rl.WindowShouldClose() {
        rl.UpdateMusicStream(game.Music)
        game.HandleInput()
        if eventTriggered(0.2){
        	game.MoveBlockDown()
        }

        rl.BeginDrawing()
        rl.ClearBackground(rl.DarkBlue)
        rl.DrawTextEx(font, "Score", rl.Vector2{X:365, Y:15}, 38, 2, rl.White)
        rl.DrawTextEx(font, "Next", rl.Vector2{X:370, Y:175}, 38, 2, rl.White)

        if game.GameOver {
        	rl.DrawTextEx(font, "GAME OVER", rl.Vector2{X:320, Y:450}, 38, 2, rl.White)
        }
        rl.DrawRectangleRounded(rl.Rectangle{X:320, Y:55, Width: 170, Height: 60},0.3, 6, rl.Color{R: 44,G: 44,B: 127,A: 255})

        scoreText := fmt.Sprintf("%d", game.Score)
        textSize := rl.MeasureTextEx(font, scoreText, 38, 2)

        rl.DrawTextEx(font, scoreText, rl.Vector2{X:320+(170 - textSize.X)/2, Y:65}, 38, 2, rl.White)
        rl.DrawRectangleRounded(rl.Rectangle{X: 320, Y: 215, Width: 170, Height: 180}, 0.3, 6, rl.Color{R: 44,G: 44,B: 127,A: 255})

        game.Draw()
        rl.EndDrawing()
    }
}


func eventTriggered(interal float64)bool {
	currentTime := rl.GetTime()
	if currentTime - LastUpdateTime >= interal {
		LastUpdateTime = currentTime
		return true
	}
	return false
}
