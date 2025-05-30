package module

import (
	"fmt"
)

type FactoryFunc func() Module

var registry = map[string]FactoryFunc{}

//NOTE: each module register itself through the init() func
// ```
// func init() {
//   Register("command", func() Module { return CommandModule{} })
// }
// ```

func Register(name string, factory FactoryFunc) {
	registry[name] = factory
}

func Get(name string) (Module, error) {
	factory, ok := registry[name]
	if !ok {
		return nil, &ModuleNotFoundError{Name: name}
	}

	return factory(), nil
}

type ModuleNotFoundError struct {
	Name string
}

func (e *ModuleNotFoundError) Error() string {
	return fmt.Sprintf("module '%s' not found", e.Name)
}
