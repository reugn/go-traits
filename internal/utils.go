package internal

// DeleteEmpty clears empty strings from a given slice
// returns a new slice without modifying the original one
func DeleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
