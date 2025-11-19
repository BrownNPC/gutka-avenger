package main

import (
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
	rl.SetTraceLogLevel(rl.LogError)
	err := engine.Run(scenes.Registered, engine.Config{
		WindowTitle: "change this in main.go",
	}, ASSETS)
	if err != nil {
		fmt.Println(err)
	}
}
