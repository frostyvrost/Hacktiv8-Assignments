package pkg

import "github.com/asaskevich/govalidator"

func ValidateStruct(payload interface{}) MessageErr {

	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return NewBadRequest(err.Error())
	}

	return nil
}
