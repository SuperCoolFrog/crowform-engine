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

func (me *Sprite) SetFlipHorizontal(isFlipped bool) {
	me.flippedH = isFlipped
}
func (me *Sprite) SetFlipVertically(isFlipped bool) {
	me.flippedV = isFlipped
}
func (me *Sprite) SetTint(color rl.Color) {
	me.colorTint = color
}

func (me *Sprite) SetOpacity(opacity float64) {
	me.setTextureOpacity(opacity)
}

func (me *Sprite) GetTexture2d() rl.Texture2D {
	return *me.getCachedTexture()
}

func (me *Sprite) SetTexture2d(texture *rl.Texture2D) {
	me.texture = texture
}
