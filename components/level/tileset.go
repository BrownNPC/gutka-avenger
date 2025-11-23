package level

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILES_PER_ROW = 1024 / TILE_SIZE
const ATLAS_MAX_TILES = TILES_PER_ROW * TILES_PER_ROW // 1024

type Tileset rl.Texture2D

func (a Tileset) GetTile(idx int) rl.RectangleInt32 {
	if idx < 0 || idx >= ATLAS_MAX_TILES {
		panic("GetTile: Tile index out of range")
	}

	x := (idx % TILES_PER_ROW) * TILE_SIZE
	y := (idx / TILES_PER_ROW) * TILE_SIZE

	return rl.RectangleInt32{
		X:      int32(x),
		Y:      int32(y),
		Width:  TILE_SIZE,
		Height: TILE_SIZE,
	}
}

func (a Tileset) DrawTile(idx int, posX, posY int32) {
	src := a.GetTile(idx)

	rl.DrawTexturePro(
		rl.Texture2D(a),
		rl.Rectangle{
			X:      float32(src.X),
			Y:      float32(src.Y),
			Width:  float32(src.Width),
			Height: float32(src.Height),
		},
		rl.Rectangle{
			X:      float32(posX),
			Y:      float32(posY),
			Width:  float32(src.Width),
			Height: float32(src.Height),
		},
		rl.Vector2{}, 0, rl.White,
	)
}
