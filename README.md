# go-traits
Extend Go struct capabilities

## About
There are various approaches to reduce Go boilerplate like go generate and others. Traits library comes to add capabilities and reduce boilerplate using embedded structs.

Traits list:  
* traits.Hash - struct unique hash generators
* traits.Jsonify - marshal struct to JSON
* traits.Stringify - stringify struct
* traits.Validator - validate struct fields

Hooks interfaces:  
* traits.bootstrap - triggered on traits.Init()
* traits.finalize - sets object finalizer via ```runtime.SetFinalizer(ptr, finalizerFunc)```

## Examples
```go
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

func main() {
	obj := Test{Num: 1, Str: "abc", Inn: Inner{make([]bool, 2)}}
	traits.Init(&obj)

	fmt.Println(obj.ToString())
	fmt.Println(obj.ToJSON())
	fmt.Println(obj.Md5Hex())
	fmt.Println(obj.Sha256Hex())
	fmt.Println(obj.Validate())
}
```
See examples folder for more.

## Contributing
Feel free to add more capabilities

## License
Licensed under the MIT License.