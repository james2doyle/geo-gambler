//************************************************************************//
// API "GeoGambler": Application Media Types
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

import "github.com/goadesign/goa"

// Location media type (default view)
//
// Identifier: location; view=default
type Location struct {
	Game      *string  `form:"game,omitempty" json:"game,omitempty" xml:"game,omitempty"`
	Latitude  *float64 `form:"latitude,omitempty" json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *float64 `form:"longitude,omitempty" json:"longitude,omitempty" xml:"longitude,omitempty"`
	Payout    *int     `form:"payout,omitempty" json:"payout,omitempty" xml:"payout,omitempty"`
	Title     *string  `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the Location media type instance.
func (mt *Location) Validate() (err error) {
	if mt.Game != nil {
		if !(*mt.Game == "numbers") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.game`, *mt.Game, []interface{}{"numbers"}))
		}
	}
	return
}

// LocationCollection is the media type for an array of Location (default view)
//
// Identifier: location; type=collection; view=default
type LocationCollection []*Location

// Validate validates the LocationCollection media type instance.
func (mt LocationCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Game != nil {
			if !(*e.Game == "numbers") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response[*].game`, *e.Game, []interface{}{"numbers"}))
			}
		}
	}
	return
}

// Result media type (default view)
//
// Identifier: result; view=default
type Result struct {
	Detail *string `form:"detail,omitempty" json:"detail,omitempty" xml:"detail,omitempty"`
}
