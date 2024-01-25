package valid

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(param any) []*ErrorResponse {
	var errorResponses []*ErrorResponse
	var validate = validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errorResponses = append(errorResponses, &element)
		}
	}
	return errorResponses
}

func Validate(param any) error {
	errs := ValidateStruct(param)
	if errs != nil {
		res, err := json.Marshal(errs)
		if err != nil {
			return err
		}
		return errors.New(string(res))
	}
	return nil
}
