package validator

import (
	"fmt"
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

// TODO: Validate errors handler
func Errors(err error) {
	// this check is only needed when your code could produce
	// an invalid value for validation such as interface with nil
	// value most including myself do not usually have code like this.
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
		return
	}

	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace())
		fmt.Println(err.StructField())
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println(err.Param())
		fmt.Println()
	}
}

func validateName(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return NameRegexp.MatchString(str)
}

func validateVendorCode(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return VendorCodeRegexp.MatchString(str)
}
