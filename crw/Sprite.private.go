package crw

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

	rl.DrawTexturePro(*sprite.getCachedTexture(), sprite.srcRect, destRect, sprite.Origin, sprite.rotation, sprite.colorTint)
}

func (sprite *Sprite) setParent(parent *Actor) {
	sprite.parent = parent
}

func (sprite *Sprite) getCachedTexture() *rl.Texture2D {
	if sprite.texture == nil {
		texture := cache.GetTexture2d(sprite.textureFileName)
		sprite.texture = &texture
		rl.SetTextureFilter(*sprite.texture, rl.FilterBilinear)
	}

	return sprite.texture
}

func (sprite *Sprite) getTexture() rl.Texture2D {
	return cache.GetTexture2d(sprite.textureFileName)
}

func (sprite *Sprite) reloadTexture() {
	sprite.texture = nil
	sprite.getCachedTexture()
}

func (sprite *Sprite) getDrawDestRect() rl.Rectangle {
	return sprite.getAnimationDestRect()
}

func (sprite *Sprite) addToUpdateQueue(item func()) {
	sprite.queueForUpdate = append(sprite.queueForUpdate, item)
}

func (sprite *Sprite) setTextureOpacity(inOpacity float64) {
	opacity := inOpacity
	// Ensure opacity is between 0.0f and 1.0f
	if inOpacity < 0.0 {
		opacity = 0.0
	}
	if inOpacity > 1.0 {
		opacity = 1.0
	}

	spriteTexture := cache.GetTexture2dForEdit(sprite.textureFileName)
	// Get image data from the texture
	image := rl.LoadImageFromTexture(spriteTexture)

	// Get the pixel data from the image
	var pixels = rl.LoadImageColors(image)

	// Modify the alpha channel for each pixel
	end := int(image.Width * image.Height)
	for i := 0; i < end; i++ {
		pixels[i].A = uint8(255 * opacity)
	}

	// Update the texture with the new image data
	rl.UpdateTexture(spriteTexture, pixels)
	sprite.texture = &spriteTexture

	// Clean up
	rl.UnloadImageColors(pixels)
	rl.UnloadImage(image)
}
