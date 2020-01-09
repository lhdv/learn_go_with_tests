package roman

import (
	"testing"
)

func TestConvertToRoman(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to I", 2, "II"},
		{"3 gets converted to I", 3, "III"},
		{"4 gets converted to I", 4, "IV"},
		{"5 gets converted to I", 5, "V"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			want := test.Want
			got := ConvertToRoman(test.Arabic)
			if want != got {
				t.Errorf("want: %q, got: %q", want, got)
			}
		})
	}

}
