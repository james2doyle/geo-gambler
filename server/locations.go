package main

import (
	"bitbucket.org/james2doyle/geo-gambler/server/app"
	"github.com/goadesign/goa"
)

// LocationsController implements the Locations resource.
type LocationsController struct {
	*goa.Controller
}

// NewLocationsController creates a Locations controller.
func NewLocationsController(service *goa.Service) *LocationsController {
	return &LocationsController{Controller: service.NewController("LocationsController")}
}

// FindAll runs the FindAll action.
func (c *LocationsController) FindAll(ctx *app.FindAllLocationsContext) error {
	// LocationsController_FindAll: start_implement

	// Put your logic here

	// LocationsController_FindAll: end_implement
	res := app.LocationCollection{}
	return ctx.OK(res)
}

// Get runs the Get action.
func (c *LocationsController) Get(ctx *app.GetLocationsContext) error {
	// LocationsController_Get: start_implement

	// Put your logic here

	// LocationsController_Get: end_implement
	res := &app.Location{}
	return ctx.OK(res)
}

// Play runs the Play action.
func (c *LocationsController) Play(ctx *app.PlayLocationsContext) error {
	// LocationsController_Play: start_implement

	// Put your logic here

	// LocationsController_Play: end_implement
	res := &app.Result{}
	return ctx.OK(res)
}
