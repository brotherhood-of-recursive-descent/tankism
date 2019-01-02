package internal

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type WindowManager struct {
	width  int32
	height int32
	title  string
	window *sdl.Window
}

func NewSDLWindowManager(width int32, height int32, title string) *WindowManager {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	return &WindowManager{
		width:  width,
		height: height,
		title:  title,
	}
}

func (wm *WindowManager) CreateWindow() (*sdl.Window, error) {
	window, err := sdl.CreateWindow(wm.title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		wm.width, wm.height, sdl.WINDOW_OPENGL)

	if err != nil {
		log.Fatalf("creating window: %s", err)
	}
	wm.window = window

	return window, err
}

func (wm *WindowManager) ToogleFullscreen() {
	isFullscreenFlagSet := wm.window.GetFlags() & uint32(sdl.WINDOW_FULLSCREEN)
	isFullscreen := isFullscreenFlagSet > 0
	if isFullscreen {
		wm.window.SetFullscreen(0)
	} else {
		wm.window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	}
	sdl.ShowCursor(int(isFullscreenFlagSet))
}
