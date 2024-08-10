package crw

import (
	"log"
	"testing"
	"time"
)

// Text
func TestWindowPositionForChildTextZero(t *testing.T) {
	a1 := BuildActor().WithPosition(0, 0, 0).Build()
	text := BuildText().Build()

	a1.AddText(text)

	if text.GetWindowPos().X != 0 {
		t.Fatalf("Expected child text X to remain 0, actual %f", text.GetWindowPos().X)
	}
	if text.GetWindowPos().Y != 0 {
		t.Fatalf("Expected child text Y to remain 0, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildTextZero")
}

func TestWindowPositionForChildTextWhenAdded(t *testing.T) {
	a1 := BuildActor().WithPosition(10, 20, 0).Build()
	text := BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	a1.Update(time.Second)

	if text.GetWindowPos().X != 15 {
		t.Fatalf("Expected child text X to be 15, actual %f", text.GetWindowPos().X)
	}
	if text.GetWindowPos().Y != 30 {
		t.Fatalf("Expected child text Y to be 30, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildTextWhenAdded")
}

func TestWindowPositionForChildTextWhenUpdated(t *testing.T) {
	a1 := BuildActor().WithPosition(0, 0, 0).Build()
	text := BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	a1.SetX(10)
	a1.SetY(20)
	a1.Update(time.Second)

	if text.GetWindowPos().X != 15 {
		t.Fatalf("Expected child text X to be 15, actual %f", text.GetWindowPos().X)
	}
	if text.GetWindowPos().Y != 30 {
		t.Fatalf("Expected child text Y to be 30, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildTextWhenUpdated")
}

func TestTextVAlignTop(t *testing.T) {
	a1 := BuildActor().WithPosition(10, 20, 0).Build()
	text := BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	a1.Update(time.Second)
	text.VAlignTop()
	text.update()

	if text.GetWindowPos().Y != 20 {
		t.Fatalf("Expected v align top to have text Y to be 20, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestTextVAlignTop")
}

func TestTextHAlignLeft(t *testing.T) {
	a1 := BuildActor().WithPosition(10, 20, 0).Build()
	text := BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	a1.Update(time.Second)
	text.HAlignLeft()
	text.update()

	if text.GetWindowPos().X != 10 {
		t.Fatalf("Expected v align top to have text X to be 10, actual %f", text.GetWindowPos().X)
	}

	log.Output(1, "[PASS]: TestTextHAlignLeft")
}

func TestPixelSizing(t *testing.T) {
	var expected float32 = 16
	var actual = ToClosestFontSize(13.23)
	if expected != actual {
		t.Fatalf("Expected 13.23 to return 16, actual %f", actual)
	}

	expected = 16
	actual = ToClosestFontSize(17.55)
	if expected != actual {
		t.Fatalf("Expected 17.55 to return 16, actual %f", actual)
	}

	expected = 48
	actual = ToClosestFontSize(45)
	if expected != actual {
		t.Fatalf("Expected 45 to return 48, actual %f", actual)
	}

	expected = 48
	actual = ToClosestFontSize(44)
	if expected != actual {
		t.Fatalf("Expected 44 to return 48, actual %f", actual)
	}

	log.Output(1, "[PASS]: TestPixelSizing")
}
