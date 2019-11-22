package traits

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/reugn/go-traits/internal"
)

type stringify interface {
	ToString() string
	setStringify(i interface{})
}

// A Stringify trait extends struct with enhanced ToString representation method
type Stringify struct {
	self interface{}
}

func (str *Stringify) setStringify(i interface{}) {
	str.self = i
}

// ToString returns struct string representation
func (str *Stringify) ToString() string {
	v := reflect.ValueOf(str.self).Elem()
	t := reflect.TypeOf(str.self).Elem()
	return recursiveToString(v, t, "")
}

func recursiveToString(v reflect.Value, t reflect.Type, ptr string) string {
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Struct:
			//filter out traits package structs
			if v.Field(i).Type().PkgPath() == traitsPackage {
				values[i] = ""
			} else {
				values[i] = recursiveToString(v.Field(i), v.Field(i).Type(), "")
			}
		case reflect.Ptr:
			if v.Field(i).Elem().Kind() == reflect.Struct {
				//filter out traits package structs
				if v.Field(i).Type().PkgPath() == traitsPackage {
					values[i] = ""
				} else {
					values[i] = recursiveToString(v.Field(i).Elem(), v.Field(i).Elem().Type(), "&")
				}
			} else {
				values[i] = fmt.Sprintf("%+v", v.Field(i).Elem())
			}
		case reflect.Interface:
			values[i] = fmt.Sprintf("i[%+v]", v.Field(i))
		case reflect.Chan:
			values[i] = fmt.Sprintf("chan[%+v]", v.Field(i))
		default:
			values[i] = fmt.Sprintf("%+v", v.Field(i))
		}
	}
	values = internal.DeleteEmpty(values)
	return ptr + t.Name() + "(" + strings.Join(values, ",") + ")"
}
