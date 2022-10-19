package lib

import (
	"github.com/co0p/tankism/lib/sound"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	currentScene Scene
	nextScene    Scene
	scenes       map[string]Scene
	ScreenWidth  int
	ScreenHeight int

	SoundManager sound.SoundManager
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes: make(map[string]Scene),
	}
}

func (sm *SceneManager) RegisterScene(sceneKey string, scene Scene) {
	sm.scenes[sceneKey] = scene

	if len(sm.scenes) == 1 {
		sm.ChangeScene(sceneKey)
	}
}

func (sm *SceneManager) ChangeScene(sceneKey string) {
	scene, ok := sm.scenes[sceneKey]
	if !ok {
		panic("Invalid sceneKey " + sceneKey)
	}

	if sm.currentScene == nil {
		sm.currentScene = scene
		sm.currentScene.Init(sm)
	} else {
		sm.nextScene = scene
	}
}

func (sm *SceneManager) Draw(screen *ebiten.Image) {
	sm.currentScene.Draw(screen)
}

func (sm *SceneManager) Update() error {
	if sm.nextScene != nil {
		sm.currentScene = sm.nextScene
		sm.currentScene.Init(sm)
		sm.nextScene = nil
	}
	return sm.currentScene.Update()
}

func (sm *SceneManager) SetWindowDimension(w int, h int) {
	sm.ScreenWidth = w
	sm.ScreenHeight = h
}

func (sm *SceneManager) CurrentScene() Scene {
	return sm.currentScene
}
