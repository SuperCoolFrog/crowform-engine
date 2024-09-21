package crw

import (
	"crowform/internal/cache"
	"fmt"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSoundBuilder struct {
	fileName string
	game     *Game
}

type GameSound struct {
	sound rl.Sound
	game  *Game
}

func BuildSound(filename string, game *Game) *GameSoundBuilder {
	return &GameSoundBuilder{
		fileName: filename,
		game:     game,
	}
}

var soundsCache map[string]rl.Sound = make(map[string]rl.Sound)

func (builder *GameSoundBuilder) getSound() rl.Sound {
	if !builder.game.hasInitAudio {
		rl.InitAudioDevice()
		builder.game.hasInitAudio = true
	}

	key := builder.fileName
	if s, ok := soundsCache[key]; ok {
		return s
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := fmt.Sprintf("%s%s%s%s%s", exPath, string(os.PathSeparator), cache.GetSetting[string](cache.SettingName_AssetDirectory), string(os.PathSeparator), builder.fileName)

	nuSound := rl.LoadSound(fullPath)

	soundsCache[key] = nuSound

	return nuSound
}

func (builder *GameSoundBuilder) Build() *GameSound {
	gs := &GameSound{
		game: builder.game,
	}

	gs.sound = builder.getSound()
	rl.SetSoundVolume(gs.sound, 0.35)

	return gs
}

func (gameSound *GameSound) Play() {
	gameSound.game.addSoundToQ(gameSound.sound)
}

func unloadSoundsCache() {
	for _, s := range soundsCache {
		rl.UnloadSound(s)
	}

	for k := range soundsCache {
		delete(soundsCache, k)
	}
}
