package ecs

type Entity struct {
	components map[ComponentType]Component
}

func newEntity() *Entity {
	return &Entity{
		components: make(map[ComponentType]Component),
	}
}

func (e *Entity) AddComponent(c Component) {
	e.components[c.Type()] = c
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
