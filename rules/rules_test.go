package rules

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	for _, test := range addCases {
		var rules map[string][]string
		if len(test.rules) > 0 {
			rules = test.rules
		} else {
			rules = map[string][]string{}
		}

		e := &Engine{rules: rules}
		err := e.Add(test.country, test.addRules)
		if test.expectErr {
			if err == nil {
				t.Errorf(
					"FAIL: '%s'\nAdd(%s, %s)\n  expected an error, received nil",
					test.desc,
					test.country,
					test.addRules,
				)
			}
			continue
		}

		if err != nil {
			t.Errorf(
				"FAIL: '%s'\nAdd(%s, %s)\n  expected no error, received error: %s",
				test.desc,
				test.country,
				test.addRules,
				err,
			)
			continue
		}

		formats, found := e.rules[test.country]
		if !found {
			t.Errorf(
				"FAIL: '%s'\nAdd(%s, %s)\n  expected country rules, but country was not found",
				test.desc,
				test.country,
				test.addRules,
			)
			continue
		}

		if !reflect.DeepEqual(test.expectRules, formats) {
			t.Errorf(
				"FAIL: '%s'\nAdd(%s, %s)\n  expected rules: %s\n  received rules: %s",
				test.desc,
				test.country,
				test.addRules,
				test.expectRules,
				formats,
			)
		}
	}
}

func TestDelete(t *testing.T) {
	for _, test := range delCases {
		var rules map[string][]string
		if len(test.rules) > 0 {
			rules = test.rules
		} else {
			rules = map[string][]string{}
		}

		e := &Engine{rules: rules}
		err := e.Delete(test.country)
		if test.expectErr {
			if err == nil {
				t.Errorf(
					"FAIL: '%s'\nDelete(%s)\n  expected an error, received nil",
					test.desc,
					test.country,
				)
			}
			continue
		}

		if err != nil {
			t.Errorf(
				"FAIL: '%s'\nDelete(%s)\n  expected no error, received error: %s",
				test.desc,
				test.country,
				err,
			)
			continue
		}

		_, found := e.rules[test.country]
		if found {
			t.Errorf(
				"FAIL: '%s'\nDelete(%s)\n  expected no country rules, but country found",
				test.desc,
				test.country,
			)
			continue
		}
	}
}
