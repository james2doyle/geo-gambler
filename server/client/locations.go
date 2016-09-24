package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// FindAllLocationsPath computes a request path to the FindAll action of Locations.
func FindAllLocationsPath() string {
	return fmt.Sprintf("/locations")
}

// FindAllLocations makes a request to the FindAll action endpoint of the Locations resource
func (c *Client) FindAllLocations(ctx context.Context, path string, game *string) (*http.Response, error) {
	req, err := c.NewFindAllLocationsRequest(ctx, path, game)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFindAllLocationsRequest create the request corresponding to the FindAll action endpoint of the Locations resource.
func (c *Client) NewFindAllLocationsRequest(ctx context.Context, path string, game *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if game != nil {
		values.Set("game", *game)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetLocationsPath computes a request path to the Get action of Locations.
func GetLocationsPath(id int) string {
	return fmt.Sprintf("/locations/%v", id)
}

// GetLocations makes a request to the Get action endpoint of the Locations resource
func (c *Client) GetLocations(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetLocationsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetLocationsRequest create the request corresponding to the Get action endpoint of the Locations resource.
func (c *Client) NewGetLocationsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PlayLocationsPath computes a request path to the Play action of Locations.
func PlayLocationsPath(id int) string {
	return fmt.Sprintf("/locations/play/%v", id)
}

// PlayLocations makes a request to the Play action endpoint of the Locations resource
func (c *Client) PlayLocations(ctx context.Context, path string, payload *Play, contentType string) (*http.Response, error) {
	req, err := c.NewPlayLocationsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPlayLocationsRequest create the request corresponding to the Play action endpoint of the Locations resource.
func (c *Client) NewPlayLocationsRequest(ctx context.Context, path string, payload *Play, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
