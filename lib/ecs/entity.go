package ecs

import (
	"github.com/co0p/tankism/lib/uuid"
)

type Entity struct {
	ID         string
	components map[ComponentType]Component
}

func newEntity() *Entity {
	return &Entity{
		ID:         uuid.New(),
		components: make(map[ComponentType]Component),
	}
}

func (e *Entity) AddComponent(c Component) {
	e.components[c.Type()] = c
}

func (e *Entity) AddComponents(cs ...Component) {

	for _, c := range cs {
		e.components[c.Type()] = c
	}
}

func (e *Entity) RemoveComponent(c ComponentType) {
	delete(e.components, c)
}

func (e *Entity) HasComponent(c ComponentType) bool {
	_, ok := e.components[c]
	return ok
}

func (e *Entity) GetComponent(c ComponentType) Component {
	if e.HasComponent(c) {
		return e.components[c]
	}
	panic("failed to get component " + c)
}
