package design

import (
	"github.com/asaskevich/govalidator"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Location = MediaType("Location", func() {
	Attributes(func() {
		Attribute("id", Integer, func() {
			Minimum(0)
		})
		Attribute("title", String)
		Attribute("game", String, func() {
			Enum("numbers")
		})
		Attribute("wallet", Integer)
		Attribute("distance", Number)
		Attribute("latitude", Number)
		Attribute("longitude", Number)
		Attribute("can_play", Boolean, func() {
			Default(false)
		})
		Required("id", "title", "game", "wallet", "latitude", "longitude", "can_play", "distance")
	})
	View("default", func() {
		Attribute("id")
		Attribute("distance")
		Attribute("title")
		Attribute("game")
		Attribute("wallet")
		Attribute("latitude")
		Attribute("longitude")
		Attribute("can_play")
	})
})

var PlayResult = MediaType("PlayResult", func() {
	Attributes(func() {
		Attribute("detail", String)
		Attribute("status", Boolean)
		Attribute("won", Boolean)
		Attribute("user", User)
		Attribute("location", Location)
		Required("detail", "status", "won", "user", "location")
	})
	View("default", func() {
		Attribute("detail")
		Attribute("status")
		Attribute("won")
		Attribute("user")
		Attribute("location")
	})
})

var Error = MediaType("Error", func() {
	Attributes(func() {
		Attribute("error", String)
		Required("error")
	})

	View("default", func() {
		Attribute("error")
	})
})

var User = MediaType("User", func() {
	Attributes(func() {
		Attribute("id", String, func() {
			Pattern(govalidator.UUID)
		})
		Attribute("name")
		Attribute("credit", Integer, func() {
			Minimum(0)
		})
		Required("name")
		View("default", func() {
			Attribute("id")
			Attribute("name")
			Attribute("credit")
		})
	})
})
