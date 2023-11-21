package pkg

import (
	"github.com/asaskevich/govalidator"
)

func ValidateStruct(s interface{}) Error {

	_, err := govalidator.ValidateStruct(s)

	if err != nil {
		return NewBadRequestError(err.Error())
	}

	return nil
}
