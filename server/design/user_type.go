package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Play = Type("Play", func() {
	Attribute("number", Integer, func() {
		Minimum(0)
		Maximum(9)
	})
	Attribute("bet", Integer, func() {
		Minimum(10)
	})
	Required("number", "bet")
})
