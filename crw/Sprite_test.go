package crw

import (
	"log"
	"math"
	"testing"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestLinearAnimation(t *testing.T) {
	s := BuildSprite().WithDestRect(99, 60, 10, 20).Build()
	s.SetLinear(time.Second*3, rl.Vector2{X: -99, Y: -60})

	s.update(0)

	if s.animationProgressRect.X != 0 {
		t.Fatalf("AnimationProgressRect should init at x = 0, actual x = %f", s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != 0 {
		t.Fatalf("AnimationProgressRect should init at y = 0, actual y = %f", s.animationProgressRect.Y)
	}

	s.update(time.Second)

	if s.animationProgressRect.X != 33 {
		t.Fatalf("AnimationProgressRect should update(1) to x = 33, actual x = %f", s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != 20 {
		t.Fatalf("AnimationProgressRect should update(1) to y = 20, actual y = %f", s.animationProgressRect.Y)
	}

	s.update(time.Second)

	if s.animationProgressRect.X != 66 {
		t.Fatalf("AnimationProgressRect should update(2) to x = 66, actual x = %f", s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != 40 {
		t.Fatalf("AnimationProgressRect should update(2) to y = 40, actual y = %f", s.animationProgressRect.Y)
	}

	s.update(time.Second)

	if s.animationProgressRect.X != s.GetWindowDestRect().X {
		t.Fatalf("AnimationProgressRect should update(3) to WindowDestRect.X, expected %f, actual %f", s.GetWindowDestRect().X, s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != s.GetWindowDestRect().Y {
		t.Fatalf("AnimationProgressRect should update(3) to WindowDestRect.Y, expected %f, actual %f", s.GetWindowDestRect().Y, s.animationProgressRect.Y)
	}

	s.update(time.Second)

	if s.animationProgressRect.X != s.GetWindowDestRect().X {
		t.Fatalf("AnimationProgressRect should update(4) to WindowDestRect.X, expected %f, actual %f", s.GetWindowDestRect().X, s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != s.GetWindowDestRect().Y {
		t.Fatalf("AnimationProgressRect should update(4) to WindowDestRect.Y, expected %f, actual %f", s.GetWindowDestRect().Y, s.animationProgressRect.Y)
	}

	log.Output(1, "[PASS]: TestLinearAnimation")
}

func TestEaseInAnimation(t *testing.T) {
	s := BuildSprite().WithDestRect(99, 60, 10, 20).Build()
	s.SetEaseIn(time.Second*3, rl.Vector2{X: -99, Y: -60})

	s.update(0)

	if s.animationProgressRect.X != 0 {
		t.Fatalf("AnimationProgressRect should init at x = 0, actual x = %f", s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != 0 {
		t.Fatalf("AnimationProgressRect should init at y = 0, actual y = %f", s.animationProgressRect.Y)
	}

	s.update(time.Second)

	expectedX := float32(math.Trunc(99 * timingQuad(.33)))
	actualX := float32(math.Trunc(float64(s.animationProgressRect.X)))
	if actualX != expectedX {
		t.Fatalf("AnimationProgressRect should update(1) to trunc'd x = %f, actual x = %f", expectedX, actualX)
	}

	expectedY := float32(math.Trunc(60 * timingQuad(.33)))
	actualY := float32(math.Trunc(float64(s.animationProgressRect.Y)))
	if actualY != expectedY {
		t.Fatalf("AnimationProgressRect should update(1) to trunc'd y = %f, actual y = %f", expectedY, actualY)
	}

	s.update(time.Second)

	expectedX = float32(math.Trunc(99 * timingQuad(.66)))
	actualX = float32(math.Trunc(float64(s.animationProgressRect.X)))
	if actualX != expectedX {
		t.Fatalf("AnimationProgressRect should update(2) to trunc'd x = %f, actual x = %f", expectedX, actualX)
	}

	expectedY = float32(math.Trunc(60 * timingQuad(.66)))
	actualY = float32(math.Trunc(float64(s.animationProgressRect.Y)))
	if actualY != expectedY {
		t.Fatalf("AnimationProgressRect should update(2) to trunc'd y = %f, actual y = %f", expectedY, actualY)
	}

	s.update(time.Second)

	if s.animationProgressRect.X != s.GetWindowDestRect().X {
		t.Fatalf("AnimationProgressRect should update(3) to WindowDestRect.X, expected %f, actual %f", s.GetWindowDestRect().X, s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != s.GetWindowDestRect().Y {
		t.Fatalf("AnimationProgressRect should update(3) to WindowDestRect.Y, expected %f, actual %f", s.GetWindowDestRect().Y, s.animationProgressRect.Y)
	}

	s.update(time.Second)

	if s.animationProgressRect.X != s.GetWindowDestRect().X {
		t.Fatalf("AnimationProgressRect should update(4) to WindowDestRect.X, expected %f, actual %f", s.GetWindowDestRect().X, s.animationProgressRect.X)
	}

	if s.animationProgressRect.Y != s.GetWindowDestRect().Y {
		t.Fatalf("AnimationProgressRect should update(4) to WindowDestRect.Y, expected %f, actual %f", s.GetWindowDestRect().Y, s.animationProgressRect.Y)
	}

	log.Output(1, "[PASS]: TestEaseInAnimation")
}
