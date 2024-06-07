package tests

import (
	"crowform/pkg"
	"log"
	"testing"
	"time"
)

func TestAnimationInitialFrameIdx(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
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
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
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
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
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
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
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
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()

	anim.Update(time.Second)

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 10 || actualRect.Y != 30 {
		t.Fatalf("Failed Update1 Src Rect Actual {X: %f, Y: %f}, expected {X: 10, Y: 30},", actualRect.X, actualRect.Y)
	}

	anim.Update(time.Second)

	actualRect = anim.GetCurrentSrcRect()

	if actualRect.X != 20 || actualRect.Y != 40 {
		t.Fatalf("Failed Update2 Src Rect Actual {X: %f, Y: %f}, expected {X: 10, Y: 30},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationUpdateSrcRect")
}

func TestAnimationUpdateSrcRectLoops(t *testing.T) {
	anim := pkg.BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
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
