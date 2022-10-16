package ecs

type EntityManager struct {
	entities []*Entity
}

func NewEntityManager(entities []Entity) *EntityManager {
	em := &EntityManager{}

	for _, v := range entities {
		e := em.NewEntity()

		for _, c := range v.components {
			e.AddComponent(c)
		}
	}

	return em
}

func (em *EntityManager) NewEntity() *Entity {
	e := newEntity()
	em.entities = append(em.entities, e)
	return e
}

func (em *EntityManager) Size() int {
	return len(em.entities)
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

func (em *EntityManager) RemoveEntity(e *Entity) {

	idx := -1
	for i, entity := range em.entities {
		if entity == e {
			idx = i
			break
		}
	}

	if idx != -1 {
		em.entities = append(em.entities[:idx], em.entities[idx+1:]...)
	}
}
