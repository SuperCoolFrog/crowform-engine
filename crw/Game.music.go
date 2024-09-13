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
	rl.SetMusicVolume(game.currentMusic, 0.10)
	rl.PlayMusicStream(game.currentMusic)
	// rl.ResumeMusicStream(game.currentMusic)
}

func (game *Game) UnloadMusic() {
	if game.hasLoadedMusic {
		rl.UnloadMusicStream(game.currentMusic)
		game.hasLoadedMusic = false
	}
}
