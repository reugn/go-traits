package traits

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"reflect"
	"runtime"
	"strings"

	"github.com/asaskevich/govalidator"
)

type finalize interface {
	Finalize()
}

type bootstrap interface {
	Bootstrap()
}

type jsonify interface {
	ToJSON() (string, error)
	setJsonify(i interface{})
}

type stringify interface {
	ToString() string
	setStringify(i interface{})
}

type hash interface {
	Md5() [16]byte
	Md5Hex() string
	Sha256() [32]byte
	Sha256Hex() string
	HashCode() uint32
	setHasher(i interface{})
}

type validator interface {
	Validate() (bool, error)
	setValidator(i interface{})
}

var traitsPackage = reflect.TypeOf(Jsonify{}).PkgPath()

// Jsonify trait provides ToJson capability
type Jsonify struct {
	self interface{}
}

func (js *Jsonify) setJsonify(i interface{}) {
	js.self = i
}

// ToJSON serializes struct to json format
func (js *Jsonify) ToJSON() (string, error) {
	b, err := json.Marshal(js.self)
	return string(b), err
}

// Stringify trait provides ToString capability
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
	values = deleteEmpty(values)
	return ptr + t.Name() + "(" + strings.Join(values, ",") + ")"
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Hash trait provides unique hash generators
type Hash struct {
	self interface{}
}

func (h *Hash) setHasher(i interface{}) {
	h.self = i
}

// Md5 byte array for self struct
func (h *Hash) Md5() [16]byte {
	jsonBytes, _ := json.Marshal(h.self)
	return md5.Sum(jsonBytes)
}

// Md5Hex string for self struct
func (h *Hash) Md5Hex() string {
	md5Bytes := h.Md5()
	return hex.EncodeToString(md5Bytes[:])
}

// Sha256 byte array for self struct
func (h *Hash) Sha256() [32]byte {
	jsonBytes, _ := json.Marshal(h.self)
	return sha256.Sum256(jsonBytes)
}

// Sha256Hex string for self struct
func (h *Hash) Sha256Hex() string {
	sha256Bytes := h.Sha256()
	return hex.EncodeToString(sha256Bytes[:])
}

// HashCode uint32 for self struct
func (h *Hash) HashCode() uint32 {
	jsonBytes, _ := json.Marshal(h.self)
	h32 := fnv.New32a()
	h32.Write(jsonBytes)
	return h32.Sum32()
}

// Validator trait provides struct fields validation capability
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

//Init traits capabilities for given object
func Init(obj interface{}) error {
	if o, ok := obj.(jsonify); ok {
		o.setJsonify(obj)
	}
	if o, ok := obj.(stringify); ok {
		o.setStringify(obj)
	}
	if o, ok := obj.(hash); ok {
		o.setHasher(obj)
	}
	if o, ok := obj.(validator); ok {
		o.setValidator(obj)
	}
	if o, ok := obj.(bootstrap); ok {
		o.Bootstrap()
	}
	if _, ok := obj.(finalize); ok {
		var fin = func(f finalize) {
			f.Finalize()
		}
		runtime.SetFinalizer(obj, fin)
	}
	return nil
}
