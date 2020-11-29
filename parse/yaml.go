package parse

import (
	"gopkg.in/yaml.v2"
)

// Parse parses a yaml file into a set of rules
func Parse(data []byte) (map[string][]string, error) {
	var rules map[string][]string
	err := yaml.Unmarshal(data, &rules)
	if err != nil || len(rules) == 0 {
		return map[string][]string{}, err
	}

	return rules, nil
}
