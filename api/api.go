package api

import (
	types "github.com/ernesto-jimenez/example-failing-goa-design/api/gen/types"
)

//go:generate goa gen github.com/ernesto-jimenez/example-failing-goa-design/design/example -o .

var Example types.ExternalType
