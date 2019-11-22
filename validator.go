package traits

import "github.com/asaskevich/govalidator"

type validator interface {
	Validate() (bool, error)
	setValidator(i interface{})
}

// A Validator trait adds fields validation capability
type Validator struct {
	self interface{}
}

func (v *Validator) setValidator(i interface{}) {
	v.self = i
}

// Validate validates struct fields
func (v *Validator) Validate() (bool, error) {
	return govalidator.ValidateStruct(v.self)
}
