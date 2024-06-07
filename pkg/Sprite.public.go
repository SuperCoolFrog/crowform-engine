package pkg

import (
	"crowform/internal/cache"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sprite *Sprite) Draw() {
	if sprite.texture == nil {
		texture := cache.GetTexture2d(sprite.textureFileName)
		sprite.texture = &texture
	}
	rl.DrawTexturePro(*sprite.texture, sprite.srcRect, sprite.DestRect, sprite.Origin, sprite.rotation, sprite.colorTint)
}
