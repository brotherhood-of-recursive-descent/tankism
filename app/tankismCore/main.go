package main

import "github.com/hajimehoshi/ebiten/v2"

type TankismCore struct {
}

// Draw renders the game screen.
func (t *TankismCore) Draw(screen *ebiten.Image) {
	// Add rendering logic here
}

// Update updates the game state.
func (t *TankismCore) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	return nil
}

// Layout specifies the screen dimensions.
func (t *TankismCore) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Set the screen dimensions
	return outsideWidth, outsideHeight
}

func main() {
	// Initialize the TankismCore application
	app := &TankismCore{}

	// Start the application
	ebiten.SetFullscreen(true)
	ebiten.RunGame(app)
}
