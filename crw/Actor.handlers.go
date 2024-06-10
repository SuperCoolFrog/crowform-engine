package crw

import rl "github.com/gen2brain/raylib-go/raylib"

func (actor *Actor) SetMouseDownHandler(handler func(mousePos rl.Vector2) bool) {
	actor.AddQueryAttr(queryAttribute_RECEIVES_MOUSE_DOWN_EVENT)
	actor.events.onMouseDown = handler
}
func (actor *Actor) UnsetMouseDownHandler() {
	actor.RemoveQueryAttr(queryAttribute_RECEIVES_MOUSE_DOWN_EVENT)
	actor.events.onMouseDown = func(mousePos rl.Vector2) bool { return true }
}

func (actor *Actor) SetMouseUpHandler(handler func(mousePos rl.Vector2) bool) {
	actor.AddQueryAttr(queryAttribute_RECEIVES_MOUSE_UP_EVENT)
	actor.events.onMouseUp = handler
}
func (actor *Actor) UnsetMouseUpHandler() {
	actor.RemoveQueryAttr(queryAttribute_RECEIVES_MOUSE_UP_EVENT)
	actor.events.onMouseUp = func(mousePos rl.Vector2) bool { return true }
}

func (actor *Actor) SetMouseMoveHandler(handler func(mousePos rl.Vector2) bool) {
	actor.AddQueryAttr(queryAttribute_RECEIVES_MOUSE_MOVE_EVENT)
	actor.events.onMouseMove = handler
}
func (actor *Actor) UnsetMouseMoveHandler() {
	actor.RemoveQueryAttr(queryAttribute_RECEIVES_MOUSE_MOVE_EVENT)
	actor.events.onMouseMove = func(mousePos rl.Vector2) bool { return false }
}

func (actor *Actor) SetMouseEnterHandler(handler func()) {
	actor.AddQueryAttr(queryAttribute_RECEIVES_MOUSE_ENTER_EVENT)
	actor.events.onMouseEnter = handler
}
func (actor *Actor) UnsetMouseEnterHandler() {
	actor.RemoveQueryAttr(queryAttribute_RECEIVES_MOUSE_ENTER_EVENT)
	actor.events.onMouseEnter = func() {}
}

func (actor *Actor) SetMouseExitHandler(handler func()) {
	actor.AddQueryAttr(queryAttribute_RECEIVES_MOUSE_EXIT_EVENT)
	actor.events.onMouseExit = handler
}

func (actor *Actor) UnsetMouseExitHandler() {
	actor.RemoveQueryAttr(queryAttribute_RECEIVES_MOUSE_EXIT_EVENT)
	actor.events.onMouseExit = func() {}
}

/*
	Key = rl.Key**
*/
func (actor *Actor) SetKeyHandler(key int32, handler func() bool) {
	actor.AddQueryAttr(queryAttribute_RECEIVES_KEY_EVENT)
	actor.events.onKeyPressed[key] = handler
}

/*
	Key = rl.Key**
*/
func (actor *Actor) UnsetKeyHandler(key int32, queryAttr QueryAttribute) {
	if len(actor.events.onKeyPressed) == 1 {
		actor.RemoveQueryAttr(queryAttribute_RECEIVES_KEY_EVENT)
	}
	delete(actor.events.onKeyPressed, key)
}

/*
	Key = rl.Key**
*/
func (actor *Actor) TriggerKeyHandler(key int32) bool {
	if !actor.IsQryType(queryAttribute_RECEIVES_KEY_EVENT) {
		return false
	}

	handler, found := actor.events.onKeyPressed[key]
	if !found {
		return false
	}

	return handler()
}

func (actor *Actor) SetHandler(queryAttr QueryAttribute, handler func(params interface{})) {
	actor.AddQueryAttr(queryAttr)
	actor.events.defined[queryAttr] = handler
}

func (actor *Actor) UnsetHandler(queryAttr QueryAttribute) {
	actor.RemoveQueryAttr(queryAttr)
	delete(actor.events.defined, queryAttr)
}

func (actor *Actor) TriggerHandler(queryAttr QueryAttribute, params interface{}) {
	if !actor.IsQryType(queryAttr) {
		return
	}

	handler, found := actor.events.defined[queryAttr]
	if !found {
		return
	}

	handler(params)
}
