package components

import (
	pkg "crowform/crw"
)

type HorizontalListBuilder struct {
	spacing   float32
	listItems []*pkg.Actor
}

func BuildHorizontalList() *HorizontalListBuilder {
	return &HorizontalListBuilder{
		listItems: make([]*pkg.Actor, 0),
	}
}

func (builder *HorizontalListBuilder) With(listItem *pkg.Actor) *HorizontalListBuilder {
	builder.listItems = append(builder.listItems, listItem)
	return builder
}

func (builder *HorizontalListBuilder) Spacing(s float32) *HorizontalListBuilder {
	builder.spacing = s
	return builder
}

func (builder *HorizontalListBuilder) Build() *pkg.Actor {
	vlist := pkg.BuildActor().Build()

	for _, item := range builder.listItems {
		item.SetX(float32(vlist.W()))
		vlist.AddChild(item)
		vlist.SetWidth(vlist.W() + item.W() + builder.spacing)

		if vlist.H() < item.H() {
			vlist.SetHeight(item.H())
		}
	}

	return vlist
}
