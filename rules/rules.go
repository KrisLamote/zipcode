package rules

import (
	"errors"
	"regexp"
)

var reg = regexp.MustCompile("[^a-zA-Z]+")

// Rules maps a country code to a slice of formatting rules
type Rules = map[string][]string

// Engine holds the rules
type Engine struct {
	rules Rules
}

// Add format rules to the given country. If the country doesn't exist, add it.
// Cntr needs to be 2 string, preferably representing an iso country code.
func (e *Engine) Add(cntr string, formats []string) error {
	if len(cntr) != 2 {
		return errors.New("the length of country code should be exactly 2")
	}

	c := reg.ReplaceAllString(cntr, "")
	if len(cntr) != len(c) {
		return errors.New("the country code should alphabetic ASCII characters only")
	}

	f, found := e.rules[cntr]
	if !found {
		f = make([]string, 0)
	}

	e.rules[cntr] = append(f, formats...)
	return nil
}
