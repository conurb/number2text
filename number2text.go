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

func dr(n int, div int) (d, r int) {
	d, r = n/div, n%div
	return
}

func under1000(n int) (st string) {
	d, r := dr(n, 100)
	if d == 0 {
		return under100(n)
	}
	st = a[d] + " hundred"
	if r > 0 {
		st += " and " + under100(r)
	}
	return
}

func under100(n int) (st string) {
	d, r := dr(n, 10)
	switch {
	case d < 2:
		st = a[n]
	case d < 10:
		st = b[d-2]
		if r > 0 {
			st += "-" + a[r]
		}
	}
	return
}

func string2slice(s string) []string {
	// get rid of possible non numeric characters
	re := regexp.MustCompile(`\d`)
	s = strings.Join(re.FindAllString(s, -1), "")

	d, r := dr(len(s), 3)

	var arr []string
	switch {
	case d == 0 && r == 0:
		return arr
	case r > 0:
		arr = append(arr, s[0:r])
		fallthrough
	case d > 0:
		re := regexp.MustCompile(`\d{3}`)
		arr = append(arr, re.FindAllString(s[r:], -1)...)
	}

	return arr
}

// Convert a number represented by a string to its textual representation
// eg: "123" gives "one hundred and twenty-three"
// usage of string to handle very big numbers (cf slice 'c')
func Convert(s string) (string, error) {
	s2s := string2slice(s)
	if len(s2s) > len(c) {
		return "", fmt.Errorf("number too large")
	}

	indexOfc := len(s2s)
	var output string
	for _, st := range s2s {
		s2i, err := strconv.Atoi(st)
		if err != nil {
			return "", fmt.Errorf("number could not be converted: %v", err)
		}
		indexOfc--
		if s2i == 0 {
			continue
		}
		output += fmt.Sprintf(" %s %s", under1000(s2i), c[indexOfc])
	}
	if output == "" && len(s2s) == 1 {
		output = "zero"
	}

	return strings.TrimSpace(output), nil
}
