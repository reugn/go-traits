package traits_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/reugn/go-traits"
)

type Inner struct {
	Arr []bool
}

type Test struct {
	traits.Hasher
	traits.Converter
	traits.Stringer
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
	obj := Test{
		Num: 1,
		Str: "abc",
		Inn: Inner{make([]bool, 2)},
	}
	traits.Init(&obj)

	assertEqual(t, obj.String(), `Test(1,abc,Inner([false false]))`)
	jsonStr, _ := obj.ToJSON()
	assertEqual(t, jsonStr, `{"num":1,"str":"abc","Inn":{"Arr":[false,false]}}`)

	assertEqual(t, obj.Md5Hex(), "a2122caf3c968cc7dd87f7783fe0abe5")
	assertEqual(t, obj.Sha256Hex(), "17a4d1650ec08d5466f6ded683c61b0becb597eaf3ba2f6a4f08dd44255ccb00")
	assertEqual(t, obj.HashCode32(), uint32(3850585125))

	obj.Num = 200000
	assertEqual(t, obj.Md5Hex(), "f42eaeaeb870fa47edbb8b236fa19d56")
	assertEqual(t, obj.Sha256Hex(), "ec8bb3fb4147a81049e99f00f64aa191ccebf77a0f17adff1320b7fa227021cf")
	assertEqual(t, obj.HashCode32(), uint32(2161010274))

	valid, _ := obj.Validate()
	assertEqual(t, valid, false)

	jsonStr = `{"num":1000,"str":"123","Inn":{"Arr":[false,true]}}`
	_ = json.Unmarshal([]byte(jsonStr), &obj)
	jsonStr, _ = obj.ToJSON()
	assertEqual(t, jsonStr, `{"num":1000,"str":"123","Inn":{"Arr":[false,true]}}`)
	valid, _ = obj.Validate()
	assertEqual(t, valid, true)

	assertEqual(t, obj.ToMap(), map[string]interface{}{
		"Num": 1000,
		"Str": "123",
		"Inn": obj.Inn,
	})

	assertEqual(t, obj.Keys(), []string{"Num", "Str", "Inn"})
	assertEqual(t, obj.Values(), []interface{}{1000, "123", obj.Inn})
}

type TestDefault struct {
	traits.Default
	i    int
	s    string
	M    map[string]int
	Ch   chan interface{}
	Sl   []string
	TPtr *Test
}

func TestTraitsDefault(t *testing.T) {
	obj := TestDefault{}
	traits.Init(&obj)

	assertEqual(t, obj.i, 0)
	assertEqual(t, obj.s, "")

	assertNotNil(t, obj.M)
	assertNotNil(t, obj.Ch)
	if obj.Sl == nil {
		t.Fatalf("%+v is nil", reflect.TypeOf(obj.Sl))
	}
	assertNotNil(t, obj.TPtr)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v != %v", a, b)
	}
}

func assertNonEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		t.Fatalf("%v == %v", a, b)
	}
}

func assertNotNil(t *testing.T, a interface{}) {
	if a == nil || reflect.ValueOf(a) == reflect.Zero(reflect.TypeOf(a)) {
		t.Fatalf("%+v is nil", reflect.TypeOf(a))
	}
}
