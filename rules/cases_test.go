package rules

type AddTest struct {
	desc        string
	country     string
	rules       []string
	addRules    []string
	expectRules []string
	expectErr   bool
}

var addCases = []AddTest{
	{
		desc:        "Add to an empty engine should insert country and rules",
		country:     "BE",
		rules:       []string{},
		addRules:    []string{"####"},
		expectRules: []string{"####"},
		expectErr:   false,
	},
	{
		desc:        "Append to existing country rules",
		country:     "BR",
		rules:       []string{"#####-###"},
		addRules:    []string{"#####"},
		expectRules: []string{"#####-###", "#####"},
		expectErr:   false,
	},
	{
		desc:        "Country codes should be length 2",
		country:     "ABC",
		rules:       []string{},
		addRules:    []string{},
		expectRules: []string{},
		expectErr:   true,
	},
	{
		desc:        "Country codes should be ASCII",
		country:     "Å",
		rules:       []string{},
		addRules:    []string{},
		expectRules: []string{},
		expectErr:   true,
	},
}
