package validation

import (
	"testing"

	"github.com/KrisLamote/zipcode/rules"
)

func TestValidation(t *testing.T) {
	t.Parallel()
	e := rules.New(rules.Rules{
		"BE": []string{"####"},
		"BR": []string{"#####-###", "#####"},
		"SK": []string{"## ###"},
	})
	for _, test := range validationCases {
		actual := Valid(e, test.zipcode, test.country)
		if actual != test.expected {
			t.Errorf("FAIL: '%s'\nValid(%s, %s): %t\nexpected: %t", test.desc, test.country, test.zipcode, actual, test.expected)
		}
	}
}
