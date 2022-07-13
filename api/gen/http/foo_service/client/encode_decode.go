// Code generated by goa v3.7.10, DO NOT EDIT.
//
// FooService HTTP client encoders and decoders
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	fooservice "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/foo_service"
	types "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/types"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildFooMethodRequest instantiates a HTTP request object with method and
// path set to call the "FooService" service "FooMethod" endpoint
func (c *Client) BuildFooMethodRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FooMethodFooServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("FooService", "FooMethod", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeFooMethodRequest returns an encoder for requests sent to the
// FooService FooMethod server.
func EncodeFooMethodRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.([]*fooservice.ExampleType)
		if !ok {
			return goahttp.ErrInvalidType("FooService", "FooMethod", "[]*fooservice.ExampleType", v)
		}
		body := NewExampleTypeRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("FooService", "FooMethod", err)
		}
		return nil
	}
}

// DecodeFooMethodResponse returns a decoder for responses returned by the
// FooService FooMethod endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeFooMethodResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FooMethodResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("FooService", "FooMethod", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateExampleTypeResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("FooService", "FooMethod", err)
			}
			res := NewFooMethodExampleTypeOK(body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("FooService", "FooMethod", resp.StatusCode, string(body))
		}
	}
}

// marshalFooserviceExampleTypeToExampleTypeRequestBody builds a value of type
// *ExampleTypeRequestBody from a value of type *fooservice.ExampleType.
func marshalFooserviceExampleTypeToExampleTypeRequestBody(v *fooservice.ExampleType) *ExampleTypeRequestBody {
	res := &ExampleTypeRequestBody{}
	if v.External != nil {
		res.External = marshalTypesExternalTypeToExternalTypeRequestBody(v.External)
	}

	return res
}

// marshalTypesExternalTypeToExternalTypeRequestBody builds a value of type
// *ExternalTypeRequestBody from a value of type *types.ExternalType.
func marshalTypesExternalTypeToExternalTypeRequestBody(v *types.ExternalType) *ExternalTypeRequestBody {
	if v == nil {
		return nil
	}
	res := &ExternalTypeRequestBody{
		Field: v.Field,
	}

	return res
}

// marshalExampleTypeRequestBodyToFooserviceExampleType builds a value of type
// *fooservice.ExampleType from a value of type *ExampleTypeRequestBody.
func marshalExampleTypeRequestBodyToFooserviceExampleType(v *ExampleTypeRequestBody) *fooservice.ExampleType {
	res := &fooservice.ExampleType{}
	if v.External != nil {
		res.External = marshalExternalTypeRequestBodyToTypesExternalType(v.External)
	}

	return res
}

// marshalExternalTypeRequestBodyToTypesExternalType builds a value of type
// *types.ExternalType from a value of type *ExternalTypeRequestBody.
func marshalExternalTypeRequestBodyToTypesExternalType(v *ExternalTypeRequestBody) *types.ExternalType {
	if v == nil {
		return nil
	}
	res := &types.ExternalType{
		Field: v.Field,
	}

	return res
}

// unmarshalExampleTypeResponseToFooserviceExampleType builds a value of type
// *fooservice.ExampleType from a value of type *ExampleTypeResponse.
func unmarshalExampleTypeResponseToFooserviceExampleType(v *ExampleTypeResponse) *fooservice.ExampleType {
	res := &fooservice.ExampleType{}
	if v.External != nil {
		res.External = unmarshalExternalTypeResponseToTypesExternalType(v.External)
	}

	return res
}

// unmarshalExternalTypeResponseToTypesExternalType builds a value of type
// *types.ExternalType from a value of type *ExternalTypeResponse.
func unmarshalExternalTypeResponseToTypesExternalType(v *ExternalTypeResponse) *types.ExternalType {
	if v == nil {
		return nil
	}
	res := &types.ExternalType{
		Field: *v.Field,
	}

	return res
}
