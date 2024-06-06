package pkg

import (
	"crowform/internal/cache"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sprite *Sprite) Draw() {
	texture := cache.GetTexture2d(sprite.textureFileName)
	rl.DrawTexturePro(texture, sprite.srcRect, sprite.DestRect, sprite.Origin, sprite.rotation, sprite.colorTint)
}
