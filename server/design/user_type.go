package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Play = Type("Play", func() {
	Attribute("latitude", Number)
	Attribute("longitude", Number)
	Attribute("number", Integer)
	Required("latitude", "longitude", "number")
})
