package pkg

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
}

func (animation *Animation) Draw() {
	if animation.loadedTexture == nil {
		texture := cache.GetTexture2d(animation.textureFileName)
		animation.loadedTexture = &texture
	}

	rl.DrawTexturePro(
		*animation.loadedTexture,
		animation.srcRect,
		animation.DestRect,
		animation.Origin,
		animation.rotation,
		animation.colorTint,
	)
}

func (animation *Animation) Update(deltaTime time.Duration) {
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

	animation.srcRect.X = frame.X
	animation.srcRect.Y = frame.Y
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
