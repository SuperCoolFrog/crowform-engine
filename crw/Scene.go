package crw

import (
	"crowform/internal/tools"
	"time"
)

type SceneUniqId string

type Scene struct {
	Actor

	SceneId    SceneUniqId
	parentGame *Game
	paused     bool
}

type SceneBuilder struct {
	parentGame *Game
	sceneId    SceneUniqId
}

func BuildScene(sceneId SceneUniqId, game *Game) *SceneBuilder {
	return &SceneBuilder{
		sceneId:    sceneId,
		parentGame: game,
	}
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

func (scene *Scene) Start() {}
func (scene *Scene) End()   {}
