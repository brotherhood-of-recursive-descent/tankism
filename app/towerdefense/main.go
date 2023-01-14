package main

import (
	"log"

	"github.com/co0p/tankism/game"
	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

type TowerDefenseDemo struct {
	game.GameScene
}

type XY struct {
	X int
	Y int
}

func (demo *TowerDefenseDemo) Init() error {

	demo.Systems = append(demo.Systems,
		&systems.SpriteRenderer{EntityManager: &demo.EntityManager},
		&systems.PerformanceMonitor{EntityManager: &demo.EntityManager},
		&systems.TextRenderer{EntityManager: &demo.EntityManager},
		&systems.Controller{EntityManager: &demo.EntityManager},
		&systems.Shaker{EntityManager: &demo.EntityManager},
		&systems.AISystem{EntityManager: &demo.EntityManager},
		&systems.CleanupSystem{EntityManager: &demo.EntityManager},
		&systems.CollisionDetection{EntityManager: &demo.EntityManager},
		&systems.CollisionResolution{EntityManager: &demo.EntityManager},
	)

	fps := demo.EntityManager.NewEntity()
	game.FPSCounter(fps, 1024)

	tilemap := demo.EntityManager.NewEntity()
	game.NewMap(tilemap, game.Tilemap{}, 1024, 1024)

	tank := demo.EntityManager.NewEntity()
	game.NewTankWithPosition(tank, 100, 400)
	tank.AddComponent(&components.Target{GroupId: 1})

	bigTank := demo.EntityManager.NewEntity()
	game.NewBigTank(bigTank, 700, 100)
	bigTank.AddComponent(&components.AI{TargetGroup: 1})

	bigTank2 := demo.EntityManager.NewEntity()
	game.NewBigTank(bigTank2, 750, 800)
	bigTank2.AddComponent(&components.AI{TargetGroup: 1})

	goal := demo.EntityManager.NewEntity()
	game.NewCrateMetal(goal, 900, 500)

	for _, v := range []XY{
		{400, 600},
		{600, 50},
		{600, 120},
		{630, 800},
		{640, 240},
		{750, 240},
		{820, 250},
		{820, 720}} {
		crate := demo.EntityManager.NewEntity()
		game.NewCrate(crate, float64(v.X), float64(v.Y))
	}

	game.NewTree(demo.EntityManager.NewEntity(), 10, 10)
	game.NewTree(demo.EntityManager.NewEntity(), 160, 50)
	game.NewTree(demo.EntityManager.NewEntity(), 50, 130)

	for _, v := range []XY{
		{400, 520},
		{460, 600},
		{700, 740},
		{760, 700},
		{900, 720},
	} {
		barricade := demo.EntityManager.NewEntity()
		game.NewBarricade(barricade, float64(v.X), float64(v.Y))
	}

	return nil
}

func main() {

	demo := TowerDefenseDemo{}
	game := game.NewGame()
	game.AddScene("TowerDefense", &demo)
	game.SetScene("TowerDefense")

	ebiten.SetFullscreen(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatalf("failed to start game: %s", err)
	}
}
