package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	x     float32 = 300
	y     float32 = 400
	speed float32 = 2

	obstacles []rl.Rectangle
)

func player_colliding() bool {
	for i := 0; i < len(obstacles); i++ {
		if rl.CheckCollisionRecs(rl.NewRectangle(x-18, y-18, 18*2, 18*2), obstacles[i]) {
			return true
		}
	}

	return false
}

func player_update() {
	// TODO: the code inside of here causes it so that moving diagonal is double the speed
	// i don't want this. i don't think. maybe i can leave it still.
	// for later consideration the collision took too long
	// i just wanna make it more workable later

	if rl.IsKeyDown(rl.KeyD) {
		x += speed

		for player_colliding() {
			x--
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		x -= speed

		for player_colliding() {
			x++
		}
	}

	if rl.IsKeyDown(rl.KeyW) {
		y -= speed

		for player_colliding() {
			y++
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		y += speed

		for player_colliding() {
			y--
		}
	}
}

func draw_obstacles() {
	for i := 0; i < len(obstacles); i++ {
		rl.DrawRectangleRec(obstacles[i], rl.Blue)
	}
}

func main() {
	rl.InitWindow(800, 800, "bullet")
	rl.SetTargetFPS(60)

	obstacles = append(obstacles, rl.NewRectangle(400-32, 400-32, 40*2, 32*2))

	for !rl.WindowShouldClose() {
		player_update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangleRec(rl.NewRectangle(x-18, y-18, 18*2, 18*2), rl.Red)
		draw_obstacles()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
