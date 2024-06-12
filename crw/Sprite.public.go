package crw

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sprite *Sprite) GetWindowDestRect() rl.Rectangle {
	if sprite.parent == nil {
		return sprite.DestRect
	}

	parentPos := sprite.parent.GetWindowPosition()
	windowDestRec := sprite.DestRect
	windowDestRec.X += parentPos.X
	windowDestRec.Y += parentPos.Y

	return windowDestRec
}

func (sprite *Sprite) GetWindowPosition() rl.Vector2 {
	destRect := sprite.GetWindowDestRect()

	return rl.Vector2{
		X: destRect.X,
		Y: destRect.Y,
	}
}

func (sprite *Sprite) GetY() float32 {
	return sprite.DestRect.Y
}
func (sprite *Sprite) GetX() float32 {
	return sprite.DestRect.X
}

func (sprite *Sprite) SetY(y float32) {
	sprite.DestRect.Y = y
}

func (sprite *Sprite) SetX(x float32) {
	sprite.DestRect.X = x
}
