package crw

import (
	"crowform/internal/tools"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/** Actor Methods: Private **/

func (actor *Actor) draw() {
	if actor.showBorder {
		rect := actor.element
		pos := actor.GetWindowPosition()
		rect.X = pos.X
		rect.Y = pos.Y

		intRect := rect.ToInt32()

		rl.DrawRectangleLines(intRect.X, intRect.Y, intRect.Width, intRect.Height, actor.borderColor)
	}
}

func (actor *Actor) updateAnimations(deltaTime time.Duration) {
	for i := range actor.Animations {
		if i >= len(actor.Animations) {
			return
		}
		anim := actor.Animations[i]
		anim.update(deltaTime)
	}
}
func (actor *Actor) updateSprites(deltaTime time.Duration) {
	for i := range actor.Sprites {
		if i >= len(actor.Sprites) {
			return
		}
		sprite := actor.Sprites[i]
		sprite.update(deltaTime)
	}
}

func (actor *Actor) runUpdateQueue() {
	tools.ForEach(actor.queueForUpdate, func(f func()) {
		f()
	})

	actor.queueForUpdate = nil
}

func (actor *Actor) getCollisionElement() rl.Rectangle {
	if !actor.CollisionElement.HasValue() {
		return actor.GetWindowRec()
	}

	collisionEl := actor.CollisionElement.Value

	pos := actor.GetWindowPosition()

	e := rl.Rectangle{
		X:      pos.X + collisionEl.X,
		Y:      pos.Y + collisionEl.Y,
		Width:  collisionEl.Width,
		Height: collisionEl.Height,
	}

	return e
}

func (me *Actor) resortChildrenByZ() {
	nu := make([]*Actor, 0)

	for i := range me.Children {
		if i < len(me.Children) {
			child := me.Children[i]

			if i == 0 {
				nu = append(nu, child)
				continue
			}

			nu = tools.InsertSorted(nu, child, func(item *Actor) bool {
				return item.GetWindowPosition().Z > child.GetWindowPosition().Z
			})
		}
	}

	me.Children = nil

	tools.ForEach(nu, func(a *Actor) {
		me.AddChild(a)
	})
}

func (actor *Actor) addChild(child *Actor) {
	if child.position.Z == 0 {
		var minNext float32 = 1

		tools.ForEach(actor.Children, func(a *Actor) {
			next := a.GetWindowPosition().Z
			if minNext < next {
				if next >= SCENE_MOUSE_ZINDEX {
					return
				} else {
					minNext = next
				}
			}
		})

		child.position.Z = minNext + 1
	}

	actor.Children = tools.InsertSorted(actor.Children, child,
		func(item *Actor) bool {
			return item.GetWindowPosition().Z > child.GetWindowPosition().Z
		})
	child.parent = actor
	child.onParentAdded(actor)
}
