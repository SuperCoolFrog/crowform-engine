package pkg

import (
	"crowform/internal/tools"
	"log"
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestClickEventNoBubble(t *testing.T) {
	expected := "Top Clicked"
	res := ""

	a1 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 1).Build()
	a1.SetMouseDownHandler(func(mousePos rl.Vector2) bool {
		res = "Bottom Clicked"
		return false
	})

	a2 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 2).Build()
	a2.SetMouseDownHandler(func(mousePos rl.Vector2) bool {
		res = expected
		return false
	})

	game := BuildGame().Build()
	scene1 := BuildScene("s1", game).Build()

	scene1.AddChild(a1)
	scene1.AddChild(a2)
	game.GoToScene("s1")
	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})

	if res != expected {
		t.Fatalf("No Bubble Failed actual %s, expected %s", res, expected)
	}

	log.Output(1, "[PASS]: TestClickEventNoBubble")
}

func TestClickCallsHandlerOnceWhileDown(t *testing.T) {
	expected := 1
	res := 0

	a2 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 2).Build()
	a2.SetMouseDownHandler(func(mousePos rl.Vector2) bool {
		res++
		return false
	})

	game := BuildGame().Build()
	scene1 := BuildScene("s1", game).Build()

	scene1.AddChild(a2)
	game.GoToScene("s1")

	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})
	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})
	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})
	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})

	if res != expected {
		t.Fatalf("Call once failed actual %d, expected %d", res, expected)
	}

	log.Output(1, "[PASS]: TestClickCallsHandlerOnceWhileDown")
}

func TestClickCallsHandlerAfterReleased(t *testing.T) {
	expected := 2
	res := 0

	a2 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 2).Build()
	a2.SetMouseDownHandler(func(mousePos rl.Vector2) bool {
		res++
		return false
	})

	game := BuildGame().Build()
	scene1 := BuildScene("s1", game).Build()

	scene1.AddChild(a2)
	game.GoToScene("s1")

	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})
	game.handleMouseLeftRelease(rl.Vector2{X: 100, Y: 100})
	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})

	if res != expected {
		t.Fatalf("Call after release failed actual %d, expected %d", res, expected)
	}

	log.Output(1, "[PASS]: TestClickCallsHandlerAfterReleased")
}

func TestClickEventBubble(t *testing.T) {
	expected := "Bottom Clicked"
	res := ""

	a1 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 1).Build()
	a1.SetMouseDownHandler(func(mousePos rl.Vector2) bool {
		res = expected
		return false
	})

	a2 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 2).Build()
	a2.SetMouseDownHandler(func(mousePos rl.Vector2) bool {
		res = "Top Clicked"
		return true
	})

	game := BuildGame().Build()
	scene1 := BuildScene("s1", game).Build()

	scene1.AddChild(a1)
	scene1.AddChild(a2)
	game.GoToScene("s1")
	game.handleMouseLeftClick(rl.Vector2{X: 100, Y: 100})

	if res != expected {
		t.Fatalf("No Bubble Failed actual %s, expected %s", res, expected)
	}

	log.Output(1, "[PASS]: TestClickEventBubble")
}

func TestKeyEvent(t *testing.T) {
	expected := "Key Received"
	res := ""

	a1 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 1).Build()
	a1.SetKeyHandler(rl.KeyA, func() bool {
		res = expected
		return true
	})

	game := BuildGame().Build()
	scene1 := BuildScene("s1", game).Build()

	scene1.AddChild(a1)
	game.GoToScene("s1")
	game.handleKeyPressed(rl.KeyA)

	if res != expected {
		t.Fatalf("Key received failed - actual %s, expected %s", res, expected)
	}

	log.Output(1, "[PASS]: TestKeyEvent")
}

func TestKeyEventBubble(t *testing.T) {
	expected := []int{1, 2}
	res := make([]int, 2)

	a1 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 1).Build()
	a1.SetKeyHandler(rl.KeyA, func() bool {
		res[0] = 1
		return false
	})
	a2 := BuildActor().WithDimensions(200, 200).WithPosition(0, 0, 1).Build()
	a2.SetKeyHandler(rl.KeyA, func() bool {
		res[1] = 2
		return false
	})

	game := BuildGame().Build()
	scene1 := BuildScene("s1", game).Build()

	scene1.AddChild(a1)
	scene1.AddChild(a2)
	game.GoToScene("s1")
	game.handleKeyPressed(rl.KeyA)

	if tools.IndexOf(res, 1) == -1 || tools.IndexOf(res, 2) == -1 {
		t.Fatalf("Key Bubble failed - actual %v, expected %v", res, expected)
	}

	log.Output(1, "[PASS]: TestKeyEventBubble")
}
