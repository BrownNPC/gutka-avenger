package editor

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/level"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Screen      render.Screen
	LevelScreen render.Screen
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	scene.Screen = render.NewScreen(ctx.SecondaryResolution)
	scene.LevelScreen = render.NewScreen(ctx.Resolution)
}

// Update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	scene.LevelScreen.BeginDrawing()
	{
		rl.ClearBackground(rl.White)
		ctx.Tileset.DrawTile(0, 0, 0)
		for x := range int32(ctx.Resolution.X / level.TILE_SIZE) {
			x *= level.TILE_SIZE
			rl.DrawLineEx(
				c.V2(x, 0).R(),
				c.V2(x, ctx.Resolution.Y).R(),
				1, rl.ColorAlpha(rl.Blue, 0.15),
			)
		}
	}
	scene.LevelScreen.EndDrawing()
	// rl.EndTextureMode()

	// scene.Screen.BeginDrawing()
	// {
	// 	rl.ClearBackground(rl.Blue)
	// }
	// scene.Screen.EndDrawing()
	return false // if true is returned, Unload is called
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "start" // the engine will switch to the scene that is registered with this id
}
