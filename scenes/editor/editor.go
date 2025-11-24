package editor

import (
	"GameFrameworkTM/components/level"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Screen          render.Screen
	LevelScreen     render.Screen
	LevelFullScreen bool
	ColorPallete    [4]int
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	scene.Screen = render.NewScreen(ctx.SecondaryResolution)
	scene.LevelScreen = render.NewScreen(ctx.Resolution)
	scene.ColorPallete = [4]int{
		0x37353EFF, 0x44444EFF, 0x715A5AFF, 0xD3DAD9FF,
	}
}

// Update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	{
		scene.LevelScreen.BeginDrawing()
		scene.drawLevel(ctx)
		rl.EndTextureMode()
	}

	if scene.LevelFullScreen {
		scene.LevelScreen.Render()
	} else {
		scene.drawEditor(ctx)
	}

	if rl.IsKeyPressed(rl.KeyTab) {
		scene.LevelFullScreen = !scene.LevelFullScreen
	}

	return false // if true is returned, Unload is called
}

func (scene *Scene) drawEditor(ctx engine.Context) {
	scene.Screen.BeginDrawing()
	{
		rl.ClearBackground(rl.GetColor(0x37353EFF))
		scene.LevelScreen.RenderEx(0, 0, ctx.Resolution.X*2, ctx.Resolution.Y*2)
	}
	scene.Screen.EndDrawing()
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "start" // the engine will switch to the scene that is registered with this id
}

func (scene *Scene) drawLevel(ctx engine.Context) {
	rl.ClearBackground(rl.White)
	for i := range int32(ctx.Resolution.X / level.TILE_SIZE) {
		x := i * level.TILE_SIZE
		ctx.Tileset.DrawTile(0, x, 0)
		rl.DrawLine(
			x, 0,
			x, int32(ctx.Resolution.Y),
			rl.Blue,
		)
	}
	for i := range int32(ctx.Resolution.Y / level.TILE_SIZE) {
		y := i * level.TILE_SIZE
		rl.DrawLine(
			0, y,
			int32(ctx.Resolution.X), y,
			rl.Blue,
		)
	}
}
