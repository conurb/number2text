package number2text

import "testing"

func TestUnder100(t *testing.T) {
	tt := []struct {
		number int
		text   string
	}{
		{0, "zero"},
		{11, "eleven"},
		{23, "twenty-three"},
		{99, "ninety-nine"},
	}

	for _, tc := range tt {
		txt := under100(tc.number)
		if txt != tc.text {
			t.Errorf("expected %v; got %v", tc.text, txt)
		}
	}
}

func TestUnder1000(t *testing.T) {
	tt := []struct {
		number int
		text   string
	}{
		{0, "zero"},
		{11, "eleven"},
		{23, "twenty-three"},
		{99, "ninety-nine"},
		{100, "one hundred"},
		{200, "two hundred"},
		{222, "two hundred and twenty-two"},
		{909, "nine hundred and nine"},
	}

	for _, tc := range tt {
		txt := under1000(tc.number)
		if txt != tc.text {
			t.Errorf("expected %v; got %v", tc.text, txt)
		}
	}
}

func TestConvert(t *testing.T) {
	tt := []struct {
		sNumber string
		sText   string
	}{
		{"0", "zero"},
		{"11", "eleven"},
		{"23", "twenty-three"},
		{"99", "ninety-nine"},
		{"100", "one hundred"},
		{"200", "two hundred"},
		{"222", "two hundred and twenty-two"},
		{"909", "nine hundred and nine"},
		{"5000100000", "five billion one hundred thousand"},
		{"5000100001", "five billion one hundred thousand one"},
		{"5000000001", "five billion one"},
		{"5339430000", "five billion three hundred and thirty-nine million four hundred and thirty thousand"},
		{"11111111111111111111111111", "eleven septillion one hundred and eleven sextillion one hundred and eleven quintillion one hundred and eleven quadrillion one hundred and eleven trillion one hundred and eleven billion one hundred and eleven million one hundred and eleven thousand one hundred and eleven"},
		{"12q45p30", "one hundred and twenty-four thousand five hundred and thirty"},
		{"0111", "one hundred and eleven"},
		{"000 5 00 0 01", "five hundred thousand one"},
		{"1 000 500 001", "one billion five hundred thousand one"},
		{"", ""},
		{"   ", ""},
	}

	for _, tc := range tt {
		txt, _ := Convert(tc.sNumber)
		if txt != tc.sText {
			t.Errorf("expected %#v; got %#v", tc.sText, txt)
		}
	}
}

func TestError(t *testing.T) {
	_, err := Convert("111111111111111111111111111111111111111111111111111")
	if err.Error() != "number too large" {
		t.Errorf("expected %#v; got %#v", "number too large", err.Error())
	}
}
