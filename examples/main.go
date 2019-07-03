package main

import (
	"fmt"

	"github.com/reugn/go-traits"
)

type Inner struct {
	Arr []bool
}

type Test struct {
	traits.Hash
	traits.Jsonify
	traits.Stringify
	traits.Validator
	Num   int    `json:"num"`
	Str   string `json:"str" valid:"numeric"`
	Inn   *Inner
	pstr  *string
	C     chan interface{} `json:"-"`
	Iface interface{}
}

func (t *Test) Bootstrap() {
	fmt.Println("Bootstrap Test struct...")
}

func (t *Test) Finalize() {
	fmt.Println("Finalize Test struct...")
}

func main() {
	str := "bar"
	obj := Test{Num: 1, Str: "abc", Inn: &Inner{make([]bool, 2)},
		pstr: &str, C: make(chan interface{}), Iface: "foo"}
	traits.Init(&obj)

	fmt.Println(obj.ToString())
	fmt.Println(obj.ToJSON())
	fmt.Println(obj.Md5Hex())
	fmt.Println(obj.Sha256Hex())
	fmt.Println(obj.Validate())
}
