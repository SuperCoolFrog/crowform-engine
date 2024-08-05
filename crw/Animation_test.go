package crw

import (
	"log"
	"testing"
	"time"
)

func TestAnimationInitialFrameIdx(t *testing.T) {
	anim := BuildAnimation().
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
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()

	anim.update(time.Second)

	actualIdx := anim.GetCurrentFrameIdx()

	if actualIdx != 1 {
		t.Fatalf("Failed updating index %d, expected %d", actualIdx, 1)
	}

	log.Output(1, "[PASS]: TestAnimationFrameIdxUpdate")
}

func TestAnimationFrameIdxLoops(t *testing.T) {
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()

	anim.update(time.Second)
	anim.update(time.Second)
	anim.update(time.Second)

	actualIdx := anim.GetCurrentFrameIdx()

	if actualIdx != 0 {
		t.Fatalf("Failed looping index %d, expected %d", actualIdx, 0)
	}

	log.Output(1, "[PASS]: TestAnimationFrameIdxLoops")
}

func TestAnimationInitialSrcRect(t *testing.T) {
	anim := BuildAnimation().
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
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()

	anim.update(time.Second)

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 10 || actualRect.Y != 30 {
		t.Fatalf("Failed Update1 Src Rect Actual {X: %f, Y: %f}, expected {X: 10, Y: 30},", actualRect.X, actualRect.Y)
	}

	anim.update(time.Second)

	actualRect = anim.GetCurrentSrcRect()

	if actualRect.X != 20 || actualRect.Y != 40 {
		t.Fatalf("Failed Update2 Src Rect Actual {X: %f, Y: %f}, expected {X: 10, Y: 30},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationUpdateSrcRect")
}

func TestAnimationUpdateSrcRectLoops(t *testing.T) {
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()

	anim.update(time.Second)
	anim.update(time.Second)
	anim.update(time.Second)

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 0 || actualRect.Y != 0 {
		t.Fatalf("Failed Looping Src Rect Actual {X: %f, Y: %f}, expected {X: 0, Y: 0},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationUpdateSrcRectLoops")
}

func TestAnimationCallOnce(t *testing.T) {
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()
	anim.SetPlayOnce(true)
	anim.SetOnComplete(func() {})

	anim.update(time.Second)
	anim.update(time.Second)
	anim.update(time.Second)
	anim.update(time.Second)

	actualRect := anim.GetCurrentSrcRect()

	if actualRect.X != 20 || actualRect.Y != 40 {
		t.Fatalf("Failed Calling Once Src Rect Actual {X: %f, Y: %f}, expected {X: 2, Y: 4},", actualRect.X, actualRect.Y)
	}

	log.Output(1, "[PASS]: TestAnimationCallOnce")
}

func TestAnimationOnCompleteCalled(t *testing.T) {
	expected := true
	actual := false
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 0).
		WithFrame(1, 3).
		WithFrame(2, 4).
		Build()
	anim.SetPlayOnce(true)
	anim.SetOnComplete(func() {
		actual = expected
	})

	anim.update(time.Second)
	anim.update(time.Second)
	anim.update(time.Second)
	anim.update(time.Second)

	if actual != expected {
		t.Fatalf("Failed Calling OnComplete")
	}

	log.Output(1, "[PASS]: TestAnimationOnCompleteCalled")
}

func TestAnimationFrameNonZeroXStartNoUpdate(t *testing.T) {
	expected := float32(20)
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(2, 0).
		Build()
	actual := anim.srcRect.X

	if actual != expected {
		t.Fatalf("Animation Start with Non Zero X fail -- expected %f, actual %f", expected, actual)
	}

	log.Output(1, "[PASS]: TestAnimationFrameNonZeroXStartNoUpdate")
}

func TestAnimationFrameNonZeroYStartNoUpdate(t *testing.T) {
	expected := float32(30)
	anim := BuildAnimation().
		WithSourceRect(0, 0, 10, 10).
		WithDestRect(0, 0, 10, 10).
		WithFramePerSec(1).
		WithFrame(0, 3).
		Build()
	actual := anim.srcRect.Y

	if actual != expected {
		t.Fatalf("Animation Start with Non Zero Y fail -- expected %f, actual %f", expected, actual)
	}

	log.Output(1, "[PASS]: TestAnimationFrameNonZeroYStartNoUpdate")
}
