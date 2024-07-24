package crw

import (
	"crowform/internal/cache"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	AnimationBuilder
	loadedTexture   *rl.Texture2D
	currentFrameIdx int
	complete        bool
	waitTime        time.Duration
	waitCounter     time.Duration
	parent          *Actor
	flippedH        bool
	flippedV        bool
}

func (animation *Animation) Draw() {
	destRect := animation.GetWindowDestRect()
	srcRect := animation.srcRect

	if animation.flippedH {
		srcRect.Width *= -1
	}
	if animation.flippedV {
		srcRect.Height *= -1
	}

	rl.DrawTexturePro(
		*animation.getTexture(),
		srcRect,
		destRect,
		animation.Origin,
		animation.rotation,
		animation.colorTint,
	)
}
func (animation *Animation) getTexture() *rl.Texture2D {
	if animation.loadedTexture == nil {
		texture := cache.GetTexture2d(animation.textureFileName)
		animation.loadedTexture = &texture
		rl.SetTextureFilter(*animation.loadedTexture, rl.FilterBilinear)
	}

	return animation.loadedTexture
}

func (animation *Animation) update(deltaTime time.Duration) {
	if animation.complete {
		return
	}
	total := animation.waitCounter + deltaTime

	if total >= animation.waitTime {
		animation.updateFrame()
		animation.waitCounter = time.Duration(0)
	} else {
		animation.waitCounter = total
	}
}

func (animation *Animation) updateFrame() {
	if animation.currentFrameIdx == len(animation.frames)-1 {
		if !animation.playOnce {
			animation.currentFrameIdx = 0
		} else {
			animation.complete = true
			animation.onComplete()
		}
	} else {
		animation.currentFrameIdx += 1
	}

	frame := animation.frames[animation.currentFrameIdx]

	animation.srcRect.X = animation.srcRect.Width * frame.X
	animation.srcRect.Y = animation.srcRect.Height * frame.Y
}

func (animation *Animation) Reset() {
	animation.complete = false
	animation.currentFrameIdx = 0
}

func (animation *Animation) ReverseFrameOrder() {
	nu := make([]rl.Vector2, len(animation.frames))
	count := len(animation.frames) - 1

	for i := count; i > -1; i-- {
		nu[count-i] = animation.frames[i]
	}

	animation.frames = nu
}

func (animation *Animation) GetTotalAnimationTime() time.Duration {
	framesCount := float64(len(animation.frames))
	return time.Duration(float64(animation.waitTime) * framesCount)
}

func (animation *Animation) GetCurrentFrameIdx() int {
	return animation.currentFrameIdx
}

func (animation *Animation) GetCurrentSrcRect() rl.Rectangle {
	return animation.srcRect
}

func (animation *Animation) SetParent(parent *Actor) {
	animation.parent = parent
}

func (animation *Animation) GetWindowDestRect() rl.Rectangle {
	if animation.parent == nil {
		return animation.DestRect
	}

	parentPos := animation.parent.GetWindowPosition()
	windowDestRec := animation.DestRect
	windowDestRec.X += parentPos.X
	windowDestRec.Y += parentPos.Y

	return windowDestRec
}

func (me *Animation) SetFlipHorizontal(isFlipped bool) {
	me.flippedH = isFlipped
}
func (me *Animation) SetFlipVertically(isFlipped bool) {
	me.flippedV = isFlipped
}

func (me *Animation) SetPlayOnce(playOnce bool) {
	me.playOnce = playOnce
}

func (me *Animation) SetOnComplete(onComplete func()) {
	me.onComplete = onComplete
}
