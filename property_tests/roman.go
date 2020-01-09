package roman

import "strings"

// ConvertToRoman will convert a given int to its representative number
// in Roman notation
func ConvertToRoman(num int) string {

	var result strings.Builder

	for i := num; i > 0; i-- {

		if i == 4 {
			result.WriteString("IV")
			break
		}

		if i == 5 {
			result.WriteString("V")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
