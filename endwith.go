package endswith

import "strings"

func EndsWith(s string, with string) bool {
	if strings.LastIndex(s, with)+len(with) == len(s) {
		return true
	}
	return false
}

func EndsWithIn(s string, within []string) bool {
	if len(within) == 0 {
		return false
	}

	for _, keyword := range within {
		if EndsWith(s, keyword) {
			return true
		}
	}

	return false
}
