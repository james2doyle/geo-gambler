package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Play = Type("Play", func() {
	Attribute("number", Integer)
	Required("number")
})
