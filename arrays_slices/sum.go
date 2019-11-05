package arraysslices

// Sum all integers from array and return the result
func Sum(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}

// SumAll will sum all elements for each slice and return them on a new slice
func SumAll(numbers ...[]int) []int {

	sumAll := make([]int, len(numbers))

	for i, n := range numbers {
		sumAll[i] = Sum(n)
	}

	return sumAll
}

// SumAllTails sums from the second item on slice until its end
func SumAllTails(numbers ...[]int) []int {

	sumAll := make([]int, len(numbers))

	for i, n := range numbers {
		if len(n) == 0 {
			sumAll[i] = 0
		} else {
			sumAll[i] = Sum(n[1:])
		}
	}

	return sumAll
}
