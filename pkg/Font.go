package pkg

import (
	"crowform/internal/cache"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	TextBuilder
	font   *rl.Font
	parent *Actor
}

func (text *Text) draw() {
	if text.font == nil {
		font := cache.GetFont(text.fontFileName, int32(math.Trunc(float64(text.fontSize))))
		text.font = &font
	}

	pos := text.GetWindowPos()

	rl.DrawTextEx(*text.font, text.text, pos, text.fontSize, 0, text.color)
}

func (text *Text) SetParent(parent *Actor) {
	text.parent = parent
}
func (text *Text) SetColor(color rl.Color) {
	text.color = color
}

func (text *Text) GetPosition() rl.Vector2 {
	return text.position
}
func (text *Text) SetPosition(position rl.Vector2) {
	text.position = position
}

func (text *Text) SetText(textString string) {
	text.text = textString
}

func (text *Text) GetWindowPos() rl.Vector2 {
	if text.parent == nil {
		return text.position
	}

	parentPos := text.parent.GetWindowPosition()
	windowDestRec := text.position
	windowDestRec.X += parentPos.X
	windowDestRec.Y += parentPos.Y

	return windowDestRec
}
