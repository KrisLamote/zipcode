package domain

import (
	"reflect"
	"testing"
)

func TestFormats(t *testing.T) {
	for _, test := range formatsCases {
		formats, err := Formats(test.country)
		if !test.expectErr {
			// dont expect an error, but we got one
			if err != nil {
				t.Errorf("FAIL: '%s'\nFormat(%s)\n  expected no error\n  received error: %s", test.desc, test.country, err)
			}
			// format is wrong
			if !reflect.DeepEqual(formats, test.formats) {
				t.Errorf("FAIL: '%s'\nFormat(%s)\n  expected formats: %s\n  received formats: %s", test.desc, test.country, test.formats, formats)
			}
		} else if err == nil {
			t.Errorf("FAIL: '%s'\nFormat(%s)\n  expected an error, but error is nil", test.desc, test.country)
		}
		if !reflect.DeepEqual(formats, test.formats) {
			t.Errorf("FAIL: '%s'\nFormats(%s)\n  expected formats: %s\n  received formats: %s", test.desc, test.country, test.formats, formats)
		}
	}
}

func TestValidation(t *testing.T) {
	for _, test := range validationCases {
		actual := Valid(test.zipcode, test.country)
		if actual != test.expected {
			t.Errorf("FAIL: '%s'\nValid(%s, %s): %t\nexpected: %t", test.desc, test.country, test.zipcode, actual, test.expected)
		}
	}
}
