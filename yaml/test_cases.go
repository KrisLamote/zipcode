package yaml

import "github.com/KrisLamote/zipcode/engine"

type parseTest struct {
	desc  string
	yaml  []byte
	rules engine.Rules
}

var parseCases = []parseTest{
	{
		desc:  "empty file",
		yaml:  []byte(""),
		rules: map[string][]string{},
	},
	{
		desc:  "Belgium",
		yaml:  []byte("BE:\n- \"####\""),
		rules: map[string][]string{"BE": {"####"}},
	},
	{
		desc:  "Brasil",
		yaml:  []byte("BR:\n- \"#####-###\"\n- \"#####\""),
		rules: map[string][]string{"BR": {"#####-###", "#####"}},
	},
	{
		desc:  "Multiple countries with multiple rules",
		yaml:  []byte("BE:\n- \"####\"\nBR:\n- \"#####-###\"\n- \"#####\""),
		rules: map[string][]string{"BE": {"####"}, "BR": {"#####-###", "#####"}},
	},
}

type errorTest struct {
	desc      string
	yaml      []byte
	yamlError bool
	ownError  string
}

// we are not going to test Yaml unmarshaling, simply if we pass those errors through
// TODO: on top of these we may need extra errors (use ownError string for these?)
var errorCases = []errorTest{
	{
		desc:      "Yalm error",
		yaml:      []byte("just a string"),
		yamlError: true,
		ownError:  "",
	},
	{
		desc:      "doesnt fit in the Rules type",
		yaml:      []byte("A: [B: [2]]"),
		yamlError: true,
		ownError:  "",
	},
}
