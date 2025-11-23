package level

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	// Which tile is this from the atlas?
	Index int
}

const TILE_SIZE = 16 //16px

type Tiles [256 / TILE_SIZE][160 / TILE_SIZE]Tile
type Level struct {
	// Background image.
	Background rl.Texture2D
	Layers     []Tiles
}
