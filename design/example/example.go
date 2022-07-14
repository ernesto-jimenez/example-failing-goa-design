package example

import (
	"github.com/ernesto-jimenez/example-failing-goa-design/design/utils"
	. "goa.design/goa/v3/dsl"
)

var _ = API("Example", func() {
	HTTP()
})

var ExampleType = Type("ExampleType", func() {
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
