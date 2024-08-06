package crw

import (
	"crowform/internal/tools"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCENE_MOUSE_ZINDEX = 1000
)

type SceneUniqId string

type Scene struct {
	Actor

	SceneId      SceneUniqId
	parentGame   *Game
	paused       bool
	onStart      func(scene *Scene)
	onEnd        func(scene *Scene)
	mousePointer *Actor
}

type SceneBuilder struct {
	parentGame *Game
	sceneId    SceneUniqId
	onStart    func(scene *Scene)
	onEnd      func(scene *Scene)
}

func BuildScene(sceneId SceneUniqId, game *Game) *SceneBuilder {
	return &SceneBuilder{
		sceneId:    sceneId,
		parentGame: game,
		onStart:    func(scene *Scene) {},
		onEnd:      func(scene *Scene) {},
	}
}

func (builder *SceneBuilder) WithOnStart(onStart func(scene *Scene)) *SceneBuilder {
	builder.onStart = onStart
	return builder
}

func (builder *SceneBuilder) WithOnEnd(onEnd func(scene *Scene)) *SceneBuilder {
	builder.onEnd = onEnd
	return builder
}

func (builder *SceneBuilder) Build() *Scene {
	sceneActor := BuildActor().
		WithDimensions(float32(builder.parentGame.windowWidth), float32(builder.parentGame.windowHeight)).
		WithPosition(0, 0, 0).
		Build()

	nuScene := &Scene{
		Actor:      *sceneActor,
		SceneId:    builder.sceneId,
		parentGame: builder.parentGame,
		paused:     false,
		onStart:    builder.onStart,
		onEnd:      builder.onEnd,
	}

	builder.parentGame.AddScene(nuScene)

	return nuScene
}

func (scene *Scene) Pause() {
	scene.paused = true

}
func (scene *Scene) Unpause() {
	scene.paused = false
}

func (scene *Scene) Update(deltaTime time.Duration) {
	if !scene.paused {
		scene.Actor.Update(deltaTime)
		return
	}

	actorsToUpdate := scene.QueryAny([]QueryAttribute{
		queryAttribute_UPDATES_WHEN_PAUSED,
	})

	tools.ForEach(actorsToUpdate, func(actor *Actor) {
		actor.Update(deltaTime)
	})
}

func (scene *Scene) Start() {
	scene.onStart(scene)
}
func (scene *Scene) End() {
	scene.onEnd(scene)
}

func (scene *Scene) ChangeMouseTexture(sprite *Sprite) {
	if scene.mousePointer != nil {
		scene.RemoveChild(scene.mousePointer)
	}

	mouse := BuildActor().
		WithPosition(50, 50, SCENE_MOUSE_ZINDEX).
		WithDimensions(sprite.srcRect.Width, sprite.srcRect.Height).
		WithOnUpdate(func(deltaTime time.Duration) {
			if scene.mousePointer == nil {
				return
			}

			rl.HideCursor()

			pos := rl.GetMousePosition()
			// 	me.mouseSprite.Element.Left = float64(mouseX) - settings.Settings.CameraXOffset
			// 	me.mouseSprite.Element.Top = float64(mouseY) - settings.Settings.CameraYOffset
			scene.mousePointer.SetX(pos.X)
			scene.mousePointer.SetY(pos.Y)
		}).
		WithAllowUpdateDuringPause().
		Build()
	mouse.AddSprite(sprite)

	scene.mousePointer = mouse
	scene.AddChild(mouse)
}

func (scene *Scene) GetGame() *Game {
	return scene.parentGame
}
