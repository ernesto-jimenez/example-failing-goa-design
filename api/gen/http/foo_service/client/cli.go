// Code generated by goa v3.7.13, DO NOT EDIT.
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
	goa "goa.design/goa/v3/pkg"
)

// BuildFooMethodPayload builds the payload for the FooService FooMethod
// endpoint from CLI flags.
func BuildFooMethodPayload(fooServiceFooMethodBody string) (*fooservice.FooMethodPayload, error) {
	var err error
	var body FooMethodRequestBody
	{
		err = json.Unmarshal([]byte(fooServiceFooMethodBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"DateField\": \"2009-08-25T06:53:12Z\",\n      \"External\": {\n         \"Field\": \"Ut mollitia id similique.\"\n      },\n      \"FieldWithExtension\": {\n         \"BarField\": {\n            \"Bar\": 12728010016061705013\n         }\n      },\n      \"SecondExternal\": {\n         \"Field\": \"Corrupti distinctio ut.\"\n      }\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.DateField", body.DateField, goa.FormatDateTime))

		if err != nil {
			return nil, err
		}
	}
	v := &fooservice.FooMethodPayload{
		DateField: fooservice.DateTime(body.DateField),
	}
	if body.FieldWithExtension != nil {
		v.FieldWithExtension = marshalFieldWithExtensionRequestBodyToFooserviceFieldWithExtension(body.FieldWithExtension)
	}
	if body.External != nil {
		v.External = marshalExternalTypeRequestBodyToTypesExternalType(body.External)
	}
	if body.SecondExternal != nil {
		v.SecondExternal = marshalSecondExternalTypeRequestBodyToTypesSecondExternalType(body.SecondExternal)
	}

	return v, nil
}
