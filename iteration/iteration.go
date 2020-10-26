package iteration

import "strings"

// Repeat prints letter n times
func Repeat(letter string, n int) string {
	// BenchmarkRepeat: 647 ns/op
	// var repeated string
	// for i := 0; i < n; i++ {
	// 	repeated += letter
	// }
	// return repeated
	// BenchmarkRepeat: 112 ns/op
	return strings.Repeat(letter, n)
}
