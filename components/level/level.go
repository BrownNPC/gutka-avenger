package level

type Tile struct {
}

const TILE_SIZE = 16 //16px

type Tiles [256 / TILE_SIZE][160 / TILE_SIZE]Tile
type Level struct {
	Layers []Tiles
}
