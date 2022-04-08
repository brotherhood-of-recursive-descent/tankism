package ecs

type EntityManager struct {
	entities []*Entity
}

func (em *EntityManager) NewEntity() *Entity {
	e := newEntity()
	em.entities = append(em.entities, e)
	return e
}

func (em *EntityManager) FindByComponents(components ...ComponentType) []*Entity {
	candidates := []*Entity{}

	for _, entity := range em.entities {

		found := 0
		for _, c := range components {
			if entity.HasComponent(c) {
				found++
			}
		}

		// works, because component is unique
		if found == len(components) {
			candidates = append(candidates, entity)
		}

	}

	return candidates
}
