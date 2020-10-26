package sum

// NOTE: there is not yet function overloading
// func Sum(numbers [5]int) int {
// 	sum := 0 // Equivalent: var sum int
// 	// for i := 0; i < len(numbers); i++ {
// 	// 	sum += numbers[i]
// 	// }
// 	// NOTE: Range returns index, value
// 	// NOTE: without '_', i.e., the blank identifier, sum.go will not compile.
// 	for _, value := range numbers {
// 		sum += value
// 	}
// 	return sum
// }

// Sum computes the sum of items in numbers array.
func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// SumAll computes n: len(numbersToSum) arrays and returns
// NOTE: This is how we write a variadic functions
func SumAll(numbersToSum ...[]int) []int {
	count := len(numbersToSum)
	sums := make([]int, count)
	for i, array := range numbersToSum {
		sums[i] = Sum(array)
	}
	return sums
}
