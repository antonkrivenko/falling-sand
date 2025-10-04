package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const FontSize = 18
const Scale = 1
const ParticalRadius = 4
const ParticalSpeed = 4
const ParticalWidth = 10
const ParticalHeight = 10

type Point struct {
	X        int32
	Y        int32
	notEmpty bool
}

var score int

// func button(title string, posX, posY int32) {
// 	rl.DrawLine(posX-10, posY-10, posX+100, posY-10, rl.Blue)
// 	rl.DrawText(title, posX, posY, 16, rl.White)
// 	// rl.DrawLineEx(rl.Vector2{10, 10}, rl.Vector2{10, 20}, 5, rl.Blue)
// }

func initPoints() [][]Point {
	var points [][]Point
	for i := 0; i < 60; i++ {
		var row []Point
		for j := 0; j < 80; j++ {
			x := int32(i * ParticalWidth)
			y := int32(j * ParticalHeight)
			row = append(row, Point{X: x, Y: y})
		}
		points = append(points, row)
	}
	return points
}

// func checkNotEmpty(points [][]Point, i, j int, point Point) bool {
// 	if j > 50 {
// 		return point.notEmpty
// 	}
// 	if j < 10 {
// 		return point.notEmpty
// 	}

// 	if points[i][j-1].notEmpty {
// 		return true
// 	}
// 	return point.notEmpty
// 	// i := point.X
// 	// j := point.Y

// 	// if j > 1 && j < 50 && points[i][j-1].notEmpty {
// 	// 	notEmpty = true
// 	// }
// 	// if j > 1 && j < 50 && !points[i][j-1].notEmpty {
// 	// 	notEmpty = false
// 	// }
// 	// return notEmpty
// }

func nextPoints(points [][]Point) [][]Point {
	rows := len(points)
	cols := len(points[0])

	// Create a copy of the current grid
	next := make([][]Point, rows)
	for i := range next {
		next[i] = make([]Point, cols)
		for j := range next[i] {
			next[i][j] = points[i][j]
		}
	}

	// Iterate from bottom to top (so falling cells don't overwrite others)
	for i := rows - 2; i >= 0; i-- {
		for j := 0; j < cols; j++ {
			if points[i][j].notEmpty {
				// Try to fall straight down
				if !points[i+1][j].notEmpty {
					next[i][j].notEmpty = false
					next[i+1][j].notEmpty = true
					continue
				}

				// Try to fall down-left
				if j > 0 && !points[i+1][j-1].notEmpty {
					next[i][j].notEmpty = false
					next[i+1][j-1].notEmpty = true
					continue
				}

				// Try to fall down-right
				if j < cols-1 && !points[i+1][j+1].notEmpty {
					next[i][j].notEmpty = false
					next[i+1][j+1].notEmpty = true
					continue
				}
			}
		}
	}

	last := rows - 1
	isFull := true
	for j := 0; j < cols; j++ {
		if !points[last][j].notEmpty {
			isFull = false
		}
	}

	if isFull {
		for j := 0; j < cols; j++ {
			next[last][j].notEmpty = false
		}
		score += 10 //increase score
		// fmt.Println("Score", score)
	}

	return next
}

// func nextPoints(points [][]Point) [][]Point {
// 	var next [][]Point

// 	for i, row := range points {
// 		var nextRow []Point
// 		for j, point := range row {
// 			notEmpty := checkNotEmpty(points, i, j, point)
// 			nextRow = append(nextRow, Point{X: point.X, Y: point.Y, notEmpty: notEmpty})
// 		}
// 		next = append(next, nextRow)
// 	}
// 	return next
// }

func main() {
	// rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(800, 600, "Falling Sand")
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	var points [][]Point
	points = initPoints()

	for !rl.WindowShouldClose() {
		mouseX := int32(float32(rl.GetMouseX()) * Scale)
		mouseY := int32(float32(rl.GetMouseY()) * Scale)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			x := int32(mouseX / ParticalWidth)
			y := int32(mouseY / ParticalHeight)
			points[y][x].notEmpty = true
			points[y][x+1].notEmpty = true
			points[y][x+2].notEmpty = true
		}

		// rl.DrawCircle(mouseX, mouseY, ParticalRadius, rl.RayWhite)

		for _, row := range points {
			for _, point := range row {
				if point.notEmpty {
					rl.DrawRectangle(point.Y, point.X, ParticalWidth, ParticalHeight, rl.DarkPurple)
				} else {
					rl.DrawRectangle(point.Y, point.X, ParticalWidth, ParticalHeight, rl.Black)
				}
			}
		}
		// button("TESTING", 100, 120)
		scoreText := fmt.Sprintf("Score %d", score)
		rl.DrawText(scoreText, 10, 10, FontSize, rl.White) //create a text for score

		rl.EndDrawing()

		points = nextPoints(points)
	}
}
