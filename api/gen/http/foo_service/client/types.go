// Code generated by goa v3.7.10, DO NOT EDIT.
//
// FooService HTTP client types
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package client

import (
	fooservice "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/foo_service"
	goa "goa.design/goa/v3/pkg"
)

// FooMethodResponseBody is the type of the "FooService" service "FooMethod"
// endpoint HTTP response body.
type FooMethodResponseBody []*ExampleTypeResponse

// ExampleTypeRequestBody is used to define fields on request body types.
type ExampleTypeRequestBody struct {
	External *ExternalTypeRequestBody `form:"External,omitempty" json:"External,omitempty" xml:"External,omitempty"`
}

// ExternalTypeRequestBody is used to define fields on request body types.
type ExternalTypeRequestBody struct {
	Field string `form:"Field" json:"Field" xml:"Field"`
}

// ExampleTypeResponse is used to define fields on response body types.
type ExampleTypeResponse struct {
	External *ExternalTypeResponse `form:"External,omitempty" json:"External,omitempty" xml:"External,omitempty"`
}

// ExternalTypeResponse is used to define fields on response body types.
type ExternalTypeResponse struct {
	Field *string `form:"Field,omitempty" json:"Field,omitempty" xml:"Field,omitempty"`
}

// NewExampleTypeRequestBody builds the HTTP request body from the payload of
// the "FooMethod" endpoint of the "FooService" service.
func NewExampleTypeRequestBody(p []*fooservice.ExampleType) []*ExampleTypeRequestBody {
	body := make([]*ExampleTypeRequestBody, len(p))
	for i, val := range p {
		body[i] = marshalFooserviceExampleTypeToExampleTypeRequestBody(val)
	}
	return body
}

// NewFooMethodExampleTypeOK builds a "FooService" service "FooMethod" endpoint
// result from a HTTP "OK" response.
func NewFooMethodExampleTypeOK(body []*ExampleTypeResponse) []*fooservice.ExampleType {
	v := make([]*fooservice.ExampleType, len(body))
	for i, val := range body {
		v[i] = unmarshalExampleTypeResponseToFooserviceExampleType(val)
	}

	return v
}

// ValidateExampleTypeResponse runs the validations defined on
// ExampleTypeResponse
func ValidateExampleTypeResponse(body *ExampleTypeResponse) (err error) {
	if body.External != nil {
		if err2 := ValidateTypesExternalTypeResponse(body.External); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateExternalTypeResponse runs the validations defined on
// ExternalTypeResponse
func ValidateExternalTypeResponse(body *ExternalTypeResponse) (err error) {
	if body.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Field", "body"))
	}
	return
}
