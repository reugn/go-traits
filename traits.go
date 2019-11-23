package traits

import (
	"reflect"
	"runtime"
)

type finalize interface {
	Finalize()
}

type bootstrap interface {
	Bootstrap()
}

var traitsPackage = reflect.TypeOf(Jsonify{}).PkgPath()

// Init traits capabilities for given object
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
	if o, ok := obj.(_default); ok {
		o.initDefault(obj)
	}
	if _, ok := obj.(finalize); ok {
		var fin = func(f finalize) {
			f.Finalize()
		}
		runtime.SetFinalizer(obj, fin)
	}
	return nil
}
