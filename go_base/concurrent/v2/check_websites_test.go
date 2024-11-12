package main

import (
	"reflect"
	"testing"
)

func TestCheckWebsite(t *testing.T) {
	urls := make([]string, 0)
	urls = append(urls, "yyj")
	urls = append(urls, "hcp", "mhp", "cjc")
	got := CheckWebsites(CheckWebsite, urls)
	want := map[string]bool {
		"yyj": false, 
		"hcp": true,
		"mhp": true,
		"cjc": true,
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}