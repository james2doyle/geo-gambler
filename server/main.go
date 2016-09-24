//go:generate goagen bootstrap -d bitbucket.org/james2doyle/geo-gambler/server/design

package main

import (
	"bitbucket.org/james2doyle/geo-gambler/server/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("GeoGambler")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Locations" controller
	c := NewLocationsController(service)
	app.MountLocationsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
