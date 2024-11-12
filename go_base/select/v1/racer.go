package main

import (
	"net/http"
	"time"
)

func Racer(url1 string, url2 string) (fastUrl string) {
	start := time.Now()
	http.Get(url1)
	Duration1 := time.Since(start)

	start = time.Now()
	http.Get(url2)
	Duration2 := time.Since(start)
	if Duration1 > Duration2 {
		return url2
	}

	return url1
}