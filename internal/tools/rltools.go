package tools

import rl "github.com/gen2brain/raylib-go/raylib"

/*
Rectangle {r1.x - r2.x, r1.y - r2.y, r1.w, r1.h}

maintains w,h from r1
*/
func RectangleSubXY(r1 rl.Rectangle, r2 rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(r1.X-r2.X, r1.Y-r2.Y, r1.Width, r1.Height)
}

/*
Rectangle {r1.x + r2.x, r1.y + r2.y, r1.w, r1.h}

maintains w,h from r1
*/
func RectangleAddXY(r1 rl.Rectangle, r2 rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(r1.X-r2.X, r1.Y-r2.Y, r1.Width, r1.Height)
}
