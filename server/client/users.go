package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateUsersPayload is the Users Create action payload.
type CreateUsersPayload struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   string  `form:"name" json:"name" xml:"name"`
}

// CreateUsersPath computes a request path to the Create action of Users.
func CreateUsersPath() string {
	return fmt.Sprintf("/user")
}

// CreateUsers makes a request to the Create action endpoint of the Users resource
func (c *Client) CreateUsers(ctx context.Context, path string, payload *CreateUsersPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateUsersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateUsersRequest create the request corresponding to the Create action endpoint of the Users resource.
func (c *Client) NewCreateUsersRequest(ctx context.Context, path string, payload *CreateUsersPayload, contentType string) (*http.Request, error) {
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

// GetUsersPath computes a request path to the Get action of Users.
func GetUsersPath() string {
	return fmt.Sprintf("/user/me")
}

// GetUsers makes a request to the Get action endpoint of the Users resource
func (c *Client) GetUsers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUsersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUsersRequest create the request corresponding to the Get action endpoint of the Users resource.
func (c *Client) NewGetUsersRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateUsersPayload is the Users Update action payload.
type UpdateUsersPayload struct {
	Credit *int    `form:"credit,omitempty" json:"credit,omitempty" xml:"credit,omitempty"`
	ID     *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name   string  `form:"name" json:"name" xml:"name"`
}

// UpdateUsersPath computes a request path to the Update action of Users.
func UpdateUsersPath() string {
	return fmt.Sprintf("/user/me")
}

// UpdateUsers makes a request to the Update action endpoint of the Users resource
func (c *Client) UpdateUsers(ctx context.Context, path string, payload *UpdateUsersPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUsersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUsersRequest create the request corresponding to the Update action endpoint of the Users resource.
func (c *Client) NewUpdateUsersRequest(ctx context.Context, path string, payload *UpdateUsersPayload, contentType string) (*http.Request, error) {
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
