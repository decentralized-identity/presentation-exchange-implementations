package util

import (
	"encoding/json"
	"fmt"

	jcs "github.com/cyberphone/json-canonicalization/go/src/webpki.org/jsoncanonicalizer"
	"github.com/gobuffalo/packr"
	"github.com/hashicorp/go-multierror"
	"gopkg.in/go-playground/validator.v9"
)

type (
	JSONSchemaFile string
)

// Generic type validator
func Validate(t interface{}) error {
	if err := validator.New().Struct(t); err != nil {
		return err
	}
	return nil
}

func ToJSON(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	return string(bytes), err
}

// Take two JSON strings, put them into a go map so that they
// can be canonically marshalled using JCS. Then, compare the resulting strings.
func CompareJSON(j1, j2 string) (bool, error) {
	var err *multierror.Error
	j1bytes, j2bytes := []byte(j1), []byte(j2)
	jcs1, err1 := jcs.Transform(j1bytes)
	err = multierror.Append(err, err1)
	jcs2, err2 := jcs.Transform(j2bytes)
	err = multierror.Append(err, err2)
	return string(jcs1) == string(jcs2), err.ErrorOrNil()
}

func GetJSONFile(box *packr.Box, name JSONSchemaFile) (string, error) {
	return box.FindString(fmt.Sprintf("%s.json", name))
}
