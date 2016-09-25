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

package app

import "github.com/goadesign/goa"

// play user type.
type play struct {
	Bet    *int `form:"bet,omitempty" json:"bet,omitempty" xml:"bet,omitempty"`
	Number *int `form:"number,omitempty" json:"number,omitempty" xml:"number,omitempty"`
}

// Validate validates the play type instance.
func (ut *play) Validate() (err error) {
	if ut.Number == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "number"))
	}
	if ut.Bet == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "bet"))
	}

	if ut.Bet != nil {
		if *ut.Bet < 10 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.bet`, *ut.Bet, 10, true))
		}
	}
	if ut.Number != nil {
		if *ut.Number < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.number`, *ut.Number, 0, true))
		}
	}
	if ut.Number != nil {
		if *ut.Number > 9 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.number`, *ut.Number, 9, false))
		}
	}
	return
}

// Publicize creates Play from play
func (ut *play) Publicize() *Play {
	var pub Play
	if ut.Bet != nil {
		pub.Bet = *ut.Bet
	}
	if ut.Number != nil {
		pub.Number = *ut.Number
	}
	return &pub
}

// Play user type.
type Play struct {
	Bet    int `form:"bet" json:"bet" xml:"bet"`
	Number int `form:"number" json:"number" xml:"number"`
}

// Validate validates the Play type instance.
func (ut *Play) Validate() (err error) {
	if ut.Bet < 10 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.bet`, ut.Bet, 10, true))
	}
	if ut.Number < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.number`, ut.Number, 0, true))
	}
	if ut.Number > 9 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.number`, ut.Number, 9, false))
	}
	return
}
