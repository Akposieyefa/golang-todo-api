package helpers

import "github.com/go-playground/validator/v10"

// validate struct
func Validate(request interface{}) error {
	validate := validator.New()
	return validate.Struct(request)
}
