package crw

import (
	"crowform/internal/tools"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type spriteAnimType int
type spriteAnimState int

type spriteAnimation struct {
	animationType         spriteAnimType
	animationTime         time.Duration
	animationDuration     time.Duration
	animationState        spriteAnimState
	animationStartRect    rl.Rectangle
	animationProgressRect rl.Rectangle
	animationOpacity      float32
	onAnimationComplete   func()
}

const (
	spriteAnimType_NONE     spriteAnimType = 0
	spriteAnimType_LINEAR   spriteAnimType = 1
	spriteAnimType_EASE_IN  spriteAnimType = 2
	spriteAnimType_EASE_OUT spriteAnimType = 3
	spriteAnimType_FADE_IN  spriteAnimType = 4
	spriteAnimType_FADE_OUT spriteAnimType = 5

	spriteAnimState_NONE        spriteAnimState = -1
	spriteAnimState_INIT        spriteAnimState = 0
	spriteAnimState_PROGRESSING spriteAnimState = 1
	spriteAnimState_COMPLETE    spriteAnimState = 2
)

func (sprite *Sprite) setupMovementAnimation(sAnimType spriteAnimType, totalTime time.Duration, startDelta rl.Vector2) {
	sprite.animationTime = time.Duration(0)
	sprite.animationDuration = totalTime

	sprite.animationType = sAnimType
	sprite.animationState = spriteAnimState_INIT

	startVector := rl.Vector2Add(sprite.GetWindowPosition(), startDelta)
	dest := sprite.GetWindowDestRect()

	sprite.animationStartRect = rl.NewRectangle(startVector.X, startVector.Y, dest.Width, dest.Height)
	sprite.animationProgressRect = sprite.animationStartRect
}

func (sprite *Sprite) SetLinear(totalTime time.Duration, startDelta rl.Vector2) *spriteAnimation {
	sprite.addToUpdateQueue(func() {
		sprite.setupMovementAnimation(spriteAnimType_LINEAR, totalTime, startDelta)
	})
	return &sprite.spriteAnimation
}

func (sprite *Sprite) SetEaseIn(totalTime time.Duration, startDelta rl.Vector2) *spriteAnimation {
	sprite.addToUpdateQueue(func() {
		sprite.setupMovementAnimation(spriteAnimType_EASE_IN, totalTime, startDelta)
	})
	return &sprite.spriteAnimation
}
func (sprite *Sprite) SetEaseOut(totalTime time.Duration, startDelta rl.Vector2) *spriteAnimation {
	sprite.addToUpdateQueue(func() {
		sprite.setupMovementAnimation(spriteAnimType_EASE_OUT, totalTime, startDelta)
	})
	return &sprite.spriteAnimation
}

func (sprite *Sprite) SetFadeIn(totalTime time.Duration) *spriteAnimation {
	sprite.addToUpdateQueue(func() {
		sprite.animationTime = time.Duration(0)
		sprite.animationDuration = totalTime
		sprite.animationType = spriteAnimType_FADE_IN
		sprite.animationState = spriteAnimState_INIT
	})
	return &sprite.spriteAnimation
}
func (sprite *Sprite) SetFadeOut(totalTime time.Duration) *spriteAnimation {
	sprite.addToUpdateQueue(func() {
		sprite.animationTime = time.Duration(0)
		sprite.animationDuration = totalTime
		sprite.animationType = spriteAnimType_FADE_OUT
		sprite.animationState = spriteAnimState_INIT
	})
	return &sprite.spriteAnimation
}

func (spriteAnim *spriteAnimation) WithOnAnimationComplete(onComplete func()) {
	spriteAnim.onAnimationComplete = onComplete
}

func (sprite *Sprite) getAnimationDestRect() rl.Rectangle {
	if sprite.animationType == spriteAnimType_NONE || sprite.animationState == spriteAnimState_COMPLETE {
		return sprite.GetWindowDestRect()
	} else {
		return sprite.animationProgressRect
	}
}

func (sprite *Sprite) updateAnimations(deltaTime time.Duration) {
	if sprite.animationState == spriteAnimState_COMPLETE {
		return
	}

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
		sprite.updateAnimationPosition(progress)
	case spriteAnimType_EASE_IN:
		progress = sprite.getEaseInProgress(timeFraction)
		sprite.updateAnimationPosition(progress)
	case spriteAnimType_EASE_OUT:
		progress = sprite.getEaseOutProgress(timeFraction)
		sprite.updateAnimationPosition(progress)
	case spriteAnimType_FADE_IN:
		progress = timingLinear(timeFraction)
		sprite.setTextureOpacity(progress)
	case spriteAnimType_FADE_OUT:
		progress = asOutFunction(timingLinear)(timeFraction)
		sprite.setTextureOpacity(progress)
	}

	if timeFraction == 1 {
		sprite.animationState = spriteAnimState_COMPLETE
		sprite.spriteAnimation.onAnimationComplete()
		sprite.spriteAnimation.onAnimationComplete = func() {}

		switch sprite.animationType {
		case spriteAnimType_FADE_IN:
			sprite.reloadTexture()
		case spriteAnimType_FADE_OUT:
			sprite.reloadTexture()
		}
	}
}

func (sprite *Sprite) updateAnimationPosition(progress float64) {
	destRect := sprite.GetWindowDestRect()
	startRect := sprite.animationStartRect
	diffRect := tools.RectangleSubXY(destRect, startRect)
	sprite.animationProgressRect.X = startRect.X + float32(progress*float64(diffRect.X))
	sprite.animationProgressRect.Y = startRect.Y + float32(progress*float64(diffRect.Y))
}

func (sprite *Sprite) getLinearProgress(timeFraction float64) float64 {
	return timingLinear(timeFraction)
}

func (sprite *Sprite) getEaseInProgress(timeFraction float64) float64 {
	return timingQuad(timeFraction)
}

func (sprite *Sprite) getEaseOutProgress(timeFraction float64) float64 {
	return asOutFunction(timingQuad)(timeFraction)
}

// https://javascript.info/js-animation

func timingLinear(in float64) float64 {
	return in
}
func timingQuad(in float64) float64 {
	return tools.RoundFloat(math.Pow(in, 2), 4)
}

// func timingBounce(in float64) float64 {
// 	var a float64 = 0
// 	var b float64 = 1

// 	for {
// 		if in >= (7-4*a)/11 {
// 			return -math.Pow((11-6*a-11*in)/4, 2) + math.Pow(b, 2)
// 		}
// 	}
// }

// accepts a timing function, returns the transformed variant
func asOutFunction(timing func(float64) float64) func(float64) float64 {
	return func(timeFraction float64) float64 {
		return 1 - timing(1-timeFraction)
	}
}
