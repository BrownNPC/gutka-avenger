package level

import (
	c "GameFrameworkTM/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const LEVEL_WIDTH = c.BaseResolutionX/ TILE_SIZE
const LEVEL_HEIGHT = c.BaseResolutionY / TILE_SIZE

type Tiles [LEVEL_WIDTH][LEVEL_HEIGHT]Tile
type Level struct {
	// Background image.
	Background rl.Texture2D
	Layers     []Tiles
}
