package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Location = MediaType("Location", func() {
	Attributes(func() {
		Attribute("title", String)
		Attribute("game", String, func() {
			Enum("numbers")
		})
		Attribute("payout", Integer)
		Attribute("latitude", Number)
		Attribute("longitude", Number)
	})
	View("default", func() {
		Attribute("title")
		Attribute("game")
		Attribute("payout")
		Attribute("latitude", Number)
		Attribute("longitude", Number)
	})
})

var Result = MediaType("Result", func() {
	Attributes(func() {
		Attribute("detail", String)
	})
	View("default", func() {
		Attribute("detail")
	})
})
