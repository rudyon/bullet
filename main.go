package main

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	player Player = Player{300, 400, 2}
	solids []Solid
)

type Player struct {
	x, y  float32
	speed float32
}

type Solid struct {
	collider rl.Rectangle
}

func player_colliding() bool {
	for i := 0; i < len(solids); i++ {
		if rl.CheckCollisionRecs(rl.NewRectangle(player.x-18, player.y-18, 18*2, 18*2), solids[i].collider) {
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
		player.x += player.speed

		for player_colliding() {
			player.x--
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		player.x -= player.speed

		for player_colliding() {
			player.x++
		}
	}

	if rl.IsKeyDown(rl.KeyW) {
		player.y -= player.speed

		for player_colliding() {
			player.y++
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		player.y += player.speed

		for player_colliding() {
			player.y--
		}
	}
}

func draw_obstacles() {
	for i := 0; i < len(solids); i++ {
		rl.DrawRectangleRec(solids[i].collider, rl.Blue)
	}
}

func update() {
	player_update()
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangleRec(rl.NewRectangle(player.x-18, player.y-18, 18*2, 18*2), rl.Red)
	draw_obstacles()

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(800, 800, "bullet")
	rl.SetTargetFPS(60)

	solids = append(solids, Solid{rl.NewRectangle(400-32, 400-32, 40*2, 32*2)})
}

func main() {

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.CloseWindow()
}
