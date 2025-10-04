package main

import rl "github.com/gen2brain/raylib-go/raylib"

const FontSize = 18
const Scale = 1
const ParticalRadius = 2
const ParticalSpeed = 4

type Point struct {
	X int32
	Y int32
}

func isOccupied(points []Point, x, y int32, selfIndex int) bool {
	for i, p := range points {
		if i == selfIndex {
			continue
		}
		dx := p.X - x
		dy := p.Y - y
		if dx*dx+dy*dy < (ParticalRadius*2)*(ParticalRadius*2) {
			return true
		}
	}
	return false
}

func main() {
	// rl.SetConfigFlags(rl.FlagWindowHighdpi)

	rl.InitWindow(800, 600, "Falling Sand")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var points []Point

	for !rl.WindowShouldClose() {
		mouseX := int32(float32(rl.GetMouseX()) * Scale)
		mouseY := int32(float32(rl.GetMouseY()) * Scale)

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			points = append(points, Point{mouseX, mouseY})
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawCircle(mouseX, mouseY, ParticalRadius, rl.RayWhite)

		for i, point := range points {
			rl.DrawCircle(point.X, point.Y, ParticalRadius, rl.Red)

			// Платформа на Y=400, шарики не падають нижче
			if point.Y < 400-ParticalRadius {
				nextY := point.Y + ParticalSpeed

				// Якщо під шариком вільно — падаємо вниз
				if !isOccupied(points, point.X, nextY, i) {
					points[i].Y = nextY
					continue
				}

				// Якщо вниз зайнято — пробуємо вниз-вліво
				if !isOccupied(points, point.X-ParticalRadius*2, nextY, i) {
					points[i].X -= ParticalRadius * 2
					points[i].Y = nextY
					continue
				}

				// Якщо вниз-вліво зайнято — пробуємо вниз-вправо
				if !isOccupied(points, point.X+ParticalRadius*2, nextY, i) {
					points[i].X += ParticalRadius * 2
					points[i].Y = nextY
					continue
				}
			}
		}

		rl.EndDrawing()
	}
}
