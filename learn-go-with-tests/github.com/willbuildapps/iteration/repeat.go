package iteration

// Repeat an input character for repeat count times.
func Repeat(input string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += input
	}
	return repeated
}
