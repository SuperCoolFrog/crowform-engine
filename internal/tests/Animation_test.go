package tests

import (
	"crowform/pkg"
	"log"
	"testing"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestAnimationInitialFrameIdx(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(rl.Vector2{X: 0, Y: 0}).
		WithFrame(rl.Vector2{X: 10, Y: 10}).
		WithFrame(rl.Vector2{X: 20, Y: 20}).
		Build()

	actualIdx := anim.GetCurrentFrameIdx()

	if actualIdx != 0 {
		t.Fatalf("Failed initial index %d, expected %d", actualIdx, 0)
	}

	log.Output(1, "[PASS]: TestAnimationFrames")
}

func TestAnimationFrameIdxUpdate(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(rl.Vector2{X: 0, Y: 0}).
		WithFrame(rl.Vector2{X: 10, Y: 10}).
		WithFrame(rl.Vector2{X: 20, Y: 20}).
		Build()

	anim.Update(time.Second)

	actualIdx := anim.GetCurrentFrameIdx()

	if actualIdx != 1 {
		t.Fatalf("Failed updating index %d, expected %d", actualIdx, 1)
	}

	log.Output(1, "[PASS]: TestAnimationFrameIdxUpdate")
}

func TestAnimationFrameIdxLoops(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(rl.Vector2{X: 0, Y: 0}).
		WithFrame(rl.Vector2{X: 10, Y: 10}).
		WithFrame(rl.Vector2{X: 20, Y: 20}).
		Build()

	anim.Update(time.Second)
	anim.Update(time.Second)
	anim.Update(time.Second)

	actualIdx := anim.GetCurrentFrameIdx()

	if actualIdx != 0 {
		t.Fatalf("Failed looping index %d, expected %d", actualIdx, 0)
	}

	log.Output(1, "[PASS]: TestAnimationFrameIdxLoops")
}

func TestAnimationInitialSrcRect(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(rl.Vector2{X: 0, Y: 0}).
		WithFrame(rl.Vector2{X: 10, Y: 10}).
		WithFrame(rl.Vector2{X: 20, Y: 20}).
		Build()

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 0 || actualRect.Y != 0 {
		t.Fatalf("Failed Src Rect Actual {X: %f, Y: %f}, expected {X: 0, Y: 0},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationInitialSrcRect")
}

func TestAnimationUpdateSrcRect(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(rl.Vector2{X: 0, Y: 0}).
		WithFrame(rl.Vector2{X: 10, Y: 10}).
		WithFrame(rl.Vector2{X: 20, Y: 20}).
		Build()

	anim.Update(time.Second)

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 10 || actualRect.Y != 10 {
		t.Fatalf("Failed Src Rect Actual {X: %f, Y: %f}, expected {X: 10, Y: 10},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationUpdateSrcRect")
}

func TestAnimationUpdateSrcRectLoops(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(rl.Vector2{X: 0, Y: 0}).
		WithFrame(rl.Vector2{X: 10, Y: 10}).
		WithFrame(rl.Vector2{X: 20, Y: 20}).
		Build()

	anim.Update(time.Second)
	anim.Update(time.Second)
	anim.Update(time.Second)

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 0 || actualRect.Y != 0 {
		t.Fatalf("Failed Looping Src Rect Actual {X: %f, Y: %f}, expected {X: 0, Y: 0},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationUpdateSrcRectLoops")
}
