package iteration

func Repeat(s string, num int) string {
	var repeated string 
	for range num {
		repeated += s
	}
	return repeated
}