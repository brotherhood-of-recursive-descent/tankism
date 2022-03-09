package tankism

import (
	_ "embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

//go:embed assets/start_background.jpg
var backgroundImage []byte

type Client struct {
	backgroundImage *ebiten.Image
	windowWidth     int
	windowHeight    int
	/* scenes scenes.Manager
	audioMixer audio.Mixer
	loader assets.Loader */
}

func (t *Client) Update() error {
	return nil
}

func (t *Client) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	scaleX, scaleY := 0.75, 0.75
	imageWidth, imageHeight := t.backgroundImage.Size()

	x := float64(t.windowWidth)*0.5 - (float64(imageWidth) * 0.5 * scaleX)
	y := float64(t.windowHeight)*0.5 - (float64(imageHeight) * 0.5 * scaleY)

	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(x, y)

	screen.Fill(color.Black)
	screen.DrawImage(t.backgroundImage, op)

	text.Draw(screen, "Loading...", FontBig, t.windowWidth-300, t.windowHeight-100, color.RGBA{0xfa, 0xfa, 0xd2, 0xff})
}

func (t *Client) Layout(outsideWidth, outsideHeight int) (int, int) {
	t.windowWidth = outsideWidth
	t.windowHeight = outsideHeight
	return outsideWidth, outsideHeight
}

func NewClient() *Client {

	img, _ := LoadImage(backgroundImage)
	sprite := ebiten.NewImageFromImage(img)

	return &Client{
		backgroundImage: sprite,
	}
}
