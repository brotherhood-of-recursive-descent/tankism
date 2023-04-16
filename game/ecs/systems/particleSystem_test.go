package systems_test

import (
	"testing"
	"time"

	"github.com/co0p/tankism/game/ecs/components"
	"github.com/co0p/tankism/game/ecs/systems"
	"github.com/co0p/tankism/lib/ecs"
)

func Test_ParticleSystem_updateTime(t *testing.T) {

	// given
	past := time.Now().Add(-time.Second)
	em := ecs.NewEntityManager(nil)
	system := systems.ParticleSystem{EntityManager: em}

	entity := em.NewEntity()
	entity.AddComponents(
		&components.Transform{},
		&components.ParticleEmitter{
			Spawn_interval: time.Millisecond,
			Last_emitted:   past,
		},
	)

	// when
	err := system.Update()

	// then
	if err != nil {
		t.Errorf("expected err to be nil, got %s\n", err)
	}

	emitter := entity.GetComponent(components.ParticleEmitterType).(*components.ParticleEmitter)
	emitter.Last_emitted.After(past)
}
