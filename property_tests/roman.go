package roman

import "strings"

// Numeral represents a numeral in Roman with its value in arabic number
type Numeral struct {
	Value  int
	Symbol string
}

// Numerals variable contains all the Roman's numerals signals and values
var Numerals = []Numeral{
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
