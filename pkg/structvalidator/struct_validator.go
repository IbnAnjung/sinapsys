package structvalidator

import (
	"strings"
	"time"

	coreerror "github.com/IbnAnjung/synapsis/pkg/error"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(obj interface{}) error
}

type structValidator struct {
	validator *validator.Validate
}

func NewStructValidator() Validator {
	v := validator.New()

	v.RegisterValidation("date_format", DateFormatValidation)

	return &structValidator{
		validator: v,
	}
}

func (v *structValidator) Validate(obj interface{}) error {
	err := v.validator.Struct(obj)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "fail validate data")
			return e
		}

		e := coreerror.NewValidationError()
		e.Errors = map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			e.Errors[err.Field()] = err.Error()
		}

		err = e

		return err
	}

	return nil
}

func DateFormatValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	params := strings.Split(fl.Param(), ";")
	format := params[0]

	if _, err := time.Parse(format, value); err != nil {
		return false
	}

	return true
}
