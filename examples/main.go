package main

import (
	"fmt"

	"github.com/reugn/go-traits"
)

type inner struct {
	Arr []bool
}

type test struct {
	traits.Hasher
	traits.Converter
	traits.Stringer
	traits.Validator

	Num   int    `json:"num"`
	Str   string `json:"str" valid:"numeric"`
	Inn   *inner
	pstr  *string
	C     chan interface{} `json:"-"`
	Iface interface{}
}

func (t *test) Bootstrap() {
	fmt.Println("Bootstrap Test struct...")
}

func (t *test) Finalize() {
	fmt.Println("Finalize Test struct...")
}

func main() {
	str := "bar"
	obj := test{
		Num:   1,
		Str:   "abc",
		Inn:   &inner{make([]bool, 2)},
		pstr:  &str,
		C:     make(chan interface{}),
		Iface: "foo",
	}
	traits.Init(&obj)

	fmt.Println(obj.String())
	fmt.Println(obj.ToJSON())
	fmt.Println(obj.Md5Hex())
	fmt.Println(obj.Sha256Hex())
	fmt.Println(obj.HashCode32())
	fmt.Println(obj.Validate())
}
