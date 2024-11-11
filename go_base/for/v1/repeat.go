package iteration

func Repeat(s string) string {
	var repeated string
	for range 4 {
		repeated += s
	}

	return repeated
}