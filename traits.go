package traits

import (
	"context"
	"log"
	"reflect"
	"runtime"
	"time"
)

// bootstrap is implemented by any value that has a Bootstrap method,
// which is executed on `traits.Init` call.
type bootstrap interface {
	Bootstrap()
}

// schedule is implemented by any value that has a Schedule method,
// which is scheduled in a separate go routine on `traits.Init` call.
// Close the context.Context to exit the scheduler goroutine.
type schedule interface {

	// Schedule returns the reccurent function, the `context.Context` to be used
	// for exiting scheduler goroutine and the schedule interval.
	Schedule() (func() error, context.Context, time.Duration)
}

// finalize is implemented by any value that has a Bootstrap method,
// which may run as soon as an object becomes unreachable.
type finalize interface {
	Finalize()
}

var traitsPackage = reflect.TypeOf(Stringer{}).PkgPath()

// Init is the mandatory function to call in order to initialize traits
// for the specified object.
func Init(obj interface{}) {
	if o, ok := obj.(converter); ok {
		o.setConverter(obj)
	}

	if o, ok := obj.(stringer); ok {
		o.setStringer(obj)
	}

	if o, ok := obj.(hasher); ok {
		o.setHasher(obj)
	}

	if o, ok := obj.(validator); ok {
		o.setValidator(obj)
	}

	if o, ok := obj.(_default); ok {
		o.initDefault(obj)
	}

	if o, ok := obj.(bootstrap); ok {
		o.Bootstrap()
	}

	if o, ok := obj.(schedule); ok {
		scheduled, ctx, interval := o.Schedule()

		go func() {
			select {
			case <-time.After(interval):
				if err := scheduled(); err != nil {
					log.Fatal(err)
				}

			case <-ctx.Done():
				log.Print("Exit a Schedule goroutine")
				return
			}
		}()
	}

	if _, ok := obj.(finalize); ok {
		runtime.SetFinalizer(obj, func(f finalize) {
			f.Finalize()
		})
	}
}
