package crw

import (
	"crowform/internal/cache"
	"crowform/internal/mog"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) PlayWav(filename string) {
	if game.isMuted {
		return
	}

	if !game.hasInitAudio {
		rl.InitAudioDevice()
		game.hasInitAudio = true
	}

	if game.hasLoadedMusic {
		game.UnloadMusic()
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := fmt.Sprintf("%s%s%s%s%s", exPath, string(os.PathSeparator), cache.GetSetting[string](cache.SettingName_AssetDirectory), string(os.PathSeparator), filename)

	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		mog.Error("%v", err)
		panic(err)
	}

	game.currentMusic = rl.LoadMusicStream(fullPath)
	game.hasLoadedMusic = true
	rl.SetMusicVolume(game.currentMusic, game.musicVolume)
	rl.PlayMusicStream(game.currentMusic)
	// rl.ResumeMusicStream(game.currentMusic)
}

func (game *Game) UnloadMusic() {
	if game.hasLoadedMusic {
		rl.UnloadMusicStream(game.currentMusic)
		game.hasLoadedMusic = false
	}
}

// Between 0 - 1
func (game *Game) SetMusicVolume(volume float32) {
	game.musicVolume = volume
	if game.hasLoadedMusic {
		rl.SetMusicVolume(game.currentMusic, game.musicVolume)
	}
}
func (game *Game) GetMusicVolume() float32 {
	return game.musicVolume
}

// Between 0 - 1
func (game *Game) SetSoundsVolume(volume float32) {
	game.soundVolume = volume
}
func (game *Game) GetSoundVolume() float32 {
	return game.soundVolume
}

func (game *Game) MuteAll() {
	game.isMuted = true
	game.UnloadMusic()
}
