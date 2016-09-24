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
	service.Mux.Handle("GET", "/locations", ctrl.MuxHandler("FindAll", h, nil))
	service.LogInfo("mount", "ctrl", "Locations", "action", "FindAll", "route", "GET /locations")

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
	service.Mux.Handle("GET", "/locations/:id", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Locations", "action", "Get", "route", "GET /locations/:id")

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
	service.Mux.Handle("POST", "/locations/play/:id", ctrl.MuxHandler("Play", h, unmarshalPlayLocationsPayload))
	service.LogInfo("mount", "ctrl", "Locations", "action", "Play", "route", "POST /locations/play/:id")
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
