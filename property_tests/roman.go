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
func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range Numerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// ConvertToArabic convert a given Roman symbol to an Arabic number
func ConvertToArabic(roman string) int {
	total := 0

	for range roman {
		total++
	}

	return total
}
