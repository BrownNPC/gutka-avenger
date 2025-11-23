package play

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"
	"fmt"

	"github.com/BrownNPC/simple-ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	// unload functions for convenience
	Defer  c.Stack[func()]
	p      *ecs.Pool
	Screen render.Screen
	// Surface is used for fonts
	Surface render.Screen
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	// Make a render texture for pixel art resolution.
	// And queue unload.
	scene.Screen = render.NewScreen(ctx.Resolution)
	scene.Surface = render.NewScreen(c.V2(600, 375))
	scene.Defer.Add(scene.Screen.Unload)
	scene.p = ecs.New(10_000)
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	// Actual game elements.
	scene.Screen.BeginDrawing()
	{
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText(fmt.Sprintf("%d Loaded tiles", ctx.Tileset.TileCount), 0, 0, 1, rl.Green)
		rl.DrawPixel(100, 20, rl.Blue)
		ctx.Tileset.DrawTile(0, 0, 0)
	}
	scene.Screen.EndDrawing()

	// Surface is used to draw fonts and GUI at a higher resolution.
	scene.Surface.BeginDrawing()
	{
		rl.ClearBackground(rl.Blank)
		rl.DrawTextEx(ctx.Font, "There is no game yet lol", c.V2(ctx.Resolution.X/2, 10).R(), 8, 0, rl.Black)
	}
	scene.Surface.EndDrawing()

	if rl.IsKeyPressed(rl.KeyA) {
		return true
	}
	return false // if true is returned, Unload is called
}

// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	// Unload everything queued.
	for _, f := range scene.Defer.Items {
		f()
	}
	return "start" // the engine will switch to the scene that is registered with this id
}
