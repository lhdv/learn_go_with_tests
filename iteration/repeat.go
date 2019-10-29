package iteration

const repeatCount = 5

// Repeat function returns a string repeated 5 times
func Repeat(char string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += char
	}

	return repeated
}
