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

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

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

// DecodeError decodes the Error instance encoded in resp body.
func (c *Client) DecodeError(resp *http.Response) (*Error, error) {
	var decoded Error
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeLocation decodes the Location instance encoded in resp body.
func (c *Client) DecodeLocation(resp *http.Response) (*Location, error) {
	var decoded Location
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeLocationCollection decodes the LocationCollection instance encoded in resp body.
func (c *Client) DecodeLocationCollection(resp *http.Response) (LocationCollection, error) {
	var decoded LocationCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Playresult media type (default view)
//
// Identifier: playresult; view=default
type Playresult struct {
	Detail   string    `form:"detail" json:"detail" xml:"detail"`
	Location *Location `form:"location" json:"location" xml:"location"`
	Status   bool      `form:"status" json:"status" xml:"status"`
	User     *User     `form:"user" json:"user" xml:"user"`
	Won      bool      `form:"won" json:"won" xml:"won"`
}

// Validate validates the Playresult media type instance.
func (mt *Playresult) Validate() (err error) {
	if mt.Detail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "detail"))
	}
	if mt.User == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user"))
	}
	if mt.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "location"))
	}

	if mt.Location != nil {
		if err2 := mt.Location.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.User != nil {
		if err2 := mt.User.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodePlayresult decodes the Playresult instance encoded in resp body.
func (c *Client) DecodePlayresult(resp *http.Response) (*Playresult, error) {
	var decoded Playresult
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
