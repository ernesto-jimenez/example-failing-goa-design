// Code generated by goa v3.7.10, DO NOT EDIT.
//
// FooService HTTP server types
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package server

import (
	fooservice "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/foo_service"
	goa "goa.design/goa/v3/pkg"
)

// FooMethodRequestBody is the type of the "FooService" service "FooMethod"
// endpoint HTTP request body.
type FooMethodRequestBody struct {
	External       *ExternalTypeRequestBody       `form:"External,omitempty" json:"External,omitempty" xml:"External,omitempty"`
	SecondExternal *SecondExternalTypeRequestBody `form:"SecondExternal,omitempty" json:"SecondExternal,omitempty" xml:"SecondExternal,omitempty"`
}

// FooMethodResponseBody is the type of the "FooService" service "FooMethod"
// endpoint HTTP response body.
type FooMethodResponseBody []*ExampleTypeResponse

// ExampleTypeResponse is used to define fields on response body types.
type ExampleTypeResponse struct {
	External       *ExternalTypeResponse       `form:"External,omitempty" json:"External,omitempty" xml:"External,omitempty"`
	SecondExternal *SecondExternalTypeResponse `form:"SecondExternal,omitempty" json:"SecondExternal,omitempty" xml:"SecondExternal,omitempty"`
}

// ExternalTypeResponse is used to define fields on response body types.
type ExternalTypeResponse struct {
	Field string `form:"Field" json:"Field" xml:"Field"`
}

// SecondExternalTypeResponse is used to define fields on response body types.
type SecondExternalTypeResponse struct {
	Field *string `form:"Field,omitempty" json:"Field,omitempty" xml:"Field,omitempty"`
}

// ExternalTypeRequestBody is used to define fields on request body types.
type ExternalTypeRequestBody struct {
	Field *string `form:"Field,omitempty" json:"Field,omitempty" xml:"Field,omitempty"`
}

// SecondExternalTypeRequestBody is used to define fields on request body types.
type SecondExternalTypeRequestBody struct {
	Field *string `form:"Field,omitempty" json:"Field,omitempty" xml:"Field,omitempty"`
}

// NewFooMethodResponseBody builds the HTTP response body from the result of
// the "FooMethod" endpoint of the "FooService" service.
func NewFooMethodResponseBody(res []*fooservice.ExampleType) FooMethodResponseBody {
	body := make([]*ExampleTypeResponse, len(res))
	for i, val := range res {
		body[i] = marshalFooserviceExampleTypeToExampleTypeResponse(val)
	}
	return body
}

// NewFooMethodPayload builds a FooService service FooMethod endpoint payload.
func NewFooMethodPayload(body *FooMethodRequestBody) *fooservice.FooMethodPayload {
	v := &fooservice.FooMethodPayload{}
	if body.External != nil {
		v.External = unmarshalExternalTypeRequestBodyToFooserviceExternalType(body.External)
	}
	if body.SecondExternal != nil {
		v.SecondExternal = unmarshalSecondExternalTypeRequestBodyToFooserviceSecondExternalType(body.SecondExternal)
	}

	return v
}

// ValidateFooMethodRequestBody runs the validations defined on
// FooMethodRequestBody
func ValidateFooMethodRequestBody(body *FooMethodRequestBody) (err error) {
	if body.External != nil {
		if err2 := ValidateExternalTypeRequestBody(body.External); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateExternalTypeRequestBody runs the validations defined on
// ExternalTypeRequestBody
func ValidateExternalTypeRequestBody(body *ExternalTypeRequestBody) (err error) {
	if body.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Field", "body"))
	}
	return
}
