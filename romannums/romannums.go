package romannums

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for arabic > 0 {
		switch {
		case arabic > 9: // 10, ...
			result.WriteString("X")
			arabic -= 10
		case arabic > 8: // 9, 10, ...
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4: // 5, 6, ...
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}
	return result.String()
}

// TODO: exercise claims that string builder is more efficient than
// plain string concat operation.
func ConvertToRomanNoStringBuilder(arabic int) string {
	result := ""
	for i := arabic; i > 0; i-- {
		if arabic == 5 {
			result += "V"
			break
		}
		if arabic == 4 {
			result += "IV"
			break
		}
		result += "I"
	}
	return result
}
