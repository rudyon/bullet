package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	window_width, window_height = 800, 800
	window_title                = "bullet"
)

var (
	running bool   = true
	player  Player = Player{300, 400, 2}
	solids  []Solid
	bullets []Bullet
)

type Player struct {
	x, y  float32
	speed float32
}

type Solid struct {
	collider rl.Rectangle
}

type Bullet struct {
	x, y     float32
	speed    float32
	velocity rl.Vector2
}

func colliding_player_solid() bool {
	for i := 0; i < len(solids); i++ {
		if rl.CheckCollisionRecs(rl.NewRectangle(player.x-18, player.y-18, 18*2, 18*2), solids[i].collider) {
			return true
		}
	}

	return false
}

func update_player() {
	// TODO: the code inside of here causes it so that moving diagonal is double the speed
	// i don't want this. i don't think. maybe i can leave it still.
	// for later consideration the collision took too long
	// i just wanna make it more workable later

	if rl.IsKeyDown(rl.KeyD) {
		player.x += player.speed

		for colliding_player_solid() {
			player.x--
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		player.x -= player.speed

		for colliding_player_solid() {
			player.x++
		}
	}

	if rl.IsKeyDown(rl.KeyW) {
		player.y -= player.speed

		for colliding_player_solid() {
			player.y++
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		player.y += player.speed

		for colliding_player_solid() {
			player.y--
		}
	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		bullets = append(bullets, Bullet{player.x, player.y, 4, rl.Vector2Subtract(rl.NewVector2(player.x, player.y), rl.GetMousePosition())})
	}
}

func draw_obstacles() {
	for i := 0; i < len(solids); i++ {
		rl.DrawRectangleRec(solids[i].collider, rl.Blue)
	}
}

func draw_bullets() {
	for i := 0; i < len(bullets); i++ {
		rl.DrawCircle(int32(bullets[i].x), int32(bullets[i].y), 16, rl.Green)
	}
}

func update() {
	running = !rl.WindowShouldClose()
	update_player()
	update_bullet()
}

func update_bullet() {
	for i := 0; i < len(bullets); i++ {
		bullets[i].x -= bullets[i].speed * rl.Vector2Normalize(bullets[i].velocity).X
		bullets[i].y -= bullets[i].speed * rl.Vector2Normalize(bullets[i].velocity).Y
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangleRec(rl.NewRectangle(player.x-18, player.y-18, 18*2, 18*2), rl.Red)
	draw_obstacles()
	draw_bullets()

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(window_width, window_height, window_title)
	rl.SetTargetFPS(60)

	solids = append(solids, Solid{rl.NewRectangle(400-32, 400-32, 40*2, 32*2)})
}

func quit() {
	rl.CloseWindow()
}
func main() {

	for running {
		update()
		draw()
	}

	quit()
}
