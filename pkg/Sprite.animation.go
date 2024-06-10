package pkg

import (
	"crowform/internal/tools"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type spriteAnimType int
type spriteAnimState int

const (
	spriteAnimType_NONE    spriteAnimType = 0
	spriteAnimType_LINEAR  spriteAnimType = 1
	spriteAnimType_EASE_IN spriteAnimType = 2

	spriteAnimState_NONE        spriteAnimState = -1
	spriteAnimState_INIT        spriteAnimState = 0
	spriteAnimState_PROGRESSING spriteAnimState = 1
	spriteAnimState_COMPLETE    spriteAnimState = 2
)

func (sprite *Sprite) SetEaseIn(totalTime time.Duration, startDelta rl.Vector2) {
	sprite.animationTime = time.Duration(0)
	sprite.animationDuration = totalTime

	sprite.animationType = spriteAnimType_EASE_IN
	sprite.animationState = spriteAnimState_INIT

	startVector := rl.Vector2Add(sprite.GetWindowPosition(), startDelta)
	dest := sprite.GetWindowDestRect()

	sprite.animationStartRect = rl.NewRectangle(startVector.X, startVector.Y, dest.Width, dest.Height)
	sprite.animationProgressRect = sprite.animationStartRect
}

func (sprite *Sprite) SetLinear(totalTime time.Duration, startDelta rl.Vector2) {
	sprite.animationTime = time.Duration(0)
	sprite.animationDuration = totalTime

	sprite.animationType = spriteAnimType_LINEAR
	sprite.animationState = spriteAnimState_INIT

	startVector := rl.Vector2Add(sprite.GetWindowPosition(), startDelta)
	dest := sprite.GetWindowDestRect()

	sprite.animationStartRect = rl.NewRectangle(startVector.X, startVector.Y, dest.Width, dest.Height)
	sprite.animationProgressRect = sprite.animationStartRect
}

func (sprite *Sprite) getAnimationDestRect() rl.Rectangle {
	if sprite.animationType == spriteAnimType_NONE || sprite.animationState == spriteAnimState_COMPLETE {
		return sprite.GetWindowDestRect()
	} else {
		return sprite.animationProgressRect
	}
}

func (sprite *Sprite) updateAnimations(deltaTime time.Duration) {
	sprite.animationState = spriteAnimState_PROGRESSING
	sprite.animationTime += deltaTime

	// timeFraction goes from 0 to 1
	timeFraction := float64(sprite.animationTime) / float64(sprite.animationDuration)
	if timeFraction > 1 {
		timeFraction = 1
	}

	// calculate the current animation progress
	var progress float64 = 0
	switch sprite.animationType {
	case spriteAnimType_LINEAR:
		progress = sprite.getLinearProgress(timeFraction)
	case spriteAnimType_EASE_IN:
		progress = sprite.getEaseInProgress(timeFraction)
	}

	destRect := sprite.GetWindowDestRect()
	startRect := sprite.animationStartRect
	diffRect := tools.RectangleSubXY(destRect, startRect)
	sprite.animationProgressRect.X = startRect.X + float32(progress*float64(diffRect.X))
	sprite.animationProgressRect.Y = startRect.Y + float32(progress*float64(diffRect.Y))

	if timeFraction == 1 {
		sprite.animationState = spriteAnimState_COMPLETE
	}
}

func (sprite *Sprite) getLinearProgress(timeFraction float64) float64 {
	return timingLinear(timeFraction)
}

func (sprite *Sprite) getEaseInProgress(timeFraction float64) float64 {
	return timingQuad(timeFraction)
}

func timingLinear(in float64) float64 {
	return in
}
func timingQuad(in float64) float64 {
	return tools.RoundFloat(math.Pow(in, 2), 4)
}
