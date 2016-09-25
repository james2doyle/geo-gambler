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
	ID   int
	Lat  *float64
	Long *float64
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
		if rctx.ID < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`id`, rctx.ID, 0, true))
		}
	}
	paramLat := req.Params["lat"]
	if len(paramLat) > 0 {
		rawLat := paramLat[0]
		if lat, err2 := strconv.ParseFloat(rawLat, 64); err2 == nil {
			tmp2 := &lat
			rctx.Lat = tmp2
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("lat", rawLat, "number"))
		}
	}
	paramLong := req.Params["long"]
	if len(paramLong) > 0 {
		rawLong := paramLong[0]
		if long, err2 := strconv.ParseFloat(rawLong, 64); err2 == nil {
			tmp3 := &long
			rctx.Long = tmp3
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("long", rawLong, "number"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetLocationsContext) OK(r *Location) error {
	ctx.ResponseData.Header().Set("Content-Type", "location")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetLocationsContext) NotFound(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// PlayLocationsContext provides the Locations Play action context.
type PlayLocationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Lat     float64
	Long    float64
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
		if rctx.ID < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`id`, rctx.ID, 0, true))
		}
	}
	paramLat := req.Params["lat"]
	if len(paramLat) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("lat"))
	} else {
		rawLat := paramLat[0]
		if lat, err2 := strconv.ParseFloat(rawLat, 64); err2 == nil {
			rctx.Lat = lat
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("lat", rawLat, "number"))
		}
	}
	paramLong := req.Params["long"]
	if len(paramLong) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("long"))
	} else {
		rawLong := paramLong[0]
		if long, err2 := strconv.ParseFloat(rawLong, 64); err2 == nil {
			rctx.Long = long
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("long", rawLong, "number"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *PlayLocationsContext) OK(r *Result) error {
	ctx.ResponseData.Header().Set("Content-Type", "result")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *PlayLocationsContext) Unauthorized(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *PlayLocationsContext) NotFound(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// CreateUsersContext provides the Users Create action context.
type CreateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUsersPayload
}

// NewCreateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the Users controller Create action.
func NewCreateUsersContext(ctx context.Context, service *goa.Service) (*CreateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createUsersPayload is the Users Create action payload.
type createUsersPayload struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUsersPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Credit != nil {
		if *payload.Credit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.credit`, *payload.Credit, 0, true))
		}
	}
	if payload.ID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, *payload.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.id`, *payload.ID, `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`))
		}
	}
	return
}

// Publicize creates CreateUsersPayload from createUsersPayload
func (payload *createUsersPayload) Publicize() *CreateUsersPayload {
	var pub CreateUsersPayload
	if payload.Credit != nil {
		pub.Credit = payload.Credit
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateUsersPayload is the Users Create action payload.
type CreateUsersPayload struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   string  `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUsersPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Credit != nil {
		if *payload.Credit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.credit`, *payload.Credit, 0, true))
		}
	}
	if payload.ID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, *payload.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.id`, *payload.ID, `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`))
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateUsersContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "user")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateUsersContext) Unauthorized(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// GetUsersContext provides the Users Get action context.
type GetUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewGetUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the Users controller Get action.
func NewGetUsersContext(ctx context.Context, service *goa.Service) (*GetUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetUsersContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "user")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *GetUsersContext) Unauthorized(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetUsersContext) NotFound(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// UpdateUsersContext provides the Users Update action context.
type UpdateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UpdateUsersPayload
}

// NewUpdateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the Users controller Update action.
func NewUpdateUsersContext(ctx context.Context, service *goa.Service) (*UpdateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// updateUsersPayload is the Users Update action payload.
type updateUsersPayload struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUsersPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Credit != nil {
		if *payload.Credit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.credit`, *payload.Credit, 0, true))
		}
	}
	if payload.ID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, *payload.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.id`, *payload.ID, `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`))
		}
	}
	return
}

// Publicize creates UpdateUsersPayload from updateUsersPayload
func (payload *updateUsersPayload) Publicize() *UpdateUsersPayload {
	var pub UpdateUsersPayload
	if payload.Credit != nil {
		pub.Credit = payload.Credit
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateUsersPayload is the Users Update action payload.
type UpdateUsersPayload struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   string  `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUsersPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	if payload.Credit != nil {
		if *payload.Credit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`raw.credit`, *payload.Credit, 0, true))
		}
	}
	if payload.ID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, *payload.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`raw.id`, *payload.ID, `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`))
		}
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUsersContext) OK(r *User) error {
	ctx.ResponseData.Header().Set("Content-Type", "user")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdateUsersContext) Unauthorized(r *Error) error {
	ctx.ResponseData.Header().Set("Content-Type", "error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}
