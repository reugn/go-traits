package traits

import "reflect"

type _default interface {
	initDefault(i interface{})
}

// Default trait for struct fields initialization
// Will work for exported fields only
type Default struct{}

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
			fv := reflect.New(ft.Type.Elem())
			initStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}
