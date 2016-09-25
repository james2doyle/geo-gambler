package main

import (
	"bitbucket.org/james2doyle/geo-gambler/server/app"
	"github.com/goadesign/goa"
	"github.com/satori/go.uuid"
)

// UsersController implements the Users resource.
type UsersController struct {
	*goa.Controller
}

// NewUsersController creates a Users controller.
func NewUsersController(service *goa.Service) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController")}
}

// Get runs the Get action.
func (c *UsersController) Get(ctx *app.GetUsersContext) error {
	if ctx.Value(auth_user) == guest {
		return ctx.Unauthorized(&app.Error{
			Error: "Invalid authorization",
		})
	}

	res := ctx.Value(auth_user).(*app.User)
	return ctx.OK(res)
}

func (c *UsersController) Create(ctx *app.CreateUsersContext) error {
	if ctx.Value(auth_user) != guest {
		return ctx.Unauthorized(&app.Error{
			Error: "Invalid authorization",
		})
	}

	user := &app.User{
		Name:   ctx.Payload.Name,
		Credit: new(int),
		ID:     new(string),
	}
	*user.ID = uuid.NewV4().String()
	*user.Credit = 1000

	Users = append(Users, user)
	return ctx.OK(user)
}

func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	if ctx.Value(auth_user) == guest {
		return ctx.Unauthorized(&app.Error{
			Error: "Invalid authorization",
		})
	}

	user := ctx.Value(auth_user).(*app.User)
	user.Name = ctx.Payload.Name
	return ctx.OK(user)
}

var Users = []*app.User{}
