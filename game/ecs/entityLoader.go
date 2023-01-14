package ecs

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/co0p/tankism/lib/ecs"
)

type EntityLoader struct {
	registry map[ecs.ComponentType]LoadableComponent
}

type LoadableComponent interface {
	Type() ecs.ComponentType
	Load([]byte) LoadableComponent
}

func NewEntityLoader() EntityLoader {
	el := EntityLoader{
		registry: make(map[ecs.ComponentType]LoadableComponent),
	}
	return el
}

func (e *EntityLoader) Register(cs ...LoadableComponent) {
	for _, v := range cs {
		e.registry[v.Type()] = v
	}
}

// Load returns a new entity from given by adding json unmarshalled components
func (el *EntityLoader) Load(e *ecs.Entity, data []byte) error {
	cs := []ecs.Component{}
	var componentMap map[ecs.ComponentType]any

	err := json.Unmarshal(data, &componentMap)
	if err != nil {
		return errors.New("failed to load entity: " + err.Error())
	}

	for k, v := range componentMap {

		t, ok := el.registry[k]
		if !ok {
			log.Fatalf("could not find type %T in registry\n", t)
		}
		v1 := v.(map[string]interface{})

		jjsss, _ := json.Marshal(v1)
		component := t.Load(jjsss)

		cs = append(cs, component)
	}

	e.AddComponents(cs...)
	return err
}
