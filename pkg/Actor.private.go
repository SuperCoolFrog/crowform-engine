package pkg

import (
	"crowform/internal/tools"
	"sort"
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

func (actor *Actor) doActions(deltaTime time.Duration, allActions []ActorAction, idx int, onComplete func()) {
	if idx > len(allActions)-1 {
		onComplete()
		return
	}

	allActions[idx].do(deltaTime, actor, func() {
		actor.doActions(deltaTime, allActions, idx+1, onComplete)
	})
}

func (actor *Actor) ActionsSetAsReady() {
	actor.actorActionState = ActorActionState_READY
}

func (actor *Actor) ActionsSetAsProcessing() {
	actor.actorActionState = ActorActionState_PROCESSING
}

func (actor *Actor) ActionsSetAsStop() {
	actor.actorActionState = ActorActionState_STOP
}

func (actor *Actor) runActions(deltaTime time.Duration) {
	if actor.actorActionState != ActorActionState_READY {
		return
	}

	actor.ActionsSetAsProcessing()

	actions := tools.FilterSlice(actor.actions, func(a ActorAction) bool { return a.when(actor) })

	if len(actions) == 0 {
		actor.ActionsSetAsReady()
		return
	}

	sort.Slice(actions, func(i, j int) bool {
		return actions[i].index < actions[j].index
	})

	actor.doActions(deltaTime, actions, 0, func() {
		actor.ActionsSetAsReady()
	})
}

func (actor *Actor) getCollisionElement() rl.Rectangle {
	if !actor.CollisionElement.HasValue() {
		return actor.element
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

func (actor *Actor) GetWindowPosition() (position rl.Vector3) {
	position.X = actor.position.X
	position.Y = actor.position.Y
	position.Z = actor.position.Z

	if actor.parent != nil {
		parentPos := actor.parent.GetWindowPosition()

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
				return item.GetWindowPosition().Z > child.GetWindowPosition().Z
			})
		}
	}

	me.Children = nil

	tools.ForEach(nu, func(a *Actor) {
		me.AddChild(a)
	})
}
