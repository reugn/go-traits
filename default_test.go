package traits_test

import (
	"testing"

	"github.com/reugn/go-traits"
)

type inner struct {
	I int
}

type defs struct {
	traits.Default

	Chan   chan bool
	Slice  []int
	Map    map[string]int
	Inn    inner
	InnPtr *inner
}

func TestDefaultTrait(t *testing.T) {
	df := defs{}
	traits.Init(&df)

	assertNotNil(t, df.Chan)
	assertNotNil(t, df.Slice)
	assertNotNil(t, df.Map)
	assertNotNil(t, df.Inn)
	assertNotNil(t, df.InnPtr)

	assertEqual(t, df.Inn.I, 0)
	assertEqual(t, df.InnPtr.I, 0)
}
