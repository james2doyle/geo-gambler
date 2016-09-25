package main

import (
	"fmt"
	"math"
	"math/rand"

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
	return ctx.OK(Locations)
}

// Get runs the Get action.
func (c *LocationsController) Get(ctx *app.GetLocationsContext) error {
	if len(Locations) <= ctx.ID {
		err := &app.Error{
			Error: "Not Found",
		}
		return ctx.NotFound(err)
	}

	loc := Locations[ctx.ID]
	loc.CanPlay, loc.Distance = CanPlayWithDistance(ctx.Lat, ctx.Long, loc.Latitude, loc.Longitude)

	return ctx.OK(loc)
}

// Play runs the Play action.
func (c *LocationsController) Play(ctx *app.PlayLocationsContext) error {
	if ctx.Value(auth_user) == guest {
		return ctx.Unauthorized(&app.Error{
			Error: "Invalid authorization",
		})
	}

	if len(Locations) <= ctx.ID {
		err := &app.Error{
			Error: "Not Found",
		}
		return ctx.NotFound(err)
	}

	user := ctx.Value(auth_user).(*app.User)

	if user == nil {
		return ctx.Unauthorized(&app.Error{
			Error: "Invalid authorization",
		})
	}

	location := Locations[ctx.ID]
	canPlay, _ := CanPlayWithDistance(&ctx.Lat, &ctx.Long, location.Latitude, location.Longitude)

	if !canPlay {
		res := &app.Playresult{
			Detail:   "You are too far away to play at this location",
			Status:   false,
			Won:      false,
			User:     user,
			Location: location,
		}
		return ctx.OK(res)
	}

	betAmount := ctx.Payload.Bet
	winAmount := betAmount + betAmount/10

	if betAmount > *user.Credit {
		res := &app.Playresult{
			Detail:   "Not enough credit in your wallet",
			Status:   false,
			Won:      false,
			User:     user,
			Location: location,
		}
		return ctx.OK(res)
	}

	if winAmount > location.Wallet {
		res := &app.Playresult{
			Detail:   "Not enough winnings for you to play at this location",
			Status:   false,
			Won:      false,
			User:     user,
			Location: location,
		}
		return ctx.OK(res)
	}

	number := rand.Int31n(10)
	if int(number) != ctx.Payload.Number {
		*user.Credit -= betAmount
		location.Wallet += betAmount
		res := &app.Playresult{
			Detail:   fmt.Sprintf("You lost %d tokens", betAmount),
			Status:   true,
			Won:      false,
			User:     user,
			Location: location,
		}
		return ctx.OK(res)
	}

	*user.Credit += winAmount
	location.Wallet -= winAmount

	if location.Wallet == 0 {
		location.Wallet = 10
	}

	res := &app.Playresult{
		Detail:   fmt.Sprintf("You won %d tokens", winAmount),
		Status:   true,
		Won:      true,
		User:     user,
		Location: location,
	}
	return ctx.OK(res)
}

func CanPlayWithDistance(latUser, longUser *float64, latLoc, longLoc float64) (bool, float64) {
	if latUser == nil || longUser == nil {
		return false, -1.0
	}

	dist := distance(*latUser, *longUser, latLoc, longLoc) * 1000
	return dist <= 50.0, dist
}

func distance(latUser, longUser, latLoc, longLoc float64) (distKM float64) {
	xUser, yUser, zUser := xyz(latUser, longUser)
	xLoc, yLoc, zLoc := xyz(latLoc, longLoc)
	return math.Sqrt(math.Pow(xUser-xLoc, 2) + math.Pow(yUser-yLoc, 2) + math.Pow(zUser-zLoc, 2))
}

func xyz(lat, long float64) (x, y, z float64) {
	theta := deg2rad(lat)
	theta2 := deg2rad(long)
	R := float64(6371)
	x = R * math.Cos(theta) * math.Cos(theta2)
	y = R * math.Cos(theta) * math.Sin(theta2)
	z = R * math.Sin(theta)
	return
}

func deg2rad(deg float64) float64 {
	return deg * math.Pi / 180
}

var Locations = app.LocationCollection{
	&app.Location{
		ID:        0,
		Title:     "Telus Gardens",
		Latitude:  49.281355,
		Longitude: -123.116686,
		Wallet:    50,
		Game:      "numbers",
	},
	&app.Location{
		ID:        1,
		Title:     "Chill Winston",
		Latitude:  49.283623,
		Longitude: -123.103991,
		Wallet:    50,
		Game:      "numbers",
	},
	&app.Location{
		ID:        2,
		Title:     "Microsoft Vancouver",
		Latitude:  49.282191,
		Longitude: -123.118707,
		Wallet:    50,
		Game:      "numbers",
	},
	&app.Location{
		ID:        3,
		Title:     "UBC Robson",
		Latitude:  49.282680,
		Longitude: -123.120789,
		Wallet:    50,
		Game:      "numbers",
	},
	&app.Location{
		ID:        4,
		Title:     "Medina",
		Latitude:  49.280566,
		Longitude: -123.116964,
		Wallet:    50,
		Game:      "numbers",
	},
	&app.Location{
		ID:        5,
		Title:     "TELUS Store Richards",
		Latitude:  49.281017,
		Longitude: -123.116647,
		Wallet:    50,
		Game:      "numbers",
	},
}
