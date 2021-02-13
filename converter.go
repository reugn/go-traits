package traits

import (
	"encoding/json"
	"encoding/xml"
	"reflect"

	yaml "gopkg.in/yaml.v3"
)

type converter interface {
	ToJSON() (string, error)
	ToJSONIndent(prefix, indent string) (string, error)
	ToYAML() (string, error)
	ToXML() (string, error)
	ToXMLIndent(prefix, indent string) (string, error)

	ToMap() map[string]interface{}
	Keys() []string
	Values() []interface{}

	setConverter(i interface{})
}

// Converter provides various marshalling methods to an embedding struct.
type Converter struct {
	self interface{}
}

var _ converter = (*Converter)(nil)

func (conv *Converter) setConverter(i interface{}) {
	conv.self = i
}

// ToJSON returns the JSON encoding of an embedding struct.
func (conv *Converter) ToJSON() (string, error) {
	b, err := json.Marshal(conv.self)
	return string(b), err
}

// ToJSONIndent is like ToJSON but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
func (conv *Converter) ToJSONIndent(prefix, indent string) (string, error) {
	b, err := json.MarshalIndent(conv.self, prefix, indent)
	return string(b), err
}

// ToYAML serializes an embedding struct into a YAML document.
func (conv *Converter) ToYAML() (string, error) {
	b, err := yaml.Marshal(conv.self)
	return string(b), err
}

// ToXML returns the XML encoding of an embedding struct.
func (conv *Converter) ToXML() (string, error) {
	b, err := xml.Marshal(conv.self)
	return string(b), err
}

// ToXMLIndent works like ToXML, but each XML element begins on a new
// indented line that starts with prefix and is followed by one or more
// copies of indent according to the nesting depth.
func (conv *Converter) ToXMLIndent(prefix, indent string) (string, error) {
	b, err := xml.MarshalIndent(conv.self, prefix, indent)
	return string(b), err
}

// ToMap converts an embedding struct to a map.
// Skips unexported fields.
func (conv *Converter) ToMap() map[string]interface{} {
	v := reflect.Indirect(reflect.ValueOf(conv.self))
	fieldMap := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() && v.Field(i).Type().PkgPath() != traitsPackage {
			fieldMap[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
	}
	return fieldMap
}

// Keys returns a slice of exported field names of an embedding struct
func (conv *Converter) Keys() []string {
	v := reflect.Indirect(reflect.ValueOf(conv.self))
	keys := make([]string, 0)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() && v.Field(i).Type().PkgPath() != traitsPackage {
			keys = append(keys, v.Type().Field(i).Name)
		}
	}
	return keys
}

// Values returns a slice of exported field values of an embedding struct
func (conv *Converter) Values() []interface{} {
	v := reflect.Indirect(reflect.ValueOf(conv.self))
	values := make([]interface{}, 0)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() && v.Field(i).Type().PkgPath() != traitsPackage {
			values = append(values, v.Field(i).Interface())
		}
	}
	return values
}
