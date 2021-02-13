package traits

import "github.com/asaskevich/govalidator"

type validator interface {
	Validate() (bool, error)

	setValidator(i interface{})
}

// Validator adds the Validate function to an embedding struct,
// which can be used to validate structure fields by tags.
//
// List of available validators for struct fields:
// "email"
// "url"
// "dialstring"
// "requrl"
// "requri"
// "alpha"
// "utfletter"
// "alphanum"
// "utfletternum"
// "numeric"
// "utfnumeric"
// "utfdigit"
// "hexadecimal"
// "hexcolor"
// "rgbcolor"
// "lowercase"
// "uppercase"
// "int"
// "float"
// "null"
// "uuid"
// "uuidv3"
// "uuidv4"
// "uuidv5"
// "creditcard"
// "isbn10"
// "isbn13"
// "json"
// "multibyte"
// "ascii"
// "printableascii"
// "fullwidth"
// "halfwidth"
// "variablewidth"
// "base64"
// "datauri"
// "ip"
// "port"
// "ipv4"
// "ipv6"
// "dns"
// "host"
// "mac"
// "latitude"
// "longitude"
// "ssn"
// "semver"
// "rfc3339"
// "rfc3339WithoutZone"
// "ISO3166Alpha2"
// "ISO3166Alpha3"
//
// Example:
//
// type Example struct {
// 		Str   string `valid:"numeric"`
//		Email string `valid:"email"`
// }
//
type Validator struct {
	self interface{}
}

var _ validator = (*Validator)(nil)

func (v *Validator) setValidator(i interface{}) {
	v.self = i
}

// Validate uses tags to validate struct fields.
// The result will be equal to `false` if there are any errors.
func (v *Validator) Validate() (bool, error) {
	return govalidator.ValidateStruct(v.self)
}
