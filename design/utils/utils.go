package utils

import (
	"goa.design/goa/v3/dsl"
)

var ExternalType = dsl.Type("ExternalType", func() {
	dsl.Meta("struct:pkg:path", "types")
	dsl.Attribute("Field", dsl.String)
	dsl.Required("Field")
})
