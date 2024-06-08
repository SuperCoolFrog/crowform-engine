package pkg

import (
	"crowform/internal/cache"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sprite *Sprite) draw() {
	destRect := sprite.GetWindowDestRect()

	rl.DrawTexturePro(*sprite.getTexture(), sprite.srcRect, destRect, sprite.Origin, sprite.rotation, sprite.colorTint)
}

func (sprite *Sprite) setParent(parent *Actor) {
	sprite.parent = parent
}

func (sprite *Sprite) getTexture() *rl.Texture2D {
	if sprite.texture == nil {
		texture := cache.GetTexture2d(sprite.textureFileName)
		sprite.texture = &texture
		rl.SetTextureFilter(*sprite.texture, rl.FilterBilinear)
	}

	return sprite.texture
}
