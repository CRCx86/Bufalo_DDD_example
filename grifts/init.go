package grifts

import (
	"location_service_v1/ls_v2/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
