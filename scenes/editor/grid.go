package editor

import (
	"GameFrameworkTM/components/level"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Functions for the actual editing grid/level

func (scene *Scene) drawLevelGrid(ctx engine.Context) {
	for i := range int32(ctx.Resolution.X/level.TILE_SIZE) + 1 {
		x := i * level.TILE_SIZE
		ctx.Tileset.DrawTile(0, x, 0)
		rl.DrawLine(
			x, 0,
			x, int32(ctx.Resolution.Y),
			rl.GetColor(uint(scene.ColorPallete[2])),
		)
	}
	for i := range int32(ctx.Resolution.Y/level.TILE_SIZE) + 1 {
		y := i * level.TILE_SIZE
		rl.DrawLine(
			0, y,
			int32(ctx.Resolution.X), y,
			rl.GetColor(uint(scene.ColorPallete[2])),
		)
	}
}
