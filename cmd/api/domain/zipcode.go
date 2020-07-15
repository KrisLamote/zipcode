package domain

import (
	"errors"
)

var formats = map[string][]string{
	"BE": {"####"},
	"BR": {"#####-###", "#####"},
	"SK": {"## ###"},
}

const (
	zero = '0'
	nine = '9'
	hash = '#'
)

// Formats gives the known formats for the given country
func Formats(cntr string) ([]string, error) {
	f, found := formats[cntr]
	if !found {
		return []string{""}, errors.New("country not found")
	}
	return f, nil
}

// Valid checks whether a zipcode is valid for the given country
func Valid(zip string, cntr string) bool {
	fs, err := Formats(cntr)
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
