package iteration

func Repeat(s string, num int) string {
	var repeated string 
	for _ = range num {
		repeated += s
	}
	return repeated
}