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
	player  Player = Player{rl.NewVector2(300, 400), rl.NewVector2(0, 0), 2}
	solids  []Solid
	bullets []Bullet
)

type Player struct {
	position rl.Vector2
	velocity rl.Vector2
	speed    float32
}

type Solid struct {
	collider rl.Rectangle
}

type Bullet struct {
	position rl.Vector2
	velocity rl.Vector2
	speed    float32
}

func colliding_player_solid() bool {
	for i := 0; i < len(solids); i++ {
		if rl.CheckCollisionRecs(rl.NewRectangle(player.position.X-18, player.position.Y-18, 18*2, 18*2), solids[i].collider) {
			return true
		}
	}

	return false
}

func input() {
	if rl.IsKeyDown(rl.KeyD) {
		player.velocity.X = +1
	}

	if rl.IsKeyDown(rl.KeyA) {
		player.velocity.X = -1
	}

	if rl.IsKeyDown(rl.KeyW) {
		player.velocity.Y = -1
	}

	if rl.IsKeyDown(rl.KeyS) {
		player.velocity.Y = +1

	}

	if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
		bullets = append(bullets, Bullet{player.position, rl.Vector2Subtract(rl.GetMousePosition(), player.position), 4})
	}
}

func update_player() {
	player.velocity = rl.Vector2Normalize(player.velocity)

	player.position = rl.Vector2Add(player.position, rl.Vector2Scale(player.velocity, player.speed))

	for colliding_player_solid() {
		player.position = rl.Vector2Subtract(player.position, player.velocity)
	}

	player.velocity = rl.NewVector2(0, 0)
}

func draw_obstacles() {
	for i := 0; i < len(solids); i++ {
		rl.DrawRectangleRec(solids[i].collider, rl.Blue)
	}
}

func draw_bullets() {
	for i := 0; i < len(bullets); i++ {
		rl.DrawCircleV(bullets[i].position, 16, rl.Green)
	}
}

func update() {
	running = !rl.WindowShouldClose()
	update_player()
	update_bullet()
}

func update_bullet() {

	for i := 0; i < len(bullets); i++ {
		bullets[i].velocity = rl.Vector2Normalize(bullets[i].velocity)
		bullets[i].position = rl.Vector2Add(bullets[i].position, rl.Vector2Scale(bullets[i].velocity, bullets[i].speed))
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangleRec(rl.NewRectangle(player.position.X-18, player.position.Y-18, 18*2, 18*2), rl.Red)
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
		input()
		update()
		draw()
	}

	quit()
}
