package traits

import "reflect"

type _default interface {
	initDefault(i interface{})
}

// Default adds default initialization to an embedding struct fields.
// Note that it applies to exported fields only.
// No public methods exposed for this trait.
type Default struct{}

// initDefault is called from the `traits.Init` function.
func (d *Default) initDefault(i interface{}) {
	v := reflect.ValueOf(i).Elem()
	t := reflect.TypeOf(i).Elem()

	initStruct(t, v)
}

func initStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))

		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))

		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))

		case reflect.Struct:
			initStruct(ft.Type, f)

		case reflect.Ptr:
			ptr := reflect.New(ft.Type.Elem())
			initStruct(ft.Type.Elem(), ptr.Elem())
			f.Set(ptr)

		default:
		}
	}
}
