package romannums

import "testing"

var testcases = []struct {
	Description string
	Arabic      int
	Want        string
}{
	{"1 gets converted to I", 1, "I"},
	{"2 gets converted to II", 2, "II"},
	{"3 gets converted to III", 3, "III"},
	{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
	{"5 gets converted to V", 5, "V"},
	{"8 gets converted to VIII", 8, "VIII"},
	{"9 gets converted to IX", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
}

func BenchmarkConvertToRoman(t *testing.B) {
	for _, test := range testcases {
		ConvertToRoman(test.Arabic)
	}
}

func BenchmarkConvertToRomanNoStringBuilder(t *testing.B) {
	for _, test := range testcases {
		ConvertToRomanNoStringBuilder(test.Arabic)
	}
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range testcases {
		t.Run(test.Description, func(t *testing.T) {
			t.Helper()
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
