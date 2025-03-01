package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"

	"part_handler/internal/pkg/errors"
)

var (
	VendorCodeRegexp = regexp.MustCompile(`^[a-zA-Z\d\-_]+$`)
	NameRegexp       = regexp.MustCompile(`^[a-zA-Z\d\s]+$`)
)

type Validator struct {
	Validate *validator.Validate
}

// New returns a new validator with instance of 'validate'
func New() *Validator {
	validate := validator.New()
	_ = validate.RegisterValidation("name", validateName)
	_ = validate.RegisterValidation("vendor-code", validateVendorCode)
	return &Validator{Validate: validate}
}

// Struct validates a structs exposed fields
func (v *Validator) Struct(s interface{}) error {
	if err := v.Validate.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fields := make([]string, 0, len(validationErrors))
		for _, err := range err.(validator.ValidationErrors) {
			fields = append(fields, err.Namespace())
		}

		return errors.ValidationError.New(strings.Join(fields, ","))
	}
	return nil
}

func validateName(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return NameRegexp.MatchString(str)
}

func validateVendorCode(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return VendorCodeRegexp.MatchString(str)
}
