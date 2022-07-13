// Code generated by goa v3.7.10, DO NOT EDIT.
//
// FooService HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package client

import (
	"encoding/json"
	"fmt"

	fooservice "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/foo_service"
)

// BuildFooMethodPayload builds the payload for the FooService FooMethod
// endpoint from CLI flags.
func BuildFooMethodPayload(fooServiceFooMethodBody string) (*fooservice.FooMethodPayload, error) {
	var err error
	var body FooMethodRequestBody
	{
		err = json.Unmarshal([]byte(fooServiceFooMethodBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"External\": {\n         \"Field\": \"Eos quisquam et nisi molestiae.\"\n      }\n   }'")
		}
	}
	v := &fooservice.FooMethodPayload{}
	if body.External != nil {
		v.External = marshalTypesExternalTypeRequestBodyToTypesExternalType(body.External)
	}

	return v, nil
}
