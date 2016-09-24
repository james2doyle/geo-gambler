package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Locations", func() {
	BasePath("/locations")
	Action("Get", func() {
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, Location)
	})

	Action("FindAll", func() {
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("game", String, func() {
				Enum("numbers")
			})
		})
		Response(OK, CollectionOf(Location))
	})

	Action("Play", func() {
		Routing(
			POST("/play/:id"),
		)
		Params(func() {
			Param("id", Integer)
		})
		Payload(Play)
		Response(OK, Result)
	})
})
