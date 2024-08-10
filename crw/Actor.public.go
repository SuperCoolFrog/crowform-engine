package crw

import (
	"crowform/internal/tools"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/** Actor Methods: Public **/

/** -- Loop Func **/

func (actor *Actor) Update(deltaTime time.Duration) {
	actor.runUpdateQueue()
	actor.onUpdate(deltaTime)
	actor.runActions(deltaTime)
	actor.updateAnimations(deltaTime)
	actor.updateSprites(deltaTime)

	tools.ForEach(actor.Children, func(child *Actor) {
		child.Update(deltaTime)
	})
	tools.ForEach(actor.Texts, func(t *Text) {
		t.update()
	})
}

func (actor *Actor) Draw() {
	actor.draw()

	tools.ForEach(actor.Sprites, func(s *Sprite) {
		s.draw()
	})

	tools.ForEach(actor.Texts, func(t *Text) {
		t.draw()
	})

	tools.ForEach(actor.Animations, func(a *Animation) {
		a.Draw()
	})

	tools.ForEach(actor.Children, func(child *Actor) {
		child.Draw()
	})
}

/** // Loop Func **/

func (actor *Actor) GetElement() rl.Rectangle {
	return actor.element
}

// Shorthand for actor.GetElement().Height
func (actor *Actor) H() float32 {
	return actor.element.Height
}

// Shorthand for actor.GetElement().Width
func (actor *Actor) W() float32 {
	return actor.element.Width
}

// Shorthand for actor.GetPosition().X
func (actor *Actor) X() float32 {
	return actor.GetPosition().X
}

// Shorthand for actor.GetPosition().Y
func (actor *Actor) Y() float32 {
	return actor.GetPosition().Y
}
func (actor *Actor) Z() float32 {
	return actor.GetPosition().Z
}

// Updates Element Width
func (actor *Actor) SetWidth(width float32) {
	actor.element.Width = width
}

// Updates Element Width
func (actor *Actor) SetHeight(height float32) {
	actor.element.Height = height
}

// Updates Element Width, Height
func (actor *Actor) SetWidthHeight(width float32, height float32) {
	actor.SetWidth(width)
	actor.SetHeight(height)
}

func (actor *Actor) GetPosition() rl.Vector3 {
	return actor.position
}

// Updates Element and Position X
func (actor *Actor) SetX(x float32) {
	actor.element.X = x
	actor.position.X = x
}

// Updates Element and Position Y
func (actor *Actor) SetY(y float32) {
	actor.element.Y = y
	actor.position.Y = y
}

// Updates Position Z
func (actor *Actor) SetZ(z float32) {
	actor.position.Z = z

	if actor.parent != nil {
		actor.parent.resortChildrenByZ()
	}
}

// Updates Element and Position X,Y
func (actor *Actor) SetXY(x float32, y float32) {
	actor.SetX(x)
	actor.SetY(y)
}

// Updates Element X,Y and Position X,Y,Z
func (actor *Actor) SetXYZ(x float32, y float32, z float32) {
	actor.element.X = x
	actor.position.X = x
	actor.element.Y = y
	actor.position.Y = y
	actor.position.Z = z
}

func (actor *Actor) AddChild(child *Actor) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		actor.addChild(child)
	})
}

func (actor *Actor) SetOnParentAdded(handler func(parent *Actor)) {
	actor.onParentAdded = handler
}

func (me *Actor) Intersects(other *Actor) bool {
	elA := me.getCollisionElement()
	elB := other.getCollisionElement()

	return rl.CheckCollisionRecs(elA, elB)
}

func (actor *Actor) AddQueryAttr(qryAttr QueryAttribute) {
	if tools.IndexOf(actor.QueryAttributes, qryAttr) == -1 {
		actor.QueryAttributes = append(actor.QueryAttributes, qryAttr)
	}
}
func (actor *Actor) RemoveQueryAttr(qryAttr QueryAttribute) {
	actor.QueryAttributes = tools.Remove(actor.QueryAttributes, qryAttr)
}

