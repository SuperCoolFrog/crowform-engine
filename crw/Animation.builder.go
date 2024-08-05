package crw

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimationBuilder struct {
	Animation *Animation
}

func BuildAnimation() *AnimationBuilder {
	return &AnimationBuilder{
		Animation: &Animation{
			textureFileName: "",
			Origin:          rl.Vector2{X: 0, Y: 0},
			srcRect:         rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
			DestRect:        rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
			rotation:        0,
			colorTint:       rl.White,
			frames:          make([]rl.Vector2, 0),
		},
	}
}

func (builder *AnimationBuilder) WithTexture(filename string) *AnimationBuilder {
	builder.Animation.textureFileName = filename
	return builder
}
func (builder *AnimationBuilder) WithFramePerSec(fps int32) *AnimationBuilder {
	builder.Animation.framesPerSecond = fps
	return builder
}

func (builder *AnimationBuilder) WithSourceRect(x float32, y float32, width float32, height float32) *AnimationBuilder {
	builder.Animation.srcRect = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	return builder
}
func (builder *AnimationBuilder) WithDestRect(x float32, y float32, width float32, height float32) *AnimationBuilder {
	builder.Animation.DestRect = rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	return builder
}

func (builder *AnimationBuilder) WithRotation(rotation float32) *AnimationBuilder {
	builder.Animation.rotation = rotation
	return builder
}
func (builder *AnimationBuilder) WithColorTint(color rl.Color) *AnimationBuilder {
	builder.Animation.colorTint = color
	return builder
}
func (builder *AnimationBuilder) WithFrame(frameX float32, frameY float32) *AnimationBuilder {
	builder.Animation.frames = append(builder.Animation.frames, rl.Vector2{X: frameX, Y: frameY})
	return builder
}
func (builder *AnimationBuilder) WithPlayOnce(then func()) *AnimationBuilder {
	builder.Animation.playOnce = true
	builder.Animation.onComplete = then
	return builder
}

func (builder *AnimationBuilder) Build() *Animation {
	builder.Animation.waitTime = time.Duration(float64(time.Second) / float64(builder.Animation.framesPerSecond))
	builder.Animation.waitCounter = time.Duration(0)

	if len(builder.Animation.frames) > 0 {
		firstFrame := builder.Animation.frames[0]

		builder.Animation.srcRect.X = firstFrame.X * builder.Animation.srcRect.Width
		builder.Animation.srcRect.Y = firstFrame.Y * builder.Animation.srcRect.Height
	}

	return builder.Animation
}
