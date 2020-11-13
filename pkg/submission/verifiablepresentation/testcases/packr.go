package testcases

import (
	"github.com/gobuffalo/packr"

	"go.wday.io/credentials-open-source/presentation-exchange/pkg/util"
)

var box = packr.NewBox(".")

const (
	SampleSubmission             util.JSONSchemaFile = "sample-submission"
	SampleVerifiablePresentation util.JSONSchemaFile = "sample-verifiable-presentation"
)

func GetJSONFile(name util.JSONSchemaFile) (string, error) {
	return util.GetJSONFile(&box, name)
}
