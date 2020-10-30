package fizzbuzz

import (
	"fmt"
)

// Compute solves parametrized fizzbuzz
func Compute(int1, int2, limit int, str1, str2 string) []string {
	if limit < 1 {
		limit = 0
	}
	strs := make([]string, 0, limit)
	var s string
	for i := 1; i <= limit; i++ {
		s = ""
		if i%int1 == 0 {
			s = str1
		}
		if i%int2 == 0 {
			s += str2
		}
		if len(s) == 0 {
			s = fmt.Sprint(i)
		}
		strs = append(strs, s)
	}
	return strs
}
