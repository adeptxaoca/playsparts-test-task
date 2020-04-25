package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	VendorCodeRegexp = regexp.MustCompile(`^[a-zA-Z\d\-_]+$`)
	NameRegexp       = regexp.MustCompile(`^[a-zA-Z\d\s]+$`)
)

type Validator struct {
	Validate *validator.Validate
}

func New() *Validator {
	validate := validator.New()
	_ = validate.RegisterValidation("name", validateName)
	_ = validate.RegisterValidation("vendor-code", validateVendorCode)
	return &Validator{Validate: validate}
}

func validateName(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return NameRegexp.MatchString(str)
}

func validateVendorCode(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return VendorCodeRegexp.MatchString(str)
}
