// Code generated by goa v3.7.10, DO NOT EDIT.
//
// FooService HTTP server encoders and decoders
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package server

import (
	"context"
	"io"
	"net/http"

	fooservice "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/foo_service"
	types "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/types"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeFooMethodResponse returns an encoder for responses returned by the
// FooService FooMethod endpoint.
func EncodeFooMethodResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*fooservice.ExampleType)
		enc := encoder(ctx, w)
		body := NewFooMethodResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeFooMethodRequest returns a decoder for requests sent to the FooService
// FooMethod endpoint.
func DecodeFooMethodRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body FooMethodRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateFooMethodRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewFooMethodExampleType(&body)

		return payload, nil
	}
}

// unmarshalExternalTypeRequestBodyToTypesExternalType builds a value of type
// *types.ExternalType from a value of type *ExternalTypeRequestBody.
func unmarshalExternalTypeRequestBodyToTypesExternalType(v *ExternalTypeRequestBody) *types.ExternalType {
	if v == nil {
		return nil
	}
	res := &types.ExternalType{
		Field: *v.Field,
	}

	return res
}

// marshalFooserviceExampleTypeToExampleTypeResponse builds a value of type
// *ExampleTypeResponse from a value of type *fooservice.ExampleType.
func marshalFooserviceExampleTypeToExampleTypeResponse(v *fooservice.ExampleType) *ExampleTypeResponse {
	res := &ExampleTypeResponse{}
	if v.External != nil {
		res.External = marshalTypesExternalTypeToTypesExternalTypeResponse(v.External)
	}

	return res
}

// marshalTypesExternalTypeToTypesExternalTypeResponse builds a value of type
// *types.ExternalTypeResponse from a value of type *types.ExternalType.
func marshalTypesExternalTypeToTypesExternalTypeResponse(v *types.ExternalType) *types.ExternalTypeResponse {
	if v == nil {
		return nil
	}
	res := &types.ExternalTypeResponse{
		Field: v.Field,
	}

	return res
}