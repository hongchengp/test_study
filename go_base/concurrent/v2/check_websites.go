package main


type CheckWebsiteFunc func(string) bool

type Result struct {
	url string
	res bool
}

func CheckWebsites(cw CheckWebsiteFunc, urls []string) map[string]bool {
	res := make(map[string]bool)
	results := make(chan *Result)
	for _, url := range urls {
		go func(url string) {
			results <- &Result{url, cw(url)}
		}(url)
	}

	for range len(urls) {
		result := <- results
		res[result.url] = result.res
	}

	return res
}