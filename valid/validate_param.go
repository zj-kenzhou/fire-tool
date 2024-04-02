package valid

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateTableName(fl validator.FieldLevel) bool {
	res, err := regexp.MatchString("^[A-Z0-9_].{0,30}$", fl.Field().String())
	if err != nil {
		return false
	}
	if !res {
		return false
	}
	return true
}

func ValidateField(fl validator.FieldLevel) bool {
	res, err := regexp.MatchString("^[a-z0-9_].{0,30}$", fl.Field().String())
	if err != nil {
		return false
	}
	if !res {
		return false
	}
	return true
}

func ValidateStruct(param any) []*ErrorResponse {
	var errorResponses []*ErrorResponse
	var validate = validator.New()
	err := validate.RegisterValidation("tableName", ValidateTableName, false)
	if err != nil {
		panic(err)
	}
	err = validate.RegisterValidation("tableField", ValidateField, false)
	if err != nil {
		panic(err)
	}
	err = validate.Struct(param)
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
