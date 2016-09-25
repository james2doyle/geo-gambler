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

// Error media type (default view)
//
// Identifier: error; view=default
type Error struct {
	Error string `form:"error" json:"error" xml:"error"`
}

// Validate validates the Error media type instance.
func (mt *Error) Validate() (err error) {
	if mt.Error == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "error"))
	}

	return
}

// Location media type (default view)
//
// Identifier: location; view=default
type Location struct {
	CanPlay   bool    `form:"can_play" json:"can_play" xml:"can_play"`
	Distance  float64 `form:"distance" json:"distance" xml:"distance"`
	Game      string  `form:"game" json:"game" xml:"game"`
	ID        int     `form:"id" json:"id" xml:"id"`
	Latitude  float64 `form:"latitude" json:"latitude" xml:"latitude"`
	Longitude float64 `form:"longitude" json:"longitude" xml:"longitude"`
	Title     string  `form:"title" json:"title" xml:"title"`
	Wallet    int     `form:"wallet" json:"wallet" xml:"wallet"`
}

// Validate validates the Location media type instance.
func (mt *Location) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Game == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "game"))
	}

	if !(mt.Game == "numbers") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.game`, mt.Game, []interface{}{"numbers"}))
	}
	if mt.ID < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 0, true))
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
		if e.Title == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "title"))
		}
		if e.Game == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "game"))
		}

		if !(e.Game == "numbers") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response[*].game`, e.Game, []interface{}{"numbers"}))
		}
		if e.ID < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response[*].id`, e.ID, 0, true))
		}
	}
	return
}

// Result media type (default view)
//
// Identifier: result; view=default
type Result struct {
	Detail string `form:"detail" json:"detail" xml:"detail"`
}

// Validate validates the Result media type instance.
func (mt *Result) Validate() (err error) {
	if mt.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "detail"))
	}

	return
}

// User media type (default view)
//
// Identifier: user; view=default
type User struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   string  `form:"name" json:"name" xml:"name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.Credit != nil {
		if *mt.Credit < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.credit`, *mt.Credit, 0, true))
		}
	}
	if mt.ID != nil {
		if ok := goa.ValidatePattern(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`, *mt.ID); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, *mt.ID, `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`))
		}
	}
	return
}
