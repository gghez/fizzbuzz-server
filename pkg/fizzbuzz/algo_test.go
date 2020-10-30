package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestCompute(t *testing.T) {
	cases := []struct {
		name              string
		int1, int2, limit int
		str1, str2        string
		expectedResults   []string
	}{
		{"limit 0", 3, 5, 0, "fizz", "buzz", []string{}},
		{"limit -2", 3, 5, -2, "fizz", "buzz", []string{}},
		{"limit 1", 3, 5, 1, "fizz", "buzz", []string{"1"}},
		{"limit 10", 3, 5, 10, "fizz", "buzz", []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			results := Compute(c.int1, c.int2, c.limit, c.str1, c.str2)
			if  !reflect.DeepEqual(results, c.expectedResults) {
				t.Fail()
			}
		})
	}
}
