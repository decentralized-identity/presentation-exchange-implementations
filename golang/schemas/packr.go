package schemas

import (
	"github.com/gobuffalo/packr"

	"go.wday.io/credentials-open-source/presentation-exchange/pkg/util"
)

var box = packr.NewBox(".")

const (
	PresentationDefinition util.JSONSchemaFile = "presentation-definition"
	PresentationSubmission util.JSONSchemaFile = "presentation-submission"
)

func GetJSONFile(name util.JSONSchemaFile) (string, error) {
	return util.GetJSONFile(&box, name)
}
