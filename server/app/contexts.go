//************************************************************************//
// API "GeoGambler": Application Contexts
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
	"strconv"
)

// FindAllLocationsContext provides the Locations FindAll action context.
type FindAllLocationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Game *string
}

// NewFindAllLocationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the Locations controller FindAll action.
func NewFindAllLocationsContext(ctx context.Context, service *goa.Service) (*FindAllLocationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := FindAllLocationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramGame := req.Params["game"]
	if len(paramGame) > 0 {
		rawGame := paramGame[0]
		rctx.Game = &rawGame
		if rctx.Game != nil {
			if !(*rctx.Game == "numbers") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`game`, *rctx.Game, []interface{}{"numbers"}))
			}
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FindAllLocationsContext) OK(r LocationCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "location; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// GetLocationsContext provides the Locations Get action context.
type GetLocationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewGetLocationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the Locations controller Get action.
func NewGetLocationsContext(ctx context.Context, service *goa.Service) (*GetLocationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetLocationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetLocationsContext) OK(r *Location) error {
	ctx.ResponseData.Header().Set("Content-Type", "location")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// PlayLocationsContext provides the Locations Play action context.
type PlayLocationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *Play
}

// NewPlayLocationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the Locations controller Play action.
func NewPlayLocationsContext(ctx context.Context, service *goa.Service) (*PlayLocationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := PlayLocationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *PlayLocationsContext) OK(r *Result) error {
	ctx.ResponseData.Header().Set("Content-Type", "result")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
