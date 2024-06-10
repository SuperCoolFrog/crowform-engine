package crw

import (
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) checkInputEvents() {
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		game.handleMouseLeftClick(mousePos)
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		game.handleMouseLeftRelease(mousePos)
	}

	keyPressed := rl.GetKeyPressed()
	// 0 means none
	if keyPressed != 0 {
		game.handleKeyPressed(keyPressed)
	}

	game.handleMouseEnter(mousePos)
	game.handleMouseExit(mousePos)
	game.handleMouseMove(mousePos)
}

var lastMouseDownActor *Actor = nil

func (game *Game) handleMouseLeftClick(mousePos rl.Vector2) {
	if game.currentScene == nil {
		return
	}
	if lastMouseDownActor != nil {
		return
	}

	actors := game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_MOUSE_DOWN_EVENT})
	clickActors := make([]*Actor, 0)

	for _, actor := range actors {
		if rl.CheckCollisionPointRec(mousePos, actor.getCollisionElement()) {
			clickActors = append(clickActors, actor)
		}
	}

	sort.Slice(clickActors, func(i, j int) bool {
		// Sort descending
		return clickActors[i].GetWindowPosition().Z > clickActors[j].GetWindowPosition().Z
	})

	for _, actor := range clickActors {
		// If click handler returns false then break, otherwise bubble
		if !actor.events.onMouseDown(mousePos) {
			lastMouseDownActor = actor
			break
		}
	}
}

func (game *Game) handleMouseLeftRelease(mousePos rl.Vector2) {
	if game.currentScene == nil {
		return
	}

	actors := game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_MOUSE_UP_EVENT})
	clickActors := make([]*Actor, 0)

	for _, actor := range actors {
		if rl.CheckCollisionPointRec(mousePos, actor.getCollisionElement()) {
			clickActors = append(clickActors, actor)
		}
	}

	sort.Slice(clickActors, func(i, j int) bool {
		// Sort descending
		return clickActors[i].GetWindowPosition().Z > clickActors[j].GetWindowPosition().Z
	})

	if lastMouseDownActor != nil && len(clickActors) > 0 {
		if lastMouseDownActor == clickActors[0] {
			lastMouseDownActor.events.onMouseUp(mousePos)
			lastMouseDownActor = nil
			return
		} else {
			lastMouseDownActor = nil
		}
	}

	for _, actor := range clickActors {
		// If click handler returns false then break, otherwise bubble
		if !actor.events.onMouseUp(mousePos) {
			break
		}
	}

	lastMouseDownActor = nil
}

var mouseOverTarget *Actor

func (game *Game) handleMouseEnter(mousePos rl.Vector2) {
	if game.currentScene == nil {
		return
	}

	if mouseOverTarget != nil {
		return
	}

	actors := game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_MOUSE_ENTER_EVENT})
	mouseActors := make([]*Actor, 0)

	for _, actor := range actors {
		if rl.CheckCollisionPointRec(mousePos, actor.getCollisionElement()) {
			mouseActors = append(mouseActors, actor)
		}
	}

	if len(mouseActors) == 0 {
		actors = game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_MOUSE_EXIT_EVENT})

		for _, actor := range actors {
			if rl.CheckCollisionPointRec(mousePos, actor.getCollisionElement()) {
				mouseActors = append(mouseActors, actor)
			}
		}

		if len(mouseActors) > 0 {
			mouseOverTarget = mouseActors[0]
		}

		return
	}

	sort.Slice(mouseActors, func(i, j int) bool {
		// Sort descending
		return mouseActors[i].GetWindowPosition().Z > mouseActors[j].GetWindowPosition().Z
	})

	mouseOverTarget = mouseActors[0]

	mouseOverTarget.events.onMouseEnter()
}

func (game *Game) handleMouseExit(mousePos rl.Vector2) {
	if mouseOverTarget == nil {
		return
	}

	if !rl.CheckCollisionPointRec(mousePos, mouseOverTarget.getCollisionElement()) {
		mouseOverTarget.events.onMouseExit()
		mouseOverTarget = nil
	}
}

func (game *Game) handleMouseMove(mousePos rl.Vector2) {
	if game.currentScene == nil {
		return
	}

	actors := game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_MOUSE_MOVE_EVENT})

	for i := range actors {
		if i >= len(actors) {
			break
		}

		actor := actors[i]

		if actor.events.onMouseMove(mousePos) {
			return
		}
	}
}

func (game *Game) handleKeyPressed(key int32) {
	if game.currentScene == nil {
		return
	}

	actors := game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_KEY_EVENT})

	handled := false
	i := 0

	for !handled {
		if i >= len(actors) {
			return
		}

		handled = actors[i].TriggerKeyHandler(key)

		i++
	}
}
