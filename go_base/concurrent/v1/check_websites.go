package main

type CheckWebsiteFunc func(string) bool

func CheckWebsites(cw CheckWebsiteFunc, urls []string) map[string]bool{
	res := make(map[string]bool)
	for _, url := range urls {
		res[url] = cw(url)
	}
	return res
}