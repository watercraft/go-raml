package codegen

import (
	"github.com/watercraft/go-raml/codegen/capnp"
	"github.com/watercraft/go-raml/raml"
)

// GenerateCapnp generates capnp schema from RAML specs
func GenerateCapnp(apiDef *raml.APIDefinition, dir, lang, pkg string) error {
	return capnp.GenerateCapnp(apiDef, dir, lang, pkg)
}
