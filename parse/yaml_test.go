package parse

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()
	for _, test := range parseCases {
		actual, err := Parse(test.yaml)
		if err != nil {
			t.Errorf("FAIL: '%s'\nParse()\n  expected no error\n  received error: %s", test.desc, err)
			continue
		}
		if !reflect.DeepEqual(actual, test.rules) {
			t.Errorf("FAIL: '%s'\nParse()\n  expected rules: %s\n  received rules: %s", test.desc, test.rules, actual)
		}
	}

	for _, test := range errorCases {
		actual, err := Parse(test.yaml)
		if err == nil {
			t.Errorf("FAIL: '%s'\nParse()\n  expected an error\n  received: %v", test.desc, actual)
			continue
		}

		if test.yamlError {
			// yaml will start with "yaml:"
			received := err.Error()[0:5]
			if received != "yaml:" {
				t.Errorf("FAIL: '%s'\nParse()\n  expected a Yaml error\n  received: %s", test.desc, err)
			}
			continue
		}

		// TODO: check for specific errors
	}
}
