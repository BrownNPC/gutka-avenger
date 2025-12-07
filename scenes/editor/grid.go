package editor

import (
	"GameFrameworkTM/components/level"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Functions for the actual editing grid/level
func (scene *Scene) drawLevelGrid(ctx engine.Context) {
	for j := range int32(level.TOTAL_TILES_Y) {
		y := j * level.TILE_SIZE
		for i := range int32(level.TOTAL_TILES_X) {
			x := i * level.TILE_SIZE
			ctx.Tileset.DrawTile(0, x, y)
			tileRec := rl.NewRectangle(float32(x), float32(y), level.TILE_SIZE, level.TILE_SIZE)
			// Check if mouse is hovering on the tile.
			m := scene.LevelScreen.VirtualMouse()
			isMouseInsideTile := rl.CheckCollisionPointRec(
				m, tileRec)

			if isMouseInsideTile {
				rl.DrawRectangleLinesEx(tileRec, 1.0, rl.GetColor(scene.ColorPallete[1]))
			}
		}
	}
}
