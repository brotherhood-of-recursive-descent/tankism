package ecs

type ComponentType string

type Component interface {
	Type() ComponentType
}
