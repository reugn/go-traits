# go-traits
![Test](https://github.com/reugn/go-traits/workflows/Test/badge.svg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/reugn/go-traits)](https://pkg.go.dev/github.com/reugn/go-traits)
[![Go Report Card](https://goreportcard.com/badge/github.com/reugn/go-traits)](https://goreportcard.com/report/github.com/reugn/go-traits)

`go-traits` is a concept package that helps implement [mixin](https://en.wikipedia.org/wiki/Mixin) behavior using embedded structs and hook interfaces.

**Trait list:**  
* `traits.Hasher` - An extension for unique hash generators.
* `traits.Converter` - An extension for miscellaneous converters.
* `traits.Stringer` - `fmt.Stringer` implementation extension.
* `traits.Validator` - Struct fields validation extension.
* `traits.Default` - Struct fields initialization extension.

**Marker trait list:**
* `traits.PreventUnkeyed` - A struct to embed when you need to forbid unkeyed literals usage.
* `traits.NonComparable` - A struct to embed when you need to prevent structs comparison.

**Hook interfaces:**  
* `traits.bootstrap` - the `Bootstrap` function will be triggered on `traits.Init` call.
* `traits.schedule` - implement `traits.schedule` to schedule a function in a separate goroutine on `traits.Init` call.
* `traits.finalize` - the `Finalize` function will be set as an object finalizer via `runtime.SetFinalizer` on `traits.Init` call.

## Examples
```go
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
```
See the examples folder for more.

## Contributing
Any proposal or improvement is very welcome.

## License
Licensed under the MIT License.
