package str

import (
	"regexp"
)

// Split splits a string with the given regular expression.
func Split(s, regex string) ([]string, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	return re.Split(s, -1), nil
}

// SplitSpace splits a string with whitespace characters.
func SplitSpace(s string) []string {
	// Here ignore the error
	parts, _ := Split(s, "\\s+")
	return parts
}

// IsEmpty checks whether the input string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}
