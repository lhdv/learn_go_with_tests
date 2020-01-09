package roman

import "strings"

// Numeral represents a numeral in Roman with its value in arabic number
type Numeral struct {
	Value  int
	Symbol string
}

var Numerals = []Numeral{
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ConvertToRoman will convert a given int to its representative number
// in Roman notation
func ConvertToRoman(num int) string {

	var result strings.Builder

	for _, numeral := range Numerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}
