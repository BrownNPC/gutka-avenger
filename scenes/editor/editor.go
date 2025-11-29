package editor

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Screen       render.Screen
	LevelScreen  render.Screen
	ColorPallete [4]uint
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
		rl.ClearBackground(rl.GetColor(uint(scene.ColorPallete[3])))
		scene.drawLevelGrid(ctx)
	}
	rl.EndTextureMode()

	// Draw Editor
	scene.Screen.BeginDrawing()
	{
		rl.ClearBackground(rl.GetColor(scene.ColorPallete[0]))
		// Draw level
		scene.LevelScreen.RenderEx(0, 0, 400, 250)
		scene.drawTilePicker(ctx)
	}
	scene.Screen.EndDrawing()

	return false // if true is returned, Unload is called
}

func (scene *Scene) drawTilePicker(ctx engine.Context) {
	// Used to check if mouse cursor is inside
	canvas := rl.NewRectangle(405, 5, 190, 365)
	// check if mouse is inside tile picker
	IsMouseInside := rl.CheckCollisionCircleRec(scene.Screen.VirtualMouse(), 1, canvas)
	canvasColor := c.If(IsMouseInside, scene.ColorPallete[3], scene.ColorPallete[1])

	rl.DrawRectangleRoundedLines(canvas,
		0.1, 10, rl.GetColor(canvasColor))
	ctx.Tileset.DrawTileEx(0, 410, 10, 32, 32)
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "start" // the engine will switch to the scene that is registered with this id
}
