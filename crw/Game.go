package crw

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
	close         bool

	hasInitAudio   bool
	hasLoadedMusic bool
	currentMusic   rl.Music

	soundQ []rl.Sound
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

func resetGlobals() {
	lastMouseDownActor = nil
	mouseOverTarget = nil
	cache.ResetSettings()
	cache.RestPreload()
	cache.UnloadFontsCache()
	cache.UnloadTextureCache()
}

func (builder *GameBuilder) Build() *Game {
	resetGlobals()

	cache.SetSetting(cache.SettingName_AssetDirectory, builder.assetDirectory)

	return &Game{
		windowName:   builder.windowName,
		windowWidth:  builder.windowWidth,
		windowHeight: builder.windowHeight,
		scenes:       make(map[SceneUniqId]*Scene),
		paused:       false,
		close:        false,
		soundQ:       make([]rl.Sound, 0),
	}
}

// Game Functions: Public

func (game *Game) Start() {
	mogError := mog.Init(true)
	if mogError != nil {
		panic(mogError)
	}
	defer mog.CleanUp()

	rl.InitWindow(game.windowWidth, game.windowHeight, game.windowName)

	cache.RunPreload()

	rl.SetExitKey(0) // Unsets escape to close

	defer rl.CloseWindow()

	// rl.SetTargetFPS(120) // causes jaggy animations
	game.lastFrameTime = time.Now()

	for !(rl.WindowShouldClose() || game.close) {
		game.checkInputEvents()

		if game.hasLoadedMusic {
			rl.UpdateMusicStream(game.currentMusic)
		}

		game.playAllSounds()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		game.updateScene()

		rl.EndDrawing()
	}

	cache.UnloadTextureCache()
	cache.UnloadFontsCache()
	game.UnloadMusic()
	unloadSoundsCache()
	rl.CloseAudioDevice()
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

func (game *Game) Shutdown() {
	game.close = true
}

func (game *Game) addSoundToQ(sound rl.Sound) {
	game.soundQ = append(game.soundQ, sound)
}

func (game *Game) playAllSounds() {
	for i := range game.soundQ {
		if i < len(game.soundQ) {
			rl.PlaySound(game.soundQ[i])
		} else {
			break
		}
	}

	game.soundQ = nil
}
