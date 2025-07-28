package main

import (
	"testing"
	"time"
)

const sleepTime = 2 * time.Second

type result struct {
	string
	bool
}

type WebsiteChecker func(string) bool

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(url)}
		}(url) // pass url as parameter to capture its value, avoiding race condition
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
