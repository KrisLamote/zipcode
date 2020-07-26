package validation

import (
	"github.com/KrisLamote/zipcode/rules"
)

const (
	zero = '0'
	nine = '9'
	hash = '#'
)

// Valid checks whether a zipcode is valid for the given country
func Valid(e *rules.Engine, zip string, cntr string) bool {
	fs, err := e.Formats(cntr)
	if err != nil {
		return false
	}

	for _, format := range fs {
		if isValid(format, zip) {
			return true
		}
	}

	return false
}

func isValid(format string, zip string) bool {
	// check format length
	if len(zip) != len(format) {
		return false
	}

	// walk through format & zipcode checking each rune rule
	for i := 0; i < len(format); i = i + 1 {
		switch format[i] {
		case hash:
			if zip[i] < zero || nine < zip[i] {
				return false
			}
		// if it's something else, then it simply needs to be the same
		default:
			if format[i] != zip[i] {
				return false
			}
		}
	}

	return true
}
