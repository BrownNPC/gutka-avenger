package start

import (
	"GameFrameworkTM/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
}
// Load is called once the scene is switched to
func (scene *Scene) Load(ctx engine.Context) {
}
// update is called every frame
func (scene *Scene) Update(ctx engine.Context) (unload bool) {
	if rl.IsKeyPressed(rl.KeyEnter){
		return true		
	}
	return false// if true is returned, Unload is called
}
// called after Update returns true
func (scene *Scene) Unload(ctx engine.Context) (nextSceneID string) {
	return "play"  // the engine will switch to the scene that is registered with this id
}

