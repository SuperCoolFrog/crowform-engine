package pkg

import (
	"crowform/internal/tools"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ActorActionState string

const (
	ActorActionState_INIT       ActorActionState = "ActorActionState_INIT"
	ActorActionState_READY      ActorActionState = "ActorActionState_READY"
	ActorActionState_PROCESSING ActorActionState = "ActorActionState_PROCESSING"
	ActorActionState_STOP       ActorActionState = "ActorActionState_STOP"
)

type ActorInitAction = func(actor *Actor)
type ActorWhenCondition = func(actor *Actor) bool
type ActorActionDone = func()
type ActorDoAction = func(deltaTime time.Duration, actor *Actor, done ActorActionDone)
type ActorAction struct {
	index int
	when  ActorWhenCondition
	do    ActorDoAction
}

type ActorEventHandlers struct {
	onMouseDown  func(mousePos rl.Vector2) bool // True to bubble, otherwise false
	onMouseUp    func(mousePos rl.Vector2) bool // True to bubble, otherwise false
	onMouseMove  func(mousePos rl.Vector2) bool // True if handled, false otherwise
	onMouseEnter func()
	onMouseExit  func()
	onKeyPressed map[int32]func() bool // true if handled
	defined      map[QueryAttribute]func(params interface{})
}

type Actor struct {
	Children         []*Actor
	Sprites          []*Sprite
	QueryAttributes  []QueryAttribute
	CollisionElement tools.Maybe[rl.Rectangle]
	Scene            *Scene
	Color            rl.Color

	parent   *Actor
	element  rl.Rectangle
	position rl.Vector3
	onUpdate func(deltaTime time.Duration)

	events ActorEventHandlers

	actorActionState ActorActionState
	actions          []ActorAction

	subscribers map[string][]func()
}

/** Builder Methods **/

type ActorBuilder struct {
	actor       Actor
	ignorePause bool
	hasActions  bool
}

func BuildActor() *ActorBuilder {
	return &ActorBuilder{
		actor: Actor{
			Children:        make([]*Actor, 0),
			QueryAttributes: make([]QueryAttribute, 0),

			position: rl.Vector3{},
			element:  rl.Rectangle{},
			onUpdate: func(deltaTime time.Duration) {},
			Color:    rl.Black,
			events: ActorEventHandlers{
				onMouseDown:  func(mousePos rl.Vector2) bool { return true },
				onMouseUp:    func(mousePos rl.Vector2) bool { return true },
				onMouseMove:  func(mousePos rl.Vector2) bool { return false },
				onMouseEnter: func() {},
				onMouseExit:  func() {},
				defined:      make(map[QueryAttribute]func(params interface{})),
				onKeyPressed: make(map[int32]func() bool),
			},
			actions:          nil,
			actorActionState: ActorActionState_INIT,
			subscribers:      map[string][]func(){},
		},
		ignorePause: false,
		hasActions:  false,
	}
}

func (builder *ActorBuilder) WithPosition(x float32, y float32, z float32) *ActorBuilder {
	builder.actor.element.X = x
	builder.actor.element.Y = y
	builder.actor.position.X = x
	builder.actor.position.Y = y
	builder.actor.position.Z = z
	return builder
}
func (builder *ActorBuilder) WithDimensions(width float32, height float32) *ActorBuilder {
	builder.actor.element.Width = width
	builder.actor.element.Height = height
	return builder
}
func (builder *ActorBuilder) WithOnUpdate(onUpdate func(deltaTime time.Duration)) *ActorBuilder {
	builder.actor.onUpdate = onUpdate
	return builder
}
func (builder *ActorBuilder) WithColor(color rl.Color) *ActorBuilder {
	builder.actor.Color = color
	return builder
}
func (builder *ActorBuilder) WithAllowUpdateDuringPause() *ActorBuilder {
	builder.actor.AddQueryAttr(queryAttribute_UPDATES_WHEN_PAUSED)
	return builder
}

func (builder *ActorBuilder) WithAction(when ActorWhenCondition, do ActorDoAction) *ActorBuilder {
	if !builder.hasActions {
		builder.actor.actions = make([]ActorAction, 0)
		builder.hasActions = true
	}

	builder.actor.actions = append(builder.actor.actions, ActorAction{len(builder.actor.actions), when, do})
	return builder
}

func (builder *ActorBuilder) Build() *Actor {
	return &builder.actor
}
