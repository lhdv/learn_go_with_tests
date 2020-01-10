package roman

import "strings"

// Numeral represents a numeral in Roman with its value in arabic number
type Numeral struct {
	Value  int
	Symbol string
}

// Numerals is an array of Roman numerals
type Numerals []Numeral

// ValueOf return the int value of a given symbol
func (r Numerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

// Numerals variable contains all the Roman's numerals signals and values
var numerals = Numerals{
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

	for _, numeral := range numerals {
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

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {
			if value := numerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += numerals.ValueOf(symbol)
			}
		} else {
			total += numerals.ValueOf(symbol)
		}
	}

	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
