package di

import (
	"go.uber.org/fx"
)

type Module struct {
	Imports     []Module
	Provides    []interface{}
	Controllers []interface{}
	Routing     interface{}
}

type Modules []Module
type Deps []interface{}
type Routes []interface{}

func combineProvides(module Module) []interface{} {
	var constructors []interface{}

	if module.Provides != nil {
		constructors = append(constructors, module.Provides...)
	}
	if module.Controllers != nil {
		constructors = append(constructors, module.Controllers...)
	}

	return constructors
}

func combineInvokes(module Module) []interface{} {
	var functions []interface{}
	if module.Routing != nil {
		functions = append(functions, module.Routing)
	}
	return functions
}

func combineImports(module Module) ([]interface{}, []interface{}) {
	var invokes []interface{}
	var provides []interface{}

	invokes = append(invokes, combineInvokes(module)...)
	provides = append(provides, combineProvides(module)...)

	if module.Imports == nil {
		return invokes, provides
	}

	for _, module := range module.Imports {
		moduleInvokes, moduleProvides := combineImports(module)
		invokes = append(invokes, moduleInvokes...)
		provides = append(provides, moduleProvides...)
	}

	return invokes, provides
}

func Init(module Module) *fx.App {
	moduleInvokes, moduleProvides := combineImports(module)
	return fx.New(
		fx.Provide(moduleProvides...),
		fx.Invoke(moduleInvokes...),
	)
}
