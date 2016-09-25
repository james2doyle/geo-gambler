//************************************************************************//
// API "GeoGambler": Application User Types
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

import "github.com/goadesign/goa"

// play user type.
type play struct {
	Latitude  *float64 `form:"latitude,omitempty" json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *float64 `form:"longitude,omitempty" json:"longitude,omitempty" xml:"longitude,omitempty"`
	Number    *int     `form:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
}

// Validate validates the play type instance.
func (ut *play) Validate() (err error) {
	if ut.Latitude == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "latitude"))
	}
	if ut.Longitude == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "longitude"))
	}
	if ut.Number == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "number"))
	}

	return
}

// Publicize creates Play from play
func (ut *play) Publicize() *Play {
	var pub Play
	if ut.Latitude != nil {
		pub.Latitude = *ut.Latitude
	}
	if ut.Longitude != nil {
		pub.Longitude = *ut.Longitude
	}
	if ut.Number != nil {
		pub.Number = *ut.Number
	}
	return &pub
}

// Play user type.
type Play struct {
	Latitude  float64 `form:"latitude" json:"latitude" xml:"latitude"`
	Longitude float64 `form:"longitude" json:"longitude" xml:"longitude"`
	Number    int     `form:"number" json:"number" xml:"number"`
}
