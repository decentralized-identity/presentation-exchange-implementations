package schemas

import (
	"github.com/gobuffalo/packr"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

var box = packr.NewBox(".")

const (
	PresentationDefinition util.JSONSchemaFile = "presentation-definition"
	PresentationSubmission util.JSONSchemaFile = "presentation-submission"
)

func GetJSONFile(name util.JSONSchemaFile) (string, error) {
	return util.GetJSONFile(&box, name)
}
