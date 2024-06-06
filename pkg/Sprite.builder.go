package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteBuilder struct {
	textureFileName string
	Origin          rl.Vector2
	srcRect         rl.Rectangle
	DestRect        rl.Rectangle
	rotation        float32
	colorTint       rl.Color
}

type Sprite struct {
	SpriteBuilder
	// texture rl.Texture2D
}

func BuildSprite() *SpriteBuilder {
	return &SpriteBuilder{
		textureFileName: "",
		Origin:          rl.Vector2{X: 0, Y: 0},
		srcRect:         rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		DestRect:        rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		rotation:        0,
		colorTint:       rl.White,
	}
}

func (builder *SpriteBuilder) WithTexture(filename string) *SpriteBuilder {
	builder.textureFileName = filename
	return builder
}

func (builder *SpriteBuilder) WithSourceRect(x float32, y float32, width float32, height float32) *SpriteBuilder {
	builder.srcRect = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	return builder
}
func (builder *SpriteBuilder) WithDestRect(x float32, y float32, width float32, height float32) *SpriteBuilder {
	builder.DestRect = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	return builder
}

func (builder *SpriteBuilder) WithRotation(rotation float32) *SpriteBuilder {
	builder.rotation = rotation
	return builder
}
func (builder *SpriteBuilder) WithColorTint(color rl.Color) *SpriteBuilder {
	builder.colorTint = color
	return builder
}

func (builder *SpriteBuilder) Build() *Sprite {
	return &Sprite{
		SpriteBuilder: *builder,
	}
}
