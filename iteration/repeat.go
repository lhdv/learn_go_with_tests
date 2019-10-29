package iteration

// Repeat function returns a string repeated 5 times
func Repeat(char string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}

	return repeated
}
