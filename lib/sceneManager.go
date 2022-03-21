package lib

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneManager struct {
	currentScene Scene
	nextScene    Scene
	scenes       map[string]Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes: make(map[string]Scene),
	}
}

func (sm *SceneManager) RegisterScene(sceneKey string, scene Scene) {
	sm.scenes[sceneKey] = scene
}

func (sm *SceneManager) ChangeScene(sceneKey string) {
	scene, ok := sm.scenes[sceneKey]
	if !ok {
		panic("Invalid sceneKey " + sceneKey)
	}

	if sm.currentScene == nil {
		sm.currentScene = scene
		sm.currentScene.Init()
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
		sm.currentScene.Init()
		sm.nextScene = nil
	}
	return sm.currentScene.Update()
}

func (sm *SceneManager) WindowDimension() (int, int) {
	return sm.currentScene.WindowDimension()
}

func (sm *SceneManager) SetWindowDimension(i int, i2 int) {
	sm.currentScene.SetWindowDimension(i, i2)
}
