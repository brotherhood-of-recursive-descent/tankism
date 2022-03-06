package tankism

import "github.com/hajimehoshi/ebiten/v2"

type Client struct {
	/* scenes scenes.Manager
	audioMixer audio.Mixer
	loader assets.Loader */
}

func (t *Client) Update() error {
	return nil
}

func (t *Client) Draw(screen *ebiten.Image) {
}

func (t *Client) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func NewClient() *Client {
	return &Client{}
}
