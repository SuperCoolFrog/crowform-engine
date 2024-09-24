package cache

import (
	"crowform/internal/tools"
	"fmt"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

/** Textures **/
var textures map[string]rl.Texture2D = make(map[string]rl.Texture2D)

func GetTexture2d(filename string) rl.Texture2D {
	if texture, ok := textures[filename]; ok {
		return texture
	}

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

func GetTexture2dForEdit(filename string) rl.Texture2D {
	key := "EDIT:" + filename
	if texture, ok := textures[key]; ok {
		return texture
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := fmt.Sprintf("%s%s%s%s%s", exPath, string(os.PathSeparator), GetSetting[string](SettingName_AssetDirectory), string(os.PathSeparator), filename)

	nuTexture := rl.LoadTexture(fullPath)

	textures[key] = nuTexture

	return nuTexture
}

func UnloadTexture2d(filename string) {
	rl.UnloadTexture(textures[filename])
	delete(textures, filename)
}

func ReloadTexture2d(filename string) {
	UnloadTexture2d(filename)
	GetTexture2d(filename)
}

func UnloadTextureCache() {
	for _, texture := range textures {
		rl.UnloadTexture(texture)
	}

	for k := range textures {
		delete(textures, k)
	}
}

/** Fonts **/
var fonts map[string]rl.Font = make(map[string]rl.Font)

func GetFont(fontName string, fontSize float32, isCustom bool) rl.Font {
	key := fmt.Sprintf("%s::%f", fontName, fontSize)
	if font, ok := fonts[key]; ok {
		return font
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fullPath := fmt.Sprintf("%s%s%s%s%s", exPath, string(os.PathSeparator), GetSetting[string](SettingName_AssetDirectory), string(os.PathSeparator), fontName)

	var nuFont rl.Font
	if isCustom {
		nuFont = rl.LoadFont(fullPath) // Loades 32 by default
	} else {
		nuFont = rl.LoadFontEx(fullPath, int32(tools.TruncFloat32(fontSize)), nil, 0) // Could not get to load
	}

	fonts[key] = nuFont

	return nuFont
}

func UnloadFontsCache() {
	for _, font := range fonts {
		rl.UnloadFont(font)
	}

	for k := range fonts {
		delete(fonts, k)
	}
}
