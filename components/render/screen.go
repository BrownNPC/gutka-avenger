package render

import (
	c "GameFrameworkTM/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Screen provides a wrapper around a raylib render texture
type Screen struct {
	renderTexture      rl.RenderTexture2D
	InternalResolution c.Vec2
	TargetPosition     c.Vec2
	TargetResolution   c.Vec2
}

// Does letterboxing (black bars) to maintain the resolution.
func NewScreen(resolution c.Vec2) Screen {
	virtualWidth, virtualHeight := resolution.ToInt()
	r := Screen{
		renderTexture:      rl.LoadRenderTexture(int32(virtualWidth), int32(virtualHeight)),
		InternalResolution: resolution,
	}
	rl.SetTextureFilter(r.renderTexture.Texture, rl.FilterAnisotropic16x)
	return r

}
func (r *Screen) Unload() {
	rl.UnloadRenderTexture(r.renderTexture)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) Render() {
	target := r.renderTexture
	// RENDER SIZE
	RS := c.V2(rl.GetRenderWidth(), rl.GetRenderHeight())
	r.TargetResolution = RS //set temporarily for calculating scale
	//Fractionally scale InternalResolution
	r.TargetResolution = r.InternalResolution.Scale(r.Scale())

	// Area remaining after fractionally scaling InternalResolution
	remainingW := RS.X - r.TargetResolution.X
	remainingH := RS.Y - r.TargetResolution.Y

	// Draw final picture in center
	r.TargetPosition.X = remainingW * 0.5
	r.TargetPosition.Y = remainingH * 0.5

	rl.DrawTexturePro(
		target.Texture,
		rl.Rectangle{Width: float32(target.Texture.Width), Height: float32(-target.Texture.Height)},
		rl.NewRectangle(
			r.TargetPosition.X, r.TargetPosition.Y,
			r.TargetResolution.X, r.TargetResolution.Y,
		),
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
}
func (r *Screen) RenderEx(x, y, width, height float32) {
	target := r.renderTexture
	r.TargetPosition = c.V2(x, y)
	r.TargetResolution = c.V2(x, y)
	rl.DrawTexturePro(
		target.Texture,
		rl.Rectangle{Width: float32(target.Texture.Width), Height: float32(-target.Texture.Height)},
		rl.NewRectangle(
			r.TargetPosition.X, r.TargetPosition.Y,
			r.TargetResolution.X, r.TargetResolution.Y,
		),
		rl.Vector2{X: 0, Y: 0}, 0, rl.White,
	)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) BeginDrawing() {
	rl.BeginTextureMode(r.renderTexture)
}

// render the render texture with black bars to keep the aspect ratio
func (r *Screen) EndDrawing() {
	rl.EndTextureMode()
	r.Render()
}

// Get scaling factor.
func (r *Screen) Scale() float32 {
	return min(
		r.TargetResolution.X/r.InternalResolution.X,
		r.TargetResolution.Y/r.InternalResolution.Y,
	)
}
func (r *Screen) VirtualMouse() rl.Vector2 {
	x, y := r.VirtualMouseInt()
	return c.V2(x, y).R()
}

// Get virtual mouse coordinates scaled within the screen.
func (r *Screen) VirtualMouseInt() (int, int) {
	mx, my := rl.GetMouseX(), rl.GetMouseY()
	scale := r.Scale()

	offsetX := (float32(rl.GetRenderWidth()) - float32(r.Width)*scale) * 0.5
	offsetY := (float32(rl.GetRenderHeight()) - float32(r.Height)*scale) * 0.5

	vx := (float32(mx) - offsetX) / scale
	vy := (float32(my) - offsetY) / scale

	if vx < 0 {
		vx = 0
	}
	vx = max(vx, 0)
	vy = max(vy, 0)

	vx = min(vx, float32(r.Width))
	vy = min(vy, float32(r.Height))

	return int(vx), int(vy)
}
