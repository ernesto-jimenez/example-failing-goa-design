package example

import (
	"github.com/ernesto-jimenez/example-failing-goa-design/design/utils"
	. "goa.design/goa/v3/dsl"
)

var _ = API("Example", func() {
	HTTP()
})

var DateTime = Type("DateTime", String, func() {
	Format(FormatDateTime)
})

var Extension = Type("Extension", func() {
	Attribute("DateField", DateTime)
	Required("DateField")
})

var ExampleType = Type("ExampleType", func() {
	Extend(Extension)
	Attribute("External", utils.ExternalType)
	Attribute("SecondExternal", utils.SecondExternalType)
})

var _ = Service("FooService", func() {
	Method("FooMethod", func() {
		Payload(func() {
			Extend(ExampleType)
		})
		Result(ArrayOf(ExampleType))
		HTTP(func() {
			POST("/example")
		})
	})
})
