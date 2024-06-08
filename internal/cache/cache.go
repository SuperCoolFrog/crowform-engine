package cache

import (
	"fmt"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/** Textures **/
var textures map[string]rl.Texture2D = make(map[string]rl.Texture2D)

func GetTexture2d(filename string) rl.Texture2D {
	// if texture, ok := textures[filename]; ok {
	// 	return texture
	// }

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := fmt.Sprintf("%s%s%s%s%s", exPath, string(os.PathSeparator), GetSetting[string](SettingName_AssetDirectory), string(os.PathSeparator), filename)

	nuTexture := rl.LoadTexture(fullPath)

	textures[filename] = nuTexture

	return nuTexture
}

func UnloadTextureCache() {
	for _, texture := range textures {
		rl.UnloadTexture(texture)
	}
}

/** Fonts **/
var fonts map[string]rl.Font = make(map[string]rl.Font)

func GetFont(fontName string) rl.Font {
	if font, ok := fonts[fontName]; ok {
		return font
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := fmt.Sprintf("%s%s%s%s%s", exPath, string(os.PathSeparator), GetSetting[string](SettingName_AssetDirectory), string(os.PathSeparator), fontName)

	nuFont := rl.LoadFont(fullPath) // Loades 32 by default
	// nuFont := rl.LoadFontEx(fullPath, 96, nil, 0) // Could not get to load

	fonts[fontName] = nuFont

	return nuFont
}

func UnloadFontsCache() {
	for _, font := range fonts {
		rl.UnloadFont(font)
	}
}
