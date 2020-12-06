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

func (r RomanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
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

// ConvertToRoman converts arabic integer to a Roman numeral in strings.
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

func SplitSymbols(w string) (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractiveSymbol(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractiveSymbol(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

func ConvertToArabic(roman string) (total int) {
	splits := SplitSymbols(roman)
	fmt.Printf("%s\n", splits)
	for _, symbols := range splits {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}
