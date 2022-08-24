package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ConvertErrToSliceOfString(err error) []string {
	var errValidator []string
	for _, e := range err.(validator.ValidationErrors) {
		errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
		errValidator = append(errValidator, errMessage)
	}
	return errValidator
}
