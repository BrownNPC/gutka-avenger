package main

import (
	c "GameFrameworkTM/components"
	"GameFrameworkTM/engine"
	"GameFrameworkTM/scenes"
	"fmt"
	"io/fs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ASSETS could either come from an embedded folder (on web).
// or from the current working directory's "./assets"
var ASSETS fs.FS

// You can register scenes in scenes/register.go

// You can edit the window title in this file.
func main() {
	// rl.SetTraceLogLevel(rl.LogError)
	err := engine.Run(scenes.Registered, engine.Config{
		WindowTitle:         "Gutka Avenger",
		Resolution:          c.V2(c.BaseResolutionX, c.BaseResolutionY),
		SecondaryResolution: c.V2(c.SecondaryResolutionX, c.SecondaryResolutionY),
		TilesetPath:         "assets/tileset.png",
		FontPath:            "assets/press-start-2p.ttf",
		ExitKey:             rl.KeyNumLock,
	}, ASSETS)
	if err != nil {
		fmt.Println(err)
	}
}
