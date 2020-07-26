package engine

import "testing"

func TestValidation(t *testing.T) {
	t.Parallel()
	e := New(Rules{
		"BE": []string{"####"},
		"BR": []string{"#####-###", "#####"},
		"SK": []string{"## ###"},
	})
	for _, test := range validationCases {
		actual := e.Valid(test.zipcode, test.country)
		if actual != test.expected {
			t.Errorf("FAIL: '%s'\nValid(%s, %s): %t\nexpected: %t", test.desc, test.country, test.zipcode, actual, test.expected)
		}
	}
}
