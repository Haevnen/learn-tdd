package iteration

func Repeat(character string, n int) string {
	if n <= 0 {
		return ""
	}

	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}
