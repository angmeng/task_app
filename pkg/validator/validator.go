package validator

import (
	"github.com/asaskevich/govalidator"
)

func ValidateStruct(params interface{}) error {
	govalidator.TagMap["alpha"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, "^[a-zA-Z\\s]+$")
	})

	govalidator.TagMap["alphanum"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, "^[a-zA-Z0-9\\s]+$")
	})

	if _, err := govalidator.ValidateStruct(params); err != nil {
		return err
	}

	return nil
}
