//************************************************************************//
// API "GeoGambler": Locations TestHelpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=bitbucket.org/james2doyle/geo-gambler/server/design
// --out=$(GOPATH)/src/bitbucket.org/james2doyle/geo-gambler/server
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package test

import (
	"bitbucket.org/james2doyle/geo-gambler/server/app"
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// FindAllLocationsOK runs the method FindAll of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func FindAllLocationsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.LocationsController, game *string) (http.ResponseWriter, app.LocationCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if game != nil {
		sliceVal := []string{*game}
		query["game"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/location"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if game != nil {
		sliceVal := []string{*game}
		prms["game"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "LocationsTest"), rw, req, prms)
	findAllCtx, err := app.NewFindAllLocationsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.FindAll(findAllCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.LocationCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.LocationCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.LocationCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// GetLocationsNotFound runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetLocationsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.LocationsController, id int, lat *float64, long *float64) (http.ResponseWriter, *app.Error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if lat != nil {
		sliceVal := []string{fmt.Sprintf("%v", *lat)}
		query["lat"] = sliceVal
	}
	if long != nil {
		sliceVal := []string{fmt.Sprintf("%v", *long)}
		query["long"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/location/%v", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if lat != nil {
		sliceVal := []string{fmt.Sprintf("%v", *lat)}
		prms["lat"] = sliceVal
	}
	if long != nil {
		sliceVal := []string{fmt.Sprintf("%v", *long)}
		prms["long"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "LocationsTest"), rw, req, prms)
	getCtx, err := app.NewGetLocationsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Get(getCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt *app.Error
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Error", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// GetLocationsOK runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetLocationsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.LocationsController, id int, lat *float64, long *float64) (http.ResponseWriter, *app.Location) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if lat != nil {
		sliceVal := []string{fmt.Sprintf("%v", *lat)}
		query["lat"] = sliceVal
	}
	if long != nil {
		sliceVal := []string{fmt.Sprintf("%v", *long)}
		query["long"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/location/%v", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if lat != nil {
		sliceVal := []string{fmt.Sprintf("%v", *lat)}
		prms["lat"] = sliceVal
	}
	if long != nil {
		sliceVal := []string{fmt.Sprintf("%v", *long)}
		prms["long"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "LocationsTest"), rw, req, prms)
	getCtx, err := app.NewGetLocationsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Get(getCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Location
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Location)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Location", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// PlayLocationsNotFound runs the method Play of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PlayLocationsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.LocationsController, id int, lat float64, long float64, payload *app.Play) (http.ResponseWriter, *app.Error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{fmt.Sprintf("%v", lat)}
		query["lat"] = sliceVal
	}
	{
		sliceVal := []string{fmt.Sprintf("%v", long)}
		query["long"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/location/play/%v", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	{
		sliceVal := []string{fmt.Sprintf("%v", lat)}
		prms["lat"] = sliceVal
	}
	{
		sliceVal := []string{fmt.Sprintf("%v", long)}
		prms["long"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "LocationsTest"), rw, req, prms)
	playCtx, err := app.NewPlayLocationsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	playCtx.Payload = payload

	// Perform action
	err = ctrl.Play(playCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt *app.Error
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Error", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// PlayLocationsOK runs the method Play of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PlayLocationsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.LocationsController, id int, lat float64, long float64, payload *app.Play) (http.ResponseWriter, *app.Playresult) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 200 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{fmt.Sprintf("%v", lat)}
		query["lat"] = sliceVal
	}
	{
		sliceVal := []string{fmt.Sprintf("%v", long)}
		query["long"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/location/play/%v", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	{
		sliceVal := []string{fmt.Sprintf("%v", lat)}
		prms["lat"] = sliceVal
	}
	{
		sliceVal := []string{fmt.Sprintf("%v", long)}
		prms["long"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "LocationsTest"), rw, req, prms)
	playCtx, err := app.NewPlayLocationsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	playCtx.Payload = payload

	// Perform action
	err = ctrl.Play(playCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Playresult
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Playresult)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Playresult", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// PlayLocationsUnauthorized runs the method Play of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PlayLocationsUnauthorized(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.LocationsController, id int, lat float64, long float64, payload *app.Play) (http.ResponseWriter, *app.Error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 401 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{fmt.Sprintf("%v", lat)}
		query["lat"] = sliceVal
	}
	{
		sliceVal := []string{fmt.Sprintf("%v", long)}
		query["long"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/location/play/%v", id),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	{
		sliceVal := []string{fmt.Sprintf("%v", lat)}
		prms["lat"] = sliceVal
	}
	{
		sliceVal := []string{fmt.Sprintf("%v", long)}
		prms["long"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "LocationsTest"), rw, req, prms)
	playCtx, err := app.NewPlayLocationsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	playCtx.Payload = payload

	// Perform action
	err = ctrl.Play(playCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}
	var mt *app.Error
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Error", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}
