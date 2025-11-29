package level

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILES_PER_ROW = 1024 / TILE_SIZE
const ATLAS_MAX_TILES = TILES_PER_ROW * TILES_PER_ROW

type Tileset struct {
	TileCount int
	Texture   rl.Texture2D
}

func LoadTileset(path string, assets fs.FS) (t Tileset, err error) {
	f, err := assets.Open(path)
	if err != nil {
		err = fmt.Errorf("LoadTileset: %w", err)
		return
	}
	img, err := png.Decode(f)
	if err != nil {
		err = fmt.Errorf("LoadTileset: %w", err)
		return
	}
	// count the number of tiles in the tileset.
	var tileCount int
	for i := range ATLAS_MAX_TILES {
		if tileIsOpaque(img, i) {
			tileCount++
		}
	}
	t = Tileset{
		Texture:   rl.LoadTexture(path),
		TileCount: tileCount,
	}
	return
}

func tileIsOpaque(img image.Image, i int) bool {
	tileX := (i % TILES_PER_ROW) * TILE_SIZE
	tileY := (i / TILES_PER_ROW) * TILE_SIZE

	for x := range TILE_SIZE {
		for y := range TILE_SIZE {
			px := tileX + x
			py := tileY + y

			_, _, _, alpha := img.At(px, py).RGBA()
			if alpha != 0 {
				return true
			}
		}
	}

	return false
}

// Get tile's pixel rectangle
func (a Tileset) GetTile(idx int) rl.RectangleInt32 {
	if idx < 0 || idx >= ATLAS_MAX_TILES {
		panic("GetTile: Tile index out of range")
	}
	if idx > a.TileCount {
		rl.TraceLog(rl.LogWarning, "Tileset.GetTile: idx is %d > TileCount", idx)
		idx = 0
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
		a.Texture,
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
func (a Tileset) DrawTileEx(idx int, posX, posY, width, height int32) {
	src := a.GetTile(idx)

	rl.DrawTexturePro(
		a.Texture,
		rl.Rectangle{
			X:      float32(src.X),
			Y:      float32(src.Y),
			Width:  float32(src.Width),
			Height: float32(src.Height),
		},
		rl.Rectangle{
			X:      float32(posX),
			Y:      float32(posY),
			Width:  float32(width),
			Height: float32(height),
		},
		rl.Vector2{}, 0, rl.White,
	)
}
