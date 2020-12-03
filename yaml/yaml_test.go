package yaml

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	for _, test := range parseCases {
		actual, err := Parse(test.yaml)
		if err != nil {
			t.Errorf("FAIL: '%s'\nParse()\n  expected no error\n  received error: %s", test.desc, err)
			continue
		}
		if !reflect.DeepEqual(actual, test.rules) {
			t.Errorf("FAIL: '%s'\nParse()\n  expected rules: %s\n  received rules: %s", test.desc, test.rules, actual)
			continue
		}
	}
}

func TestErrors(t *testing.T) {
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

func TestRealFile(t *testing.T) {
	desc := "read real file"
	path := "../data/rules.yml"
	// reads the entire file (and does defer close)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("FAIL: could not read rules file %s, %s", path, err)
	}

	actual, err := Parse(data)
	if err != nil {
		t.Errorf("FAIL: '%s'\nParse()\n  expected no error\n  received error: %s", desc, err)
	}

	if len(actual) != 131 {
		t.Errorf("FAIL: '%s'\nParse()\n  expected 131 country keys\n  received: %d", desc, len(actual))
	}

	singapore, found := actual["SG"]
	if !found {
		t.Errorf("FAIL: '%s'\nParse()\n  missing expected rules for Singaore (SG)", desc)
	}

	if len(singapore) != 3 {
		t.Errorf("FAIL: '%s'\nParse()\n  expected 3 rules for Singapore, received %d", desc, len(singapore))
	}
}
