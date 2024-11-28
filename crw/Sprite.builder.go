package crw

import (
	"crowform/internal/cache"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpriteBuilder struct {
	textureFileName string
	Origin          rl.Vector2
	srcRect         rl.Rectangle
	DestRect        rl.Rectangle
	rotation        float32
	colorTint       rl.Color
	hasShader       bool
	shaderFileName  string
	initShader      func(shader *rl.Shader) rl.Shader
	updateShader    func(shader rl.Shader)
	useBlankTexture bool
}

type Sprite struct {
	SpriteBuilder
	spriteAnimation
	texture        *rl.Texture2D
	parent         *Actor
	shader         *rl.Shader
	queueForUpdate []func()
	flippedH       bool
	flippedV       bool
}

func BuildSprite() *SpriteBuilder {
	return &SpriteBuilder{
		textureFileName: "",
		Origin:          rl.Vector2{X: 0, Y: 0},
		srcRect:         rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		DestRect:        rl.Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		rotation:        0,
		colorTint:       rl.White,
		initShader:      func(shader *rl.Shader) rl.Shader { return *shader },
		updateShader:    func(shader rl.Shader) {},
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

func (builder *SpriteBuilder) WithShader(shaderFileName string) *SpriteBuilder {
	builder.shaderFileName = shaderFileName
	builder.hasShader = true
	return builder
}
func (builder *SpriteBuilder) WithInitShader(init func(*rl.Shader) rl.Shader) *SpriteBuilder {
	builder.initShader = init
	return builder
}
func (builder *SpriteBuilder) WithUpdateShader(update func(rl.Shader)) *SpriteBuilder {
	builder.updateShader = update
	return builder
}
func (builder *SpriteBuilder) WithBlankTexture() *SpriteBuilder {
	builder.useBlankTexture = true
	return builder
}

func (builder *SpriteBuilder) Build() *Sprite {
	sprite := &Sprite{
		SpriteBuilder: *builder,
	}
	sprite.animationType = spriteAnimType_NONE
	sprite.animationTime = time.Duration(0)
	sprite.animationDuration = time.Duration(0)
	sprite.animationState = spriteAnimState_NONE
	sprite.animationStartRect = sprite.DestRect
	sprite.animationProgressRect = sprite.DestRect
	sprite.onAnimationComplete = func() {}

	cache.QueueForPreload(func() {
		if builder.useBlankTexture {
			blankRenderTexture := rl.LoadRenderTexture(int32(sprite.DestRect.Width), int32(sprite.DestRect.Height))
			sprite.texture = &blankRenderTexture.Texture
		} else {
			sprite.getCachedTexture()
		}

		if sprite.hasShader {
			shader := sprite.getCachedShader()
			initedShader := sprite.initShader(shader)
			sprite.shader = &initedShader
		}
	})

	return sprite
}
