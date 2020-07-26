package rules

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
