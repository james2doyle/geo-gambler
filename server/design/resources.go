package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Locations", func() {
	BasePath("/location")
	Action("Get", func() {
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, func() {
				Minimum(0)
			})
			Param("lat", Number)
			Param("long", Number)
		})
		Response(OK, Location)
		Response(NotFound, Error)
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
			Param("id", Integer, func() {
				Minimum(0)
			})
			Param("lat", Number)
			Param("long", Number)
			Required("lat", "long")
		})
		Payload(Play)
		Response(OK, PlayResult)
		Response(NotFound, Error)
		Response(Unauthorized, Error)
	})
})

var _ = Resource("Users", func() {
	BasePath("/user")
	Action("Get", func() {
		Routing(
			GET("/me"),
		)
		Response(OK, User)
		Response(Unauthorized, Error)
		Response(NotFound, Error)
	})

	Action("Create", func() {
		Routing(
			POST("/"),
		)
		Payload(User)
		Response(OK, User)
		Response(Unauthorized, Error)
	})

	Action("Update", func() {
		Routing(
			POST("/me"),
		)
		Payload(User)
		Response(OK, User)
		Response(Unauthorized, Error)
	})
})
