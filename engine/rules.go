package engine

import (
	"errors"
	"fmt"
	"regexp"
)

var reg = regexp.MustCompile("[^a-zA-Z]+")

// Rules maps a country code to a slice of formatting rules
type Rules = map[string][]string

// Engine holds the rules
type Engine struct {
	rules Rules
}

// Empty intialises a new empty Engine
func Empty() *Engine {
	e := &Engine{rules: map[string][]string{}}
	return e
}

// New intialises a new Engine
func New(r Rules) *Engine {
	e := &Engine{rules: r}
	return e
}

// Formats gives the known formats for the given country
func (e *Engine) Formats(cntr string) ([]string, error) {
	f, found := e.rules[cntr]
	if !found {
		return []string{""}, errors.New("country not found")
	}
	return f, nil
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

// Delete all format rules to the given country.
func (e *Engine) Delete(cntr string) error {
	_, found := e.rules[cntr]
	if !found {
		return fmt.Errorf("the country code not found (%v)", cntr)
	}
	delete(e.rules, cntr)
	return nil
}

// Valid checks whether a zipcode is valid for the given country
func (e *Engine) Valid(zip string, cntr string) bool {
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

const (
	zero = '0'
	nine = '9'
	hash = '#'
)

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
