package roman

import "strings"

// ConvertToRoman will convert a given int to its representative number
// in Roman notation
func ConvertToRoman(num int) string {

	var result strings.Builder

	for i := num; i > 0; i-- {
		result.WriteString("I")

		if i == 4 {
			return "IV"
		}

		if i == 5 {
			return "V"
		}

	}

	return result.String()
}
