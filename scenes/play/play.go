package play

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/components/render"
	"GameFrameworkTM/engine"

	"github.com/BrownNPC/simple-ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	// unload functions for convenience
	Defer  c.Stack[func()]
	p      *ecs.Pool
	Screen render.Screen
}

// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
	// Make a render texture for pixel art resolution.
	// And queue unload.
	scene.Screen = render.NewScreen(ctx.Resolution)
	scene.Defer.Add(scene.Screen.Unload)
	scene.p = ecs.New(10_000)
}

// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	scene.Screen.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("There is no game yet lol", 0, int32(ctx.Resolution.Y/2), 10, rl.Black)
	scene.Screen.EndDrawing()
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
