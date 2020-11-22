package data

import (
	"fmt"
	"github.com/go-playground/validator"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	if v.Tag() == "required" {
		return fmt.Sprintf("%s is required to create a product instance", v.Field())
	} else {
		return fmt.Sprintf(
			"Key: %s Error: Field validation for %s failed on the %s tag",
			v.Namespace(),
			v.Field(),
			v.Tag(),
		)
	}
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate 
}

func NewValidation() *Validation {
	validate := validator.New()
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i)
	if errs == nil {
		return nil
	}

	var returnErrs ValidationErrors
	for _, err := range errs.(validator.ValidationErrors) {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}