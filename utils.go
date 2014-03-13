package ninja

import (
	"regexp"
)

func FindStringSubmatchMap(re *regexp.Regexp, s string) map[string]string {
	captures := make(map[string]string)

	match := re.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range re.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}
		if match[i] != "" {
			captures[name] = match[i]
		}
	}
	return captures
}
