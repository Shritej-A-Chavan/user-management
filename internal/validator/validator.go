package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {
	err := Validate.Struct(data)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		switch fieldErr.Field() {
		case "Name":
			switch fieldErr.Tag() {
			case "required":
				errors["name"] = "name is required"
			case "min":
				errors["name"] = "name must be at least 2 characters"
			case "max":
				errors["name"] = "name cannot exceed 100 characters"
			}

		case "Dob":
			if fieldErr.Tag() == "required" {
				errors["dob"] = "date of birth is required"
			}
		}
	}

	return errors
}
