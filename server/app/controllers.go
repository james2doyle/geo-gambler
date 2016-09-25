//************************************************************************//
// API "GeoGambler": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=bitbucket.org/james2doyle/geo-gambler/server/design
// --out=$(GOPATH)/src/bitbucket.org/james2doyle/geo-gambler/server
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// LocationsController is the controller interface for the Locations actions.
type LocationsController interface {
	goa.Muxer
	FindAll(*FindAllLocationsContext) error
	Get(*GetLocationsContext) error
	Play(*PlayLocationsContext) error
}

// MountLocationsController "mounts" a Locations resource controller on the given service.
func MountLocationsController(service *goa.Service, ctrl LocationsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewFindAllLocationsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.FindAll(rctx)
	}
	service.Mux.Handle("GET", "/location", ctrl.MuxHandler("FindAll", h, nil))
	service.LogInfo("mount", "ctrl", "Locations", "action", "FindAll", "route", "GET /location")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetLocationsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/location/:id", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Locations", "action", "Get", "route", "GET /location/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPlayLocationsContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Play)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Play(rctx)
	}
	service.Mux.Handle("POST", "/location/play/:id", ctrl.MuxHandler("Play", h, unmarshalPlayLocationsPayload))
	service.LogInfo("mount", "ctrl", "Locations", "action", "Play", "route", "POST /location/play/:id")
}

// unmarshalPlayLocationsPayload unmarshals the request body into the context request data Payload field.
func unmarshalPlayLocationsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &play{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	Create(*CreateUsersContext) error
	Get(*GetUsersContext) error
	Update(*UpdateUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUsersContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/user", ctrl.MuxHandler("Create", h, unmarshalCreateUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Create", "route", "POST /user")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUsersContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/user/me", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "Get", "route", "GET /user/me")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUsersContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("POST", "/user/me", ctrl.MuxHandler("Update", h, unmarshalUpdateUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Update", "route", "POST /user/me")
}

// unmarshalCreateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
