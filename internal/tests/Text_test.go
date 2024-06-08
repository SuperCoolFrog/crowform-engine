package tests

import (
	"crowform/pkg"
	"log"
	"testing"
)

// Text
func TestWindowPositionForChildTextZero(t *testing.T) {
	a1 := pkg.BuildActor().WithPosition(0, 0, 0).Build()
	text := pkg.BuildText().Build()

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
	a1 := pkg.BuildActor().WithPosition(10, 20, 0).Build()
	text := pkg.BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)

	if text.GetWindowPos().X != 15 {
		t.Fatalf("Expected child text X to be 15, actual %f", text.GetWindowPos().X)
	}
	if text.GetWindowPos().Y != 30 {
		t.Fatalf("Expected child text Y to be 30, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildTextWhenAdded")
}

func TestWindowPositionForChildTextWhenUpdated(t *testing.T) {
	a1 := pkg.BuildActor().WithPosition(0, 0, 0).Build()
	text := pkg.BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	a1.SetX(10)
	a1.SetY(20)

	if text.GetWindowPos().X != 15 {
		t.Fatalf("Expected child text X to be 15, actual %f", text.GetWindowPos().X)
	}
	if text.GetWindowPos().Y != 30 {
		t.Fatalf("Expected child text Y to be 30, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestWindowPositionForChildTextWhenUpdated")
}

func TestTextVAlignTop(t *testing.T) {
	a1 := pkg.BuildActor().WithPosition(10, 20, 0).Build()
	text := pkg.BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	text.VAlignTop()

	if text.GetWindowPos().Y != 20 {
		t.Fatalf("Expected v align top to have text Y to be 20, actual %f", text.GetWindowPos().Y)
	}

	log.Output(1, "[PASS]: TestTextVAlignTop")
}

func TestTextHAlignLeft(t *testing.T) {
	a1 := pkg.BuildActor().WithPosition(10, 20, 0).Build()
	text := pkg.BuildText().WithPosition(5, 10).Build()

	a1.AddText(text)
	text.HAlignLeft()

	if text.GetWindowPos().X != 10 {
		t.Fatalf("Expected v align top to have text X to be 10, actual %f", text.GetWindowPos().X)
	}

	log.Output(1, "[PASS]: TestTextHAlignLeft")
}
