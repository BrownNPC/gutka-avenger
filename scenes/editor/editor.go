package editor

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/level"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Screen       render.Screen
	LevelScreen  render.Screen
	ColorPallete [4]uint
	TilePicker   [92][5][9]int
	SelectedTile level.Tile
	Defer        c.Stack[func()]
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	// Main editor screen with UI
	scene.Screen = render.NewScreen(ctx.SecondaryResolution)
	scene.Defer.Add(scene.Screen.Unload)
	// The actual grid/level
	scene.LevelScreen = render.NewScreen(ctx.Resolution)
	scene.Defer.Add(scene.LevelScreen.Unload)
	// ColorPallete used throughout the editor
	scene.ColorPallete = [4]uint{
		0x37353EFF, 0x44444EFF, 0x715A5AFF, 0xD3DAD9FF,
	}
}

// Update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	// Render Level to Render Texture, dont draw to screen yet.
	scene.LevelScreen.BeginDrawing()
	{
		rl.ClearBackground(rl.GetColor(uint(scene.ColorPallete[2])))
		scene.drawLevelGrid(ctx)
	}
	rl.EndTextureMode()

	// Draw Editor
	scene.Screen.BeginDrawing()
	{
		rl.ClearBackground(rl.GetColor(scene.ColorPallete[0]))
		// Draw level
		scene.LevelScreen.RenderEx(0, 0, 400, 250, &scene.Screen)
		scene.drawTilePicker(ctx, rl.NewRectangle(405, 5, 190, 365))
		levelGridHeight := scene.LevelScreen.TargetResolution.Y
		ctx.Tileset.DrawTileEx(scene.SelectedTile.Index, 10, int32(levelGridHeight+30), 32, 32)
	}
	scene.Screen.EndDrawing()

	return false // if true is returned, Unload is called
}

// TILE PICKER
func (scene *Scene) drawTilePicker(ctx engine.Context, canvas rl.Rectangle) {
	// Used to check if mouse cursor is inside
	// check if mouse is inside tile picker
	IsMouseInside := rl.CheckCollisionPointRec(scene.Screen.VirtualMouse(), canvas)
	canvasColor := c.If(IsMouseInside, scene.ColorPallete[3], scene.ColorPallete[1])

	// Draw tile pallete (border)
	rl.DrawRectangleRoundedLines(canvas,
		0.1, 5, rl.GetColor(canvasColor))
	scene.tilePickerTiles(ctx)
}

// Render all the loaded tiles in the tileset (ctx.Tileset)
func (scene *Scene) tilePickerTiles(ctx engine.Context) {
	var TileIndex int
	for j := range int32(9) {
		y := j*36 + 16
		for i := range int32(5) {
			x := 412 + (i * 36)
			ctx.Tileset.DrawTileEx(TileIndex, x, y, 32, 32)

			// Draw outline if mouse hovers over tile
			mouseInsideThisTile := rl.CheckCollisionPointRec(scene.Screen.VirtualMouse(),
				rl.NewRectangle(float32(x), float32(y), 32, 32))
			if mouseInsideThisTile {
				rl.DrawRectangleLines(x, y, 32, 32, rl.GetColor(scene.ColorPallete[3]))
				// Make this the SelectedTile if mouse is clicked
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					scene.SelectedTile = level.Tile{Index: TileIndex}
				}
			}
			// Only draw tiles that are available
			TileIndex++
			if TileIndex == ctx.Tileset.TileCount {
				return
			}
		}
	}
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "start" // the engine will switch to the scene that is registered with this id
}
