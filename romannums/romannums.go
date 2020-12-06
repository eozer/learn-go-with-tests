package romannums

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
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
	for _, romannum := range allRomanNumerals {
		for arabic >= romannum.Value {
			result.WriteString(romannum.Symbol)
			arabic -= romannum.Value
		}
	}
	return result.String()
}

func couldBeSubtractive(i int, symbol byte, roman string) bool {
	isSubtractiveSymbol := symbol == 'I' || symbol == 'X' || symbol == 'C'
	return i+1 < len(roman) && isSubtractiveSymbol
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i] // this is byte/uint8 type.
		fmt.Print(symbol, " ")
		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]
			if value := allRomanNumerals.ValueOf(symbol, nextSymbol); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
		fmt.Println("total: ", total)
	}

	return total
}

// NOTE: exercise claims that string builder is more efficient than
// plain string concat operation. For this exercise I did not observe
// a clear winner. Difference was small and sometimes NoStringBuilder
// approach won against the builder way.
func ConvertToRomanNoStringBuilder(arabic int) string {
	result := ""
	for _, romannum := range allRomanNumerals {
		for arabic >= romannum.Value {
			result += romannum.Symbol
			arabic -= romannum.Value
		}
	}
	return result
}
