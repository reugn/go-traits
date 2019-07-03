package traits_test

import (
	"encoding/json"
	"fmt"
	"testing"

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
	Num int    `json:"num"`
	Str string `json:"str" valid:"numeric"`
	Inn Inner
}

func (t *Test) Bootstrap() {
	fmt.Println("Bootstrap Test struct...")
}

func (t *Test) Finalize() {
	fmt.Println("Finalize Test struct...")
}

func TestTraits(t *testing.T) {
	obj := Test{Num: 1, Str: "abc", Inn: Inner{make([]bool, 2)}}
	traits.Init(&obj)

	assertEqual(t, obj.ToString(), `Test(1,abc,Inner([false false]))`)
	jsonStr, _ := obj.ToJSON()
	assertEqual(t, jsonStr, `{"num":1,"str":"abc","Inn":{"Arr":[false,false]}}`)

	assertEqual(t, obj.Md5Hex(), "a2122caf3c968cc7dd87f7783fe0abe5")
	assertEqual(t, obj.Sha256Hex(), "17a4d1650ec08d5466f6ded683c61b0becb597eaf3ba2f6a4f08dd44255ccb00")
	assertEqual(t, obj.HashCode(), uint32(3850585125))

	obj.Num = 200000
	assertEqual(t, obj.Md5Hex(), "f42eaeaeb870fa47edbb8b236fa19d56")
	assertEqual(t, obj.Sha256Hex(), "ec8bb3fb4147a81049e99f00f64aa191ccebf77a0f17adff1320b7fa227021cf")
	assertEqual(t, obj.HashCode(), uint32(2161010274))

	valid, _ := obj.Validate()
	assertEqual(t, valid, false)

	jsonStr = `{"num":1000,"str":"123","Inn":{"Arr":[false,true]}}`
	json.Unmarshal([]byte(jsonStr), &obj)
	jsonStr, _ = obj.ToJSON()
	assertEqual(t, jsonStr, `{"num":1000,"str":"123","Inn":{"Arr":[false,true]}}`)
	valid, _ = obj.Validate()
	assertEqual(t, valid, true)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
