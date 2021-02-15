package internal

// DeleteEmpty clears empty strings from the specified slice.
// Returns a new slice without modifying the original one.
func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// Check the error and panic if not nil.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
