package pkg

import (
	"crowform/internal/cache"
	"crowform/internal/tools"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (sprite *Sprite) update(deltaTime time.Duration) {
	tools.ForEach(sprite.queueForUpdate, func(f func()) {
		f()
	})
	sprite.queueForUpdate = nil

	sprite.updateAnimations(deltaTime)
}

func (sprite *Sprite) draw() {
	// destRect := sprite.GetWindowDestRect()
	destRect := sprite.getDrawDestRect()

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

func (sprite *Sprite) getDrawDestRect() rl.Rectangle {
	return sprite.getAnimationDestRect()
}

func (sprite *Sprite) addToUpdateQueue(item func()) {
	sprite.queueForUpdate = append(sprite.queueForUpdate, item)
}
