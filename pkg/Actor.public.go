package pkg

import (
	"crowform/internal/tools"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/** Actor Methods: Public **/

/** -- Loop Func **/

func (actor *Actor) Update(deltaTime time.Duration) {
	actor.onUpdate(deltaTime)
	actor.runActions(deltaTime)
	actor.updateAnimations(deltaTime)

	tools.ForEach(actor.Children, func(child *Actor) {
		child.Update(deltaTime)
	})
}

func (actor *Actor) Draw() {
	actor.draw()

	tools.ForEach(actor.Sprites, func(s *Sprite) {
		s.Draw()
	})

	tools.ForEach(actor.Animations, func(a *Animation) {
		a.Draw()
	})

	tools.ForEach(actor.Children, func(child *Actor) {
		child.Draw()
	})
}

/** // Loop Func **/

// Updates Element Width, Height
func (actor *Actor) SetWidthHeight(width float32, height float32) {
	actor.element.Width = width
	actor.element.Height = height
}

func (actor *Actor) GetPosition() rl.Vector3 {
	return actor.position
}

// Updates Element and Position X
func (actor *Actor) SetX(x float32) {
	actor.element.X = x
	actor.position.X = x

	pos := actor.windowPosition()

	tools.ForEach(actor.Sprites, func(s *Sprite) {
		s.DestRect.X = pos.X
	})
}

// Updates Element and Position Y
func (actor *Actor) SetY(y float32) {
	actor.element.Y = y
	actor.position.Y = y

	pos := actor.windowPosition()

	tools.ForEach(actor.Sprites, func(s *Sprite) {
		s.DestRect.Y = pos.Y
	})
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
	if child.position.Z == 0 {
		// Set to max Z
		max := tools.MaxFloat32Of(
			tools.MapSlice(actor.Children, func(c *Actor) float32 {
				return c.windowPosition().Z
			})...)
		child.position.Z = max + 1
	}

	actor.Children = tools.InsertSorted(actor.Children, child,
		func(item *Actor) bool {
			return item.windowPosition().Z > child.windowPosition().Z
		})
	child.parent = actor
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
	position := actor.windowPosition()
	sprite.DestRect.X = position.X
	sprite.DestRect.Y = position.Y

	actor.Sprites = append(actor.Sprites, sprite)
}
func (me *Actor) RemoveSprite(spriteToRemove *Sprite) {
	me.Sprites = tools.RemoveAll(me.Sprites, spriteToRemove)
}

func (actor *Actor) HasParent() bool {
	return actor.parent != nil
}

func (actor *Actor) GetParent() *Actor {
	return actor.parent
}

func (actor *Actor) RemoveChild(child *Actor) {
	actor.Children = tools.Remove(actor.Children, child)
	child.parent = nil
}

func (actor *Actor) RemoveSelf() {
	if actor.parent == nil {
		return
	}

	actor.parent.RemoveChild(actor)
}

func (actor *Actor) ShowBorder() {
	actor.showBorder = true
}

func (actor *Actor) HideBorder() {
	actor.showBorder = false
}

func (actor *Actor) AddAnimation(animation *Animation) {
	position := actor.windowPosition()
	animation.DestRect.X += position.X
	animation.DestRect.Y += position.Y

	actor.Animations = append(actor.Animations, animation)
}

func (actor *Actor) RemoveAnimation(anim *Animation) {
	actor.Animations = tools.Remove(actor.Animations, anim)
}
