package arr

// FindStr searches the target string from left to right
// into the array and returns the left most result.
// It returns -1 if the value is not present.
func FindStr(s []string, target string) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}

	return -1
}
