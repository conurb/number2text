package number2text

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	a = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	b = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	c = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion",
		"septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion",
		"quattuordecillion"}
)

func dr(n, div int) (d, r int) {
	d, r = n/div, n%div
	return
}

func under1000(n int) (s string) {
	d, r := dr(n, 100)
	if d == 0 {
		return under100(n)
	}
	s = a[d] + " hundred"
	if r > 0 {
		s += " and " + under100(r)
	}
	return
}

func under100(n int) (s string) {
	d, r := dr(n, 10)
	switch {
	case d < 2:
		s = a[n]
	case d < 10:
		s = b[d-2]
		if r > 0 {
			s += "-" + a[r]
		}
	}
	return
}

func groupBy3(s string) []string {
	// get rid of possible non numeric characters
	re := regexp.MustCompile(`\d`)
	s = strings.Join(re.FindAllString(s, -1), "")

	d, r := dr(len(s), 3)

	var arr []string
	switch {
	case r > 0:
		arr = append(arr, s[0:r])
		fallthrough
	case d > 0:
		re := regexp.MustCompile(`\d{3}`)
		arr = append(arr, re.FindAllString(s[r:], -1)...)
	}

	return arr
}

// Convert a number represented by a string to its textual representation.
// Input type is string to handle big numbers (until quattuordecillion).
//
// eg: Convert("123") will output "one hundred and twenty-three".
func Convert(s string) (string, error) {
	aBy3 := groupBy3(s)
	if len(aBy3) > len(c) {
		return "", fmt.Errorf("number too large")
	}

	indexOfc := len(aBy3)
	var output string
	for _, st := range aBy3 {
		i2convert, err := strconv.Atoi(st)
		if err != nil {
			return "", fmt.Errorf("number could not be converted: %v", err)
		}
		indexOfc--
		if i2convert == 0 {
			continue
		}
		output += fmt.Sprintf(" %s %s", under1000(i2convert), c[indexOfc])
	}

	if output == "" && len(aBy3) == 1 {
		output = "zero"
	}

	return strings.TrimSpace(output), nil
}
