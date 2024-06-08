package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextBuilder struct {
	fontFileName string
	position     rl.Vector2
	color        rl.Color
	text         string
	fontSize     float32
	spacing      float32
}

func BuildText() *TextBuilder {
	return &TextBuilder{
		fontFileName: "",
		text:         "",
		position:     rl.Vector2{X: 0, Y: 0},
		color:        rl.White,
		fontSize:     16,
		spacing:      1,
	}
}

func (builder *TextBuilder) WithFont(filename string) *TextBuilder {
	builder.fontFileName = filename
	return builder
}
func (builder *TextBuilder) WithPosition(x, y float32) *TextBuilder {
	builder.position = rl.Vector2{X: x, Y: y}
	return builder
}
func (builder *TextBuilder) WithText(text string) *TextBuilder {
	builder.text = text
	return builder
}
func (builder *TextBuilder) WithColor(color rl.Color) *TextBuilder {
	builder.color = color
	return builder
}
func (builder *TextBuilder) WithFontSize(fontSize float32) *TextBuilder {
	builder.fontSize = ToClosestFontSize(fontSize)
	return builder
}
func (builder *TextBuilder) WithSpacing(spacing float32) *TextBuilder {
	builder.spacing = spacing
	return builder
}

func (builder *TextBuilder) Build() *Text {
	return &Text{
		TextBuilder: *builder,
	}
}
