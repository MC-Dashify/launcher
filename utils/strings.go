package utils

import "strings"

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str || strings.HasPrefix(v, str) {
			return true
		}
	}

	return false
}
