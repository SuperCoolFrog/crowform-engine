package pkg

import (
	"crowform/internal/cache"
	"crowform/internal/mog"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameBuilder struct {
	windowName   string
	windowWidth  int32
	windowHeight int32

	assetDirectory string
}

type Game struct {
	windowName   string
	windowWidth  int32
	windowHeight int32

	scenes        map[SceneUniqId]*Scene
	currentScene  *Scene
	lastFrameTime time.Time
	paused        bool
}

func BuildGame() *GameBuilder {
	return &GameBuilder{
		assetDirectory: "assets",
	}
}

func (builder *GameBuilder) WithWindowName(name string) *GameBuilder {
	builder.windowName = name
	return builder
}
func (builder *GameBuilder) WithDimensions(width int32, height int32) *GameBuilder {
	builder.windowWidth = width
	builder.windowHeight = height
	return builder
}
func (builder *GameBuilder) WithAssetDirectory(directoryName string) *GameBuilder {
	builder.assetDirectory = directoryName
	return builder
}

func (builder *GameBuilder) Build() *Game {
	cache.SetSetting(cache.SettingName_AssetDirectory, builder.assetDirectory)

	return &Game{
		windowName:   builder.windowName,
		windowWidth:  builder.windowWidth,
		windowHeight: builder.windowHeight,
		scenes:       make(map[SceneUniqId]*Scene),
		paused:       false,
	}
}

// Game Functions: Public

func (game *Game) Start() {
	mogError := mog.Init(true)
	if mogError != nil {
		panic(mogError)
	}

	rl.InitWindow(game.windowWidth, game.windowHeight, game.windowName)

	rl.SetExitKey(0) // Unsets escape to close

	defer rl.CloseWindow()

	// rl.SetTargetFPS(120) // causes jaggy animations
	game.lastFrameTime = time.Now()

	for !rl.WindowShouldClose() {
		game.checkInputEvents()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		game.updateScene()

		rl.EndDrawing()
	}
}

func (game *Game) AddScene(scene *Scene) {
	game.scenes[scene.SceneId] = scene
}

// Game Functions: Private

func (game *Game) updateScene() {
	now := time.Now()
	delta := now.Sub(game.lastFrameTime)

	if game.currentScene != nil {
		game.currentScene.Update(delta)
		game.currentScene.Draw()
	}

	game.lastFrameTime = now
}

func (game *Game) RemoveScene(scene *Scene) {
	delete(game.scenes, scene.SceneId)
}

func (game *Game) GoToScene(sceneId SceneUniqId) {
	game.endScene()

	nextScene, found := game.scenes[sceneId]

	if !found {
		return
	}

	game.currentScene = nextScene
	game.currentScene.Start()
}

func (game *Game) endScene() {
	if game.currentScene == nil {
		return
	}
	game.currentScene.End()
	game.currentScene = nil
}
