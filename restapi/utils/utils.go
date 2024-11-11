package utils

import (
	"fmt"
	"sort"
)

func SprintMapStringInt(m map[string]int) string {
	var s string
	for k, v := range m {
		s += fmt.Sprintf("%s:\t%d\n", k, v)
	}
	return s
}

func ParseURL(urlArray []string) []string {
	res := make([]string, 0)
	for _, v := range urlArray {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func ReversreString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(s)-1-i] = runes[len(s)-1-i], runes[i]
	}
	return string(runes)
}

func DeduplicateString(s string) string {
	m := make(map[rune]bool)
	var runes []rune

	for _, r := range s {
		if m[r] {
			continue
		}
		runes = append(runes, r)
		m[r] = true
	}
	
	return string(runes)
}

func SortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}