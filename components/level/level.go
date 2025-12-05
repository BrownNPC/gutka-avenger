package level

import (
	c "GameFrameworkTM/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const TOTAL_TILES_X = c.BaseResolutionX/ TILE_SIZE
const TOTAL_TILES_Y = c.BaseResolutionY / TILE_SIZE

type Tiles [TOTAL_TILES_X][TOTAL_TILES_Y]Tile
type Level struct {
	// Background image.
	Background rl.Texture2D
	Layers     []Tiles
}
