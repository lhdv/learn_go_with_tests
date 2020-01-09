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
