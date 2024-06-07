package pkg

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimationBuilder struct {
	textureFileName string
	Origin          rl.Vector2
	srcRect         rl.Rectangle
	DestRect        rl.Rectangle
	rotation        float32
	colorTint       rl.Color
	frames          []rl.Vector2
	playOnce        bool
	onComplete      func()
	framesPerSecond int32
}

func BuildAnimation() *AnimationBuilder {
	return &AnimationBuilder{
		textureFileName: "",
		Origin:          rl.Vector2{X: 0, Y: 0},
		srcRect:         rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		DestRect:        rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		rotation:        0,
		colorTint:       rl.White,
		frames:          make([]rl.Vector2, 0),
	}
}

func (builder *AnimationBuilder) WithTexture(filename string) *AnimationBuilder {
	builder.textureFileName = filename
	return builder
}
func (builder *AnimationBuilder) WithFramePerSec(fps int32) *AnimationBuilder {
	builder.framesPerSecond = fps
	return builder
}

func (builder *AnimationBuilder) WithSourceRect(x float32, y float32, width float32, height float32) *AnimationBuilder {
	builder.srcRect = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	return builder
}
func (builder *AnimationBuilder) WithDestRect(x float32, y float32, width float32, height float32) *AnimationBuilder {
	builder.DestRect = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	return builder
}

func (builder *AnimationBuilder) WithRotation(rotation float32) *AnimationBuilder {
	builder.rotation = rotation
	return builder
}
func (builder *AnimationBuilder) WithColorTint(color rl.Color) *AnimationBuilder {
	builder.colorTint = color
	return builder
}
func (builder *AnimationBuilder) WithFrame(frameXY rl.Vector2) *AnimationBuilder {
	builder.frames = append(builder.frames, frameXY)
	return builder
}
func (builder *AnimationBuilder) WithPlayOnce(then func()) *AnimationBuilder {
	builder.playOnce = true
	builder.onComplete = then
	return builder
}

func (builder *AnimationBuilder) Build() *Animation {
	return &Animation{
		AnimationBuilder: *builder,
		waitTime:         time.Duration(float64(time.Second) / float64(builder.framesPerSecond)),
		waitCounter:      time.Duration(0),
	}
}
