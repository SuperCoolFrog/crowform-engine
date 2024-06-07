package pkg

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
