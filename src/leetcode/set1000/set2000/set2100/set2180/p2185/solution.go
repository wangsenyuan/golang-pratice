package p2185

import "strings"

func prefixCount(words []string, pref string) int {
	var res int

	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			res++
		}
	}

	return res
}
