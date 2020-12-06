package romannums

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

var testcases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
	{2888, "MMDCCCLXXXVIII"},
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
		description := fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman)
		t.Run(description, func(t *testing.T) {
			got, _ := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range testcases {
		description := fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic)
		t.Run(description, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %q, want %q", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		log.Println(arabic)
		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	cfg := quick.Config{
		MaxCount: 1000,
	}
	if err := quick.Check(assertion, &cfg); err != nil {
		t.Error("failed checks", err)
	}
}
