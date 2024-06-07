package pkg

import (
	"crowform/internal/cache"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sprite *Sprite) draw() {
	if sprite.texture == nil {
		texture := cache.GetTexture2d(sprite.textureFileName)
		sprite.texture = &texture
	}

	destRect := sprite.GetWindowDestRect()

	rl.DrawTexturePro(*sprite.texture, sprite.srcRect, destRect, sprite.Origin, sprite.rotation, sprite.colorTint)
}

func (sprite *Sprite) setParent(parent *Actor) {
	sprite.parent = parent
}
