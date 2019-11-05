package arraysslices

// Sum all integers from array and return the result
func Sum(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}
