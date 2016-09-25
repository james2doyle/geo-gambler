//go:generate goagen bootstrap -d bitbucket.org/james2doyle/geo-gambler/server/design

package main

import (
	"net/http"

	"golang.org/x/net/context"

	"bitbucket.org/james2doyle/geo-gambler/server/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("GeoGambler")
	service.WithLogger(nil)

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(CORS)
	service.Use(AuthMiddleware)
	service.Use(middleware.Recover())

	// Mount "Locations" controller
	c := NewLocationsController(service)
	app.MountLocationsController(service, c)
	users := NewUsersController(service)
	app.MountUsersController(service, users)

	// Start service
	if err := service.ListenAndServe(":8000"); err != nil {
		service.LogError("startup", "err", err)
	}
}

func CORS(h goa.Handler) goa.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		return h(ctx, w, r)
	}
}

const (
	auth_user = "auth_user"
	guest     = "guest_profile"
)

func AuthMiddleware(h goa.Handler) goa.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		auth := r.URL.Query().Get("auth")
		if auth == "" {
			ctx = context.WithValue(ctx, auth_user, guest)
			return h(ctx, w, r)
		}

		var user *app.User
		for _, u := range Users {
			if auth == *u.ID {
				user = u
			}
		}

		ctx = context.WithValue(ctx, auth_user, user)
		return h(ctx, w, r)
	}
}
