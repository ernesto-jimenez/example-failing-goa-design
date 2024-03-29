// Code generated by goa v3.7.13, DO NOT EDIT.
//
// Example HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	fooservicec "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/http/foo_service/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `foo-service foo-method
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` foo-service foo-method --body '{
      "DateField": "2009-08-25T06:53:12Z",
      "External": {
         "Field": "Ut mollitia id similique."
      },
      "FieldWithExtension": {
         "BarField": {
            "Bar": 12728010016061705013
         }
      },
      "SecondExternal": {
         "Field": "Corrupti distinctio ut."
      }
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		fooServiceFlags = flag.NewFlagSet("foo-service", flag.ContinueOnError)

		fooServiceFooMethodFlags    = flag.NewFlagSet("foo-method", flag.ExitOnError)
		fooServiceFooMethodBodyFlag = fooServiceFooMethodFlags.String("body", "REQUIRED", "")
	)
	fooServiceFlags.Usage = fooServiceUsage
	fooServiceFooMethodFlags.Usage = fooServiceFooMethodUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "foo-service":
			svcf = fooServiceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "foo-service":
			switch epn {
			case "foo-method":
				epf = fooServiceFooMethodFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "foo-service":
			c := fooservicec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "foo-method":
				endpoint = c.FooMethod()
				data, err = fooservicec.BuildFooMethodPayload(*fooServiceFooMethodBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// foo-serviceUsage displays the usage of the foo-service command and its
// subcommands.
func fooServiceUsage() {
	fmt.Fprintf(os.Stderr, `Service is the FooService service interface.
Usage:
    %[1]s [globalflags] foo-service COMMAND [flags]

COMMAND:
    foo-method: FooMethod implements FooMethod.

Additional help:
    %[1]s foo-service COMMAND --help
`, os.Args[0])
}
func fooServiceFooMethodUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] foo-service foo-method -body JSON

FooMethod implements FooMethod.
    -body JSON: 

Example:
    %[1]s foo-service foo-method --body '{
      "DateField": "2009-08-25T06:53:12Z",
      "External": {
         "Field": "Ut mollitia id similique."
      },
      "FieldWithExtension": {
         "BarField": {
            "Bar": 12728010016061705013
         }
      },
      "SecondExternal": {
         "Field": "Corrupti distinctio ut."
      }
   }'
`, os.Args[0])
}
