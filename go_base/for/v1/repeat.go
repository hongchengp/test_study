package iteration

func Repeat(s string) string {
	var repeated string
	for _ = range 4 {
		repeated += s
	}

	return repeated
}