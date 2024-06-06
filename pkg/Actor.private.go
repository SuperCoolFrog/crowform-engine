package pkg

import (
	"crowform/internal/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/** Actor Methods: Private **/

func (actor *Actor) draw() {
	intRect := actor.element.ToInt32()
	rl.DrawRectangle(intRect.X, intRect.Y, intRect.Width, intRect.Height, actor.Color)
}

func (actor *Actor) getCollisionElement() rl.Rectangle {
	if !actor.CollisionElement.HasValue() {
		return actor.element
	}

	collisionEl := actor.CollisionElement.Value

	// @Todo window position may be needed
	pos := actor.element

	e := rl.Rectangle{
		X:      pos.X + collisionEl.X,
		Y:      pos.Y + collisionEl.Y,
		Width:  collisionEl.Width,
		Height: collisionEl.Height,
	}

	return e
}

func (actor *Actor) windowPosition() (position rl.Vector3) {
	position.X = actor.position.X
	position.Y = actor.position.Y
	position.Z = actor.position.Z

	if actor.parent != nil {
		parentPos := actor.parent.windowPosition()

		position.X += parentPos.X
		position.Y += parentPos.Y

		if position.Z > 0 || position.Z < 0 {
			position.Z = parentPos.Z + position.Z
		} else {
			position.Z = parentPos.Z + 1
		}
	}

	return position
}

func (me *Actor) resortChildrenByZ() {
	nu := make([]*Actor, len(me.Children))

	for i := range me.Children {
		if i < len(me.Children) {
			child := me.Children[i]

			if i == 0 {
				nu[0] = child
				continue
			}

			nu = tools.InsertSorted(me.Children, child, func(item *Actor) bool {
				return item.windowPosition().Z > child.windowPosition().Z
			})
		}
	}

	me.Children = nil

	tools.ForEach(nu, func(a *Actor) {
		me.AddChild(a)
	})
}
