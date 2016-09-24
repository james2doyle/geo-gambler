package design

import (
    . "github.com/goadesign/goa/design"
    . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("GeoGambler", func() {
    Version("v1")
    Host("localhost:8000")
    Scheme("http")
})