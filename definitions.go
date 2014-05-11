package gory

import (
	"fmt"
	"reflect"
	"strings"
)

type definition struct {
	structType reflect.Type
	factory    Factory
}

func newDefinition(instance interface{}, factory Factory) *definition {
	return &definition{structType: reflect.TypeOf(instance), factory: factory}
}

type definitions map[string]*definition

func (defs definitions) set(name string, def *definition) {
	_, exists := defs[name]
	if exists {
		message := fmt.Sprintf("gory: '%s' has already been defined", name)
		panic(message)
	} else {
		defs[name] = def
	}
}

func (defs definitions) get(name string) *definition {
	def, exists := defs[name]
	if exists {
		return def
	} else {
		defined := strings.Join(defs.defined(), ", ")
		message := fmt.Sprintf("gory: '%s' is undefined. Defined factories: %s", name, defined)
		panic(message)
	}
}

func (defs definitions) defined() []string {
	names := []string{}
	for name, _ := range defs {
		names = append(names, name)
	}
	return names
}
