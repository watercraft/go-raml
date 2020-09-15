package codegen

import (
	"github.com/watercraft/go-raml/codegen/python"
	"github.com/watercraft/go-raml/raml"
)

func GeneratePythonCapnp(apiDef *raml.APIDefinition, dir string) error {
	return python.GeneratePythonCapnpClasses(apiDef, dir)
}
