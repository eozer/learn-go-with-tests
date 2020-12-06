package romannums

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, romannum := range romanNumerals {
		for arabic >= romannum.Value {
			result.WriteString(romannum.Symbol)
			arabic -= romannum.Value
		}
	}

	return result.String()
}

// NOTE: exercise claims that string builder is more efficient than
// plain string concat operation. For this exercise I did not observe
// a clear winner. Difference was small and sometimes NoStringBuilder
// approach won against the builder way.
func ConvertToRomanNoStringBuilder(arabic int) string {
	result := ""
	for _, romannum := range romanNumerals {
		for arabic >= romannum.Value {
			result += romannum.Symbol
			arabic -= romannum.Value
		}
	}
	return result
}
