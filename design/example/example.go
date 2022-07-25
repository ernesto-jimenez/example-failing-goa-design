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
	Attribute("FieldWithExtension", FieldWithExtension)
	Attribute("External", utils.ExternalType)
	Attribute("SecondExternal", utils.SecondExternalType)
})

var FieldWithExtension = Type("FieldWithExtension", func() {
	Extend(ParticularField)
})

var ParticularField = Type("ParticularField", func() {
	Attribute("BarField", Bar)
})

var Bar = Type("Bar", func() {
	Attribute("Bar", UInt)
	Required("Bar")
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
