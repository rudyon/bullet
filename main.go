package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	x     float32 = 300
	y     float32 = 400
	speed float32 = 2

	obstacle_rec rl.Rectangle = rl.NewRectangle(400-32, 400-32, 32*2, 32*2)
)

func player_update() {
	// TODO: the code inside of here causes it so that moving diagonal is double the speed
	// i don't want this. i don't think. maybe i can leave it still.
	// for later consideration the collision took too long
	// i just wanna make it more workable later

	if rl.IsKeyDown(rl.KeyD) {
		x += speed

		for rl.CheckCollisionRecs(rl.NewRectangle(x-18, y-18, 18*2, 18*2), obstacle_rec) {
			x--
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		x -= speed

		for rl.CheckCollisionRecs(rl.NewRectangle(x-18, y-18, 18*2, 18*2), obstacle_rec) {
			x++
		}
	}

	if rl.IsKeyDown(rl.KeyW) {
		y -= speed

		for rl.CheckCollisionRecs(rl.NewRectangle(x-18, y-18, 18*2, 18*2), obstacle_rec) {
			y++
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		y += speed

		for rl.CheckCollisionRecs(rl.NewRectangle(x-18, y-18, 18*2, 18*2), obstacle_rec) {
			y--
		}
	}
}

func main() {
	rl.InitWindow(800, 800, "bullet")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		player_update()

		rl.DrawRectangleRec(rl.NewRectangle(x-18, y-18, 18*2, 18*2), rl.Red)
		rl.DrawRectangleRec(obstacle_rec, rl.Blue)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
