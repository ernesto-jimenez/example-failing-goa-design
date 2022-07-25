// Code generated by goa v3.7.12, DO NOT EDIT.
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
	FieldWithExtension *FieldWithExtensionRequestBody `form:"FieldWithExtension,omitempty" json:"FieldWithExtension,omitempty" xml:"FieldWithExtension,omitempty"`
	External           *ExternalTypeRequestBody       `form:"External,omitempty" json:"External,omitempty" xml:"External,omitempty"`
	SecondExternal     *SecondExternalTypeRequestBody `form:"SecondExternal,omitempty" json:"SecondExternal,omitempty" xml:"SecondExternal,omitempty"`
	DateField          *string                        `form:"DateField,omitempty" json:"DateField,omitempty" xml:"DateField,omitempty"`
}

// FooMethodResponseBody is the type of the "FooService" service "FooMethod"
// endpoint HTTP response body.
type FooMethodResponseBody []*ExampleTypeResponse

// ExampleTypeResponse is used to define fields on response body types.
type ExampleTypeResponse struct {
	FieldWithExtension *FieldWithExtensionResponse `form:"FieldWithExtension,omitempty" json:"FieldWithExtension,omitempty" xml:"FieldWithExtension,omitempty"`
	External           *ExternalTypeResponse       `form:"External,omitempty" json:"External,omitempty" xml:"External,omitempty"`
	SecondExternal     *SecondExternalTypeResponse `form:"SecondExternal,omitempty" json:"SecondExternal,omitempty" xml:"SecondExternal,omitempty"`
	DateField          string                      `form:"DateField" json:"DateField" xml:"DateField"`
}

// FieldWithExtensionResponse is used to define fields on response body types.
type FieldWithExtensionResponse struct {
	BarField *BarResponse `form:"BarField,omitempty" json:"BarField,omitempty" xml:"BarField,omitempty"`
}

// BarResponse is used to define fields on response body types.
type BarResponse struct {
	Bar uint `form:"Bar" json:"Bar" xml:"Bar"`
}

// ExternalTypeResponse is used to define fields on response body types.
type ExternalTypeResponse struct {
	Field string `form:"Field" json:"Field" xml:"Field"`
}

// SecondExternalTypeResponse is used to define fields on response body types.
type SecondExternalTypeResponse struct {
	Field *string `form:"Field,omitempty" json:"Field,omitempty" xml:"Field,omitempty"`
}

// FieldWithExtensionRequestBody is used to define fields on request body types.
type FieldWithExtensionRequestBody struct {
	BarField *Bar `form:"BarField,omitempty" json:"BarField,omitempty" xml:"BarField,omitempty"`
}

// Bar is used to define fields on request body types.
type Bar struct {
	Bar *uint `form:"Bar,omitempty" json:"Bar,omitempty" xml:"Bar,omitempty"`
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
	v := &fooservice.FooMethodPayload{
		DateField: fooservice.DateTime(*body.DateField),
	}
	if body.FieldWithExtension != nil {
		v.FieldWithExtension = unmarshalFieldWithExtensionRequestBodyToFooserviceFieldWithExtension(body.FieldWithExtension)
	}
	if body.External != nil {
		v.External = unmarshalExternalTypeRequestBodyToTypesExternalType(body.External)
	}
	if body.SecondExternal != nil {
		v.SecondExternal = unmarshalSecondExternalTypeRequestBodyToTypesSecondExternalType(body.SecondExternal)
	}

	return v
}

// ValidateFooMethodRequestBody runs the validations defined on
// FooMethodRequestBody
func ValidateFooMethodRequestBody(body *FooMethodRequestBody) (err error) {
	if body.DateField == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("DateField", "body"))
	}
	if body.FieldWithExtension != nil {
		if err2 := ValidateFieldWithExtensionRequestBody(body.FieldWithExtension); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.External != nil {
		if err2 := ValidateExternalTypeRequestBody(body.External); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.DateField != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.DateField", *body.DateField, goa.FormatDateTime))
	}
	return
}

// ValidateFieldWithExtensionRequestBody runs the validations defined on
// FieldWithExtensionRequestBody
func ValidateFieldWithExtensionRequestBody(body *FieldWithExtensionRequestBody) (err error) {
	if body.BarField != nil {
		if err2 := ValidateBar(body.BarField); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateBar runs the validations defined on Bar
func ValidateBar(body *Bar) (err error) {
	if body.Bar == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Bar", "body"))
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
