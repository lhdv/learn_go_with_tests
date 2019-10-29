package iteration

const repeatCount = 5

// Repeat function returns a string repeated how many times you ask.
// If you don't know how many times you want, the default is 5 by the way.
func Repeat(char string, times int) string {

	var repeated string

	if times <= 0 {
		times = repeatCount
	}

	for i := 0; i < times; i++ {
		repeated += char
	}

	return repeated
}
