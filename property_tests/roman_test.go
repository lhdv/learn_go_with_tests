package roman

import (
	"testing"
)

func TestConvertToRoman(t *testing.T) {

	t.Run("Converting number 1", func(t *testing.T) {
		want := "I"
		got := ConvertToRoman(1)
		if got != want {
			t.Errorf("want %q got %q\n", want, got)
		}
	})

	t.Run("Converting number 2", func(t *testing.T) {
		want := "II"
		got := ConvertToRoman(2)
		if got != want {
			t.Errorf("want %q got %q\n", want, got)
		}
	})

	t.Run("Converting number 3", func(t *testing.T) {
		want := "III"
		got := ConvertToRoman(3)
		if got != want {
			t.Errorf("want %q got %q\n", want, got)
		}
	})

	t.Run("Converting number 4", func(t *testing.T) {
		want := "IV"
		got := ConvertToRoman(4)
		if got != want {
			t.Errorf("want %q got %q\n", want, got)
		}
	})

	t.Run("Converting number 5", func(t *testing.T) {
		want := "V"
		got := ConvertToRoman(5)
		if got != want {
			t.Errorf("want %q got %q\n", want, got)
		}
	})
}
