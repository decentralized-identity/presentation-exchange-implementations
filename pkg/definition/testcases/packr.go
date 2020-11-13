package testcases

import (
	"github.com/gobuffalo/packr"

	"github.com/decentralized-identity/presentation-exchange-implementations/pkg/util"
)

var box = packr.NewBox(".")

const (
	BasicPresentationDefinition       util.JSONSchemaFile = "basic-presentation-definition"
	SingleGroupPresentationDefinition util.JSONSchemaFile = "single-group-presentation-definition"
	MultiGroupPresentationDefinition  util.JSONSchemaFile = "multi-group-presentation-definition"
)

func GetJSONFile(name util.JSONSchemaFile) (string, error) {
	return util.GetJSONFile(&box, name)
}
