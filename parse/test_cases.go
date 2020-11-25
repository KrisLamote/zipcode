package parse

import "github.com/KrisLamote/zipcode/engine"

type parseTest struct {
	desc  string
	yaml  string
	rules engine.Rules
}

var parseCases = []parseTest{
	{
		desc:  "empty file",
		yaml:  "",
		rules: map[string][]string{},
	},
	{
		desc:  "Belgium",
		yaml:  "BE:\n- \"####\"",
		rules: map[string][]string{"BE": {"####"}},
	},
	{
		desc:  "Brasil",
		yaml:  "BR:\n- \"#####-###\"\n- \"#####\"",
		rules: map[string][]string{"BR": {"#####-###", "#####"}},
	},
	{
		desc:  "Multiple countries with multiple rules",
		yaml:  "BE:\n- \"####\"\nBR:\n- \"#####-###\"\n- \"#####\"",
		rules: map[string][]string{"BE": {"####"}, "BR": {"#####-###", "#####"}},
	},
}

type errorTest struct {
	desc      string
	yaml      string
	yamlError bool
	ownError  string
}

// we are not going to test Yaml unmarshaling, simply if we pass those errors through
// TODO: on top of these we may need extra errors (use ownError string for these?)
var errorCases = []errorTest{
	{
		desc:      "Yalm error",
		yaml:      "just a string",
		yamlError: true,
		ownError:  "",
	},
	{
		desc:      "doesnt fit in the Rules type",
		yaml:      "A: [B: [2]]",
		yamlError: true,
		ownError:  "",
	},
}
