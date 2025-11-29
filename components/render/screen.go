package render

import (
	c "GameFrameworkTM/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Screen provides a wrapper around a raylib render texture
type Screen struct {
	renderTexture rl.RenderTexture2D
	Width, Height int
}

// Does letterboxing (black bars) to maintain the resolution.
func NewScreen(resolution c.Vec2) Screen {
	virtualWidth, virtualHeight := resolution.ToInt()
	r := Screen{
		renderTexture: rl.LoadRenderTexture(int32(virtualWidth), int32(virtualHeight)),
		Width:         virtualWidth,
		Height:        virtualHeight,
	}
	// Use nearest neighbor / point sampling for crisp integer scaling
	rl.SetTextureFilter(r.renderTexture.Texture, rl.FilterPoint)
	return r
}

func (r *Screen) Unload() {
	rl.UnloadRenderTexture(r.renderTexture)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) Render() {
	target := r.renderTexture
	scale := r.Scale()

	// compute integer destination size and integer centered offsets
	destW := r.Width * scale
	destH := r.Height * scale
	offsetX := (rl.GetRenderWidth() - destW) / 2
	offsetY := (rl.GetRenderHeight() - destH) / 2

	rl.DrawTexturePro(
		target.Texture,
		rl.Rectangle{X: 0, Y: 0, Width: float32(target.Texture.Width), Height: float32(-target.Texture.Height)},
		rl.Rectangle{
			X:      float32(offsetX),
			Y:      float32(offsetY),
			Width:  float32(destW),
			Height: float32(destH),
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
}

func (r *Screen) RenderEx(x, y, width, height float32) {
	target := r.renderTexture
	rl.DrawTexturePro(
		target.Texture,
		rl.Rectangle{X: 0, Y: 0, Width: float32(target.Texture.Width), Height: float32(-target.Texture.Height)},
		rl.Rectangle{
			X: x, Y: y,
			Width:  width,
			Height: height,
		},
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
}

func (r *Screen) BeginDrawing() {
	rl.BeginTextureMode(r.renderTexture)
}

func (r *Screen) EndDrawing() {
	rl.EndTextureMode()
	r.Render()
}

// Get scaling factor (integer). Clamp to at least 1 to avoid zero.
func (r *Screen) Scale() float32 {
	// use ints so / is integer division (floor)
	scaleW := float32(rl.GetRenderWidth()) / float32(r.Width)
	scaleH := rl.GetRenderHeight() / r.Height
	var scale int
	if scaleW > 1 {
		scale = scaleW
	} else {
		scale = scaleH
	}

	return max(1, scale)
}

func (r *Screen) VirtualMouse() rl.Vector2 {
	x, y := r.VirtualMouseInt()
	return c.V2(x, y).R()
}

// Get virtual mouse coordinates scaled within the screen.
func (r *Screen) VirtualMouseInt() (int, int) {
	mx, my := rl.GetMouseX(), rl.GetMouseY()
	scale := r.Scale()

	// integer offsets to match Render()
	destW := r.Width * scale
	destH := r.Height * scale
	offsetX := (rl.GetRenderWidth() - destW) / 2
	offsetY := (rl.GetRenderHeight() - destH) / 2

	vx := (float32(mx) - float32(offsetX)) / float32(scale)
	vy := (float32(my) - float32(offsetY)) / float32(scale)

	// clamp to virtual resolution
	if vx < 0 {
		vx = 0
	}
	if vy < 0 {
		vy = 0
	}
	if vx > float32(r.Width) {
		vx = float32(r.Width)
	}
	if vy > float32(r.Height) {
		vy = float32(r.Height)
	}

	return int(vx), int(vy)
}
