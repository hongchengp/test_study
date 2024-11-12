package main

import (
	"fmt"
	"testing"
)

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 0)
	for i := range 10 {
		urls = append(urls, fmt.Sprintf("yyj %d", i))
	}

	for range b.N {
		CheckWebsites(CheckWebsite, urls)
	}

}