package valid

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate      *validator.Validate
	tableNameReg  = regexp.MustCompile("^[A-Z0-9_].{0,30}$")
	tableFieldRef = regexp.MustCompile("^[a-z0-9_].{0,30}$")
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateTableName(fl validator.FieldLevel) bool {
	res := tableNameReg.MatchString(fl.Field().String())
	return res
}

func ValidateField(fl validator.FieldLevel) bool {
	return tableFieldRef.MatchString(fl.Field().String())
}

func ValidateStruct(param any) []*ErrorResponse {
	var errorResponses []*ErrorResponse
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

func init() {
	var validateInstance = validator.New()
	err := validateInstance.RegisterValidation("tableName", ValidateTableName, false)
	if err != nil {
		panic(err)
	}
	err = validateInstance.RegisterValidation("tableField", ValidateField, false)
	if err != nil {
		panic(err)
	}
	validate = validateInstance
}
