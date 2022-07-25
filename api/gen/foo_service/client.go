// Code generated by goa v3.7.13, DO NOT EDIT.
//
// FooService client
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package fooservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "FooService" service client.
type Client struct {
	FooMethodEndpoint goa.Endpoint
}

// NewClient initializes a "FooService" service client given the endpoints.
func NewClient(fooMethod goa.Endpoint) *Client {
	return &Client{
		FooMethodEndpoint: fooMethod,
	}
}

// FooMethod calls the "FooMethod" endpoint of the "FooService" service.
func (c *Client) FooMethod(ctx context.Context, p *FooMethodPayload) (res []*ExampleType, err error) {
	var ires interface{}
	ires, err = c.FooMethodEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*ExampleType), nil
}
