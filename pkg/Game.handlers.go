package pkg

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

func (game *Game) handleMouseLeftClick(mousePos rl.Vector2) {
	if game.currentScene == nil {
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
		return clickActors[i].windowPosition().Z > clickActors[j].windowPosition().Z
	})

	for _, actor := range clickActors {
		// If click handler returns false then break, otherwise bubble
		if !actor.events.onMouseDown(mousePos) {
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
		return clickActors[i].windowPosition().Z > clickActors[j].windowPosition().Z
	})

	for _, actor := range clickActors {
		// If click handler returns false then break, otherwise bubble
		if !actor.events.onMouseUp(mousePos) {
			break
		}
	}
}

var mouseOverTarget *Actor

func (game *Game) handleMouseEnter(mousePos rl.Vector2) {
	if game.currentScene == nil {
		return
	}

	old := mouseOverTarget

	// Checks if actor was removed
	if mouseOverTarget != nil && mouseOverTarget.parent == nil {
		old = nil
		mouseOverTarget = nil
	}

	actors := game.currentScene.QueryAny([]QueryAttribute{queryAttribute_RECEIVES_MOUSE_ENTER_EVENT})

	for i := range actors {

		if i >= len(actors) {
			break
		}

		actor := actors[i]
		if rl.CheckCollisionPointRec(mousePos, actor.getCollisionElement()) {
			if mouseOverTarget == nil {
				mouseOverTarget = actor
				break
			}

			zIndexCurrent := mouseOverTarget.windowPosition().Z
			zIndex := actor.windowPosition().Z

			if zIndex > zIndexCurrent {
				if mouseOverTarget.IsQryType(queryAttribute_RECEIVES_MOUSE_EXIT_EVENT) {
					mouseOverTarget.events.onMouseExit()
				}

				mouseOverTarget = actor

				break
			}
		}
	}

	if old == mouseOverTarget {
		return
	}

	if mouseOverTarget != nil && mouseOverTarget.events.onMouseEnter != nil {
		mouseOverTarget.events.onMouseEnter()
	}
}

func (game *Game) handleMouseExit(mousePos rl.Vector2) {
	if mouseOverTarget == nil {
		return
	}

	if rl.CheckCollisionPointRec(mousePos, mouseOverTarget.getCollisionElement()) {
		if mouseOverTarget.IsQryType(queryAttribute_RECEIVES_MOUSE_EXIT_EVENT) {
			mouseOverTarget.events.onMouseExit()
		}
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
