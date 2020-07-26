package engine

type formatsTest struct {
	desc      string
	country   string
	formats   []string
	expectErr bool
}

var formatsCases = []formatsTest{
	{
		desc:      "BE format",
		country:   "BE",
		formats:   []string{"####"},
		expectErr: false,
	},
	{
		desc:      "SK format",
		country:   "SK",
		formats:   []string{"## ###"},
		expectErr: false,
	},
	{
		desc:      "Unknown country code",
		country:   "XX",
		formats:   []string{""},
		expectErr: true,
	},
}

type AddTest struct {
	desc        string
	rules       Rules
	country     string
	addRules    []string
	expectRules []string
	expectErr   bool
}

var addCases = []AddTest{
	{
		desc:        "Add to an empty engine should insert country and rules",
		rules:       Rules{},
		country:     "BE",
		addRules:    []string{"####"},
		expectRules: []string{"####"},
		expectErr:   false,
	},
	{
		desc:        "Append to existing country rules",
		rules:       Rules{"BR": {"#####-###"}},
		country:     "BR",
		addRules:    []string{"#####"},
		expectRules: []string{"#####-###", "#####"},
		expectErr:   false,
	},
	{
		desc:        "Country codes should be length 2",
		rules:       Rules{},
		country:     "ABC",
		addRules:    []string{},
		expectRules: []string{},
		expectErr:   true,
	},
	{
		desc:        "Country codes should be ASCII",
		rules:       Rules{},
		country:     "Å",
		addRules:    []string{},
		expectRules: []string{},
		expectErr:   true,
	},
}

type DeleteTest struct {
	desc      string
	rules     Rules
	country   string
	expectErr bool
}

var delCases = []DeleteTest{
	{
		desc:      "Delete existing country",
		rules:     Rules{"BE": {"####"}},
		country:   "BE",
		expectErr: false,
	},
	{
		desc:      "Delete non-existing country",
		rules:     Rules{"SK": {"## ###"}},
		country:   "BE",
		expectErr: true,
	},
}

type validationTest struct {
	desc     string
	zipcode  string
	country  string
	expected bool
}

var validationCases = []validationTest{
	{
		desc:     "Valid BE zipcode",
		zipcode:  "3000",
		country:  "BE",
		expected: true,
	},
	{
		desc:     "Valid SK zipcode",
		zipcode:  "83 242",
		country:  "SK",
		expected: true,
	},
	{
		desc:     "Valid BR zipcode, type1",
		zipcode:  "12345-123",
		country:  "BR",
		expected: true,
	},
	{
		desc:     "Valid BR zipcode, type2",
		zipcode:  "12345",
		country:  "BR",
		expected: true,
	},
	{
		desc:     "Invalid BR zipcode, too short",
		zipcode:  "1234",
		country:  "BR",
		expected: false,
	},
	{
		desc:     "Invalid BR zipcode, wrong separator",
		zipcode:  "12345_123",
		country:  "BR",
		expected: false,
	},
	{
		desc:     "Invalid BR zipcode, too long",
		zipcode:  "12345-1234",
		country:  "BR",
		expected: false,
	},
	{
		desc:     "Zipcode length < format length",
		zipcode:  "300",
		country:  "BE",
		expected: false,
	},
	{
		desc:     "Missing space",
		zipcode:  "83242",
		country:  "SK",
		expected: false,
	},
	{
		desc:     "Only digits",
		zipcode:  "abcd",
		country:  "BE",
		expected: false,
	},
}
