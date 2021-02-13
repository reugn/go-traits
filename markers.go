package traits

// PreventUnkeyed is a struct to embed when you need
// to forbid unkeyed literals usage.
//
// PreventUnkeyed trait doesn't require the `traits.Init` call.
//
type PreventUnkeyed struct {
	_ struct{}
}

// NonComparable is a struct to embed when you need
// to prevent structs comparison.
//
// NonComparable trait doesn't require the `traits.Init` call.
//
type NonComparable struct {
	_ [0]func()
}