// Returns actor and children that match AT LEASE ONE of the QueryAttributes
func (actor *Actor) QueryAny(qryAttrs []QueryAttribute) []*Actor {
	res := make([]*Actor, 0)

	i := tools.GetIntersects(qryAttrs, actor.QueryAttributes)

	if len(i) > 0 {
		res = append(res, actor)
	}

	for i := range actor.Children {
		if i < len(actor.Children) {
			child := actor.Children[i]
			res = append(res, child.QueryAny(qryAttrs)...)
		}
	}

	return res
}

// Returns actor and children that match ALL of the QueryAttributes
func (actor *Actor) QueryExact(qryAttrs []QueryAttribute) []*Actor {
	res := make([]*Actor, 0)

	i := tools.GetIntersects(qryAttrs, actor.QueryAttributes)

	if len(i) == len(qryAttrs) {
		res = append(res, actor)
	}

	for i := range actor.Children {
		if i < len(actor.Children) {
			child := actor.Children[i]
			res = append(res, child.QueryExact(qryAttrs)...)
		}
	}

	return res
}

func (actor *Actor) IsQryType(qryType QueryAttribute) bool {
	return tools.IndexOf(actor.QueryAttributes, qryType) > -1
}

func (actor *Actor) AddSprite(sprite *Sprite) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		sprite.setParent(actor)
		actor.Sprites = append(actor.Sprites, sprite)
	})
}
func (actor *Actor) RemoveSprite(spriteToRemove *Sprite) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		actor.Sprites = tools.RemoveAll(actor.Sprites, spriteToRemove)
	})
}

func (actor *Actor) AddText(text *Text) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		text.SetParent(actor)
		actor.Texts = append(actor.Texts, text)
	})
}

func (actor *Actor) RemoveText(textToRemove *Text) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		actor.Texts = tools.RemoveAll(actor.Texts, textToRemove)
	})
}

func (actor *Actor) HasParent() bool {
	return actor.parent != nil
}

func (actor *Actor) GetParent() *Actor {
	return actor.parent
}

func (actor *Actor) RemoveChild(child *Actor) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		actor.removeChild(child)
	})
}
func (actor *Actor) removeChild(child *Actor) {
	actor.Children = tools.Remove(actor.Children, child)
	child.runUpdateQueue()
	child.parent = nil
}

func (actor *Actor) RemoveSelf() {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		actor.removeSelf()
	})
}

func (actor *Actor) removeSelf() {
	if actor.parent == nil {
		return
	}

	actor.parent.removeChild(actor)
}

func (actor *Actor) ShowBorder() {
	actor.showBorder = true
}

func (actor *Actor) HideBorder() {
	actor.showBorder = false
}

func (actor *Actor) AddAnimation(animation *Animation) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		animation.SetParent(actor)
		actor.Animations = append(actor.Animations, animation)
	})
}

func (actor *Actor) RemoveAnimation(anim *Animation) {
	actor.queueForUpdate = append(actor.queueForUpdate, func() {
		actor.Animations = tools.Remove(actor.Animations, anim)
	})
}

func (actor *Actor) SetBorderColor(color rl.Color) {
	actor.borderColor = color
}

func (actor *Actor) ContainsPoint(point rl.Vector2) bool {
	return rl.CheckCollisionPointRec(point, actor.getCollisionElement())
}

func (actor *Actor) CollidesWithOther(other *Actor) bool {
	return rl.CheckCollisionRecs(other.getCollisionElement(), actor.getCollisionElement())
}

func (actor *Actor) CollidesWithRec(rec rl.Rectangle) bool {
	return rl.CheckCollisionRecs(actor.getCollisionElement(), rec)
}

func (actor *Actor) GetWindowRec() rl.Rectangle {
	rect := actor.element
	winPos := actor.GetWindowPosition()

	rect.X = winPos.X
	rect.Y = winPos.Y

	return rect
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
