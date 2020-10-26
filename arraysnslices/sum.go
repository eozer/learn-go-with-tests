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
	var sums []int
	for _, array := range numbersToSum {
		sums = append(sums, Sum(array))
	}
	return sums
}

// SumAllTails
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, array := range numbersToSum {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			tail := array[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
