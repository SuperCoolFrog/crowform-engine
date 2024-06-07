package components

import (
	pkg "crowform/pkg"
)

type VerticalListBuilder struct {
	spacing   float32
	listItems []*pkg.Actor
}

func BuildVerticalList() *VerticalListBuilder {
	return &VerticalListBuilder{
		listItems: make([]*pkg.Actor, 0),
	}
}

func (builder *VerticalListBuilder) With(listItem *pkg.Actor) *VerticalListBuilder {
	builder.listItems = append(builder.listItems, listItem)
	return builder
}

func (builder *VerticalListBuilder) Spacing(s float32) *VerticalListBuilder {
	builder.spacing = s
	return builder
}

func (builder *VerticalListBuilder) Build() *pkg.Actor {
	vlist := pkg.BuildActor().Build()

	for _, item := range builder.listItems {
		item.SetY(vlist.H())
		vlist.AddChild(item)
		vlist.SetHeight(vlist.H() + item.H() + builder.spacing)

		if vlist.W() < item.W() {
			vlist.SetWidth(item.W())
		}
	}

	return vlist
}
