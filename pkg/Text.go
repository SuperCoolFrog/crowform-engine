package pkg

import (
	"crowform/internal/cache"
	"crowform/internal/tools"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	TextBuilder
	font           *rl.Font
	parent         *Actor
	queueForUpdate []func()
}

func (text *Text) update() {
	tools.ForEach(text.queueForUpdate, func(f func()) {
		f()
	})

	text.queueForUpdate = nil
}

func (text *Text) draw() {
	pos := text.GetWindowPos()

	rl.DrawTextEx(text.GetFont(), text.text, pos, text.fontSize, 0, text.color)
}

func (text *Text) GetTextSize() rl.Vector2 {
	font := text.GetFont()
	return rl.MeasureTextEx(font, text.text, text.fontSize, text.spacing)
}

func (text *Text) GetFont() rl.Font {
	if text.font == nil {
		font := cache.GetFont(text.fontFileName, int32(math.Trunc(float64(text.fontSize))))
		text.font = &font
	}

	return *text.font
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

func (text *Text) vAlignCenter() {
	if text.parent == nil {
		return
	}

	parentEle := text.parent.GetElement()
	size := text.GetTextSize()

	text.position.Y = parentEle.Height/2 - size.Y/2
}
func (text *Text) VAlignCenter() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.vAlignCenter()
	})
}

func (text *Text) vAlignTop() {
	if text.parent == nil {
		return
	}

	text.position.Y = 0
}
func (text *Text) VAlignTop() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.vAlignTop()
	})
}

func (text *Text) vAlignBottom() {
	if text.parent == nil {
		return
	}

	parentEle := text.parent.GetElement()
	size := text.GetTextSize()

	text.position.Y = parentEle.Height - size.Y
}
func (text *Text) VAlignBottom() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.vAlignBottom()
	})
}

func (text *Text) hAlignCenter() {
	if text.parent == nil {
		return
	}

	parentEle := text.parent.GetElement()
	size := text.GetTextSize()

	text.position.X = parentEle.Width/2 - size.X/2
}
func (text *Text) HAlignCenter() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.hAlignCenter()
	})
}

func (text *Text) hAlignLeft() {
	if text.parent == nil {
		return
	}

	text.position.X = 0
}
func (text *Text) HAlignLeft() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.hAlignLeft()
	})
}

func (text *Text) hAlignRight() {
	if text.parent == nil {
		return
	}

	parentEle := text.parent.GetElement()
	size := text.GetTextSize()

	text.position.X = parentEle.Width - size.X
}
func (text *Text) HAlignRight() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.hAlignRight()
	})
}

func (text *Text) AlignCenterVH() {
	text.queueForUpdate = append(text.queueForUpdate, func() {
		text.hAlignCenter()
		text.vAlignCenter()
	})
}
