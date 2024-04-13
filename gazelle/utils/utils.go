package js

import (
	"strings"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Remove(s *[]string, r string) {
	for i, v := range *s {
		if v == r {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}

func Any(s []string, pred func(string) bool) bool {
	for _, v := range s {
		if pred(v) {
			return true
		}
	}

	return false
}

func MergeUnique(a []string, b []string) []string {
	unique := map[string]bool{}
	values := append(a, b...)
	result := []string{}

	for _, value := range values {
		if unique[value] {
			continue
		}

		unique[value] = true
		result = append(result, value)
	}

	return result
}

func HasPrefixes(prefixes []string, x string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(x, prefix) {
			return true
		}
	}
	return false
}

func HasSuffixes(suffixes []string, x string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(x, suffix) {
			return true
		}
	}
	return false
}

func SlicesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aMap := make(map[string]bool, len(b))
	for i := range a {
		aMap[a[i]] = true
	}
	for i := range b {
		if !aMap[b[i]] {
			return false
		}
	}
	return true
}

func SliceFromStringSet(input map[string]bool) []string {
	result := []string{}
	for s := range input {
		result = append(result, s)
	}

	return result
}
