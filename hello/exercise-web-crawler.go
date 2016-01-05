package main

import (
  "fmt"
  //"sync"
)


type Fetcher interface {
  // returns body of url and slice of urls found on web page
  Fetch(url string) (body string, urls []string, err error)
}

func contains(slice []string, target string) bool {
  for _, elem := range slice {
    if elem == target {
      return true
    }
  }
  return false
}
// Uses Fetcher to recursively crawl pages starting at url, to max depth
func Crawl(url string, depth int, fetcher Fetcher) (body string, urls []string, err error) {
  foundUrls := SafeCounter{v: make(map[string]int)}
  return _crawl(url, depth, fetcher, foundUrls)
}

func _crawl(url string, depth int, fetcher Fetcher, foundUrls SafeCounter) (body string, urls []string, err error) {
  // don't fetch same url 2x
  if depth <= 0 || foundUrls.Value(url) != 0 {
    return
  }
  body, urls, err = fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }
  foundUrls.Inc(url)
  fmt.Printf("found %s %q\n", url, body)
  // fetch urls in parallel
  done := make(chan bool)
  for _, u := range urls {
    go func(url string) {
      _crawl(url, depth-1, fetcher, foundUrls)
      done <- true
    }(u)
  }
  // only proceed after finished handling all child urls
  for range urls {
    <-done
  }
  return

}

func runWebCrawler() {
  fmt.Println("* Web Crawler exercise *")
  // execute on hard-coded results
  Crawl("http://golang.org/", 4, fetcher)
}

// hard-coded results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
  body string
  urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
  if res, ok := f[url]; ok {
    return res.body, res.urls, nil
  }
  return "", nil, fmt.Errorf("not found %s", url)
}

var fetcher = fakeFetcher {
  "http://golang.org/": &fakeResult{
    "The Go Programming Language",
    []string{
      "http://golang.org/pkg/",
      "http://golang.org/cmd/",
    },
  },
  "http://golang.org/pkg/": &fakeResult{
    "Packages",
    []string{
      "http://golang.org/",
      "http://golang.org/cmd/",
      "http://golang.org/pkg/fmt/",
      "http://golang.org/pkg/os/",
    },
  },
  "http://golang.org/pkg/fmt/": &fakeResult{
    "Package fmt",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },
  "http://golang.org/pkg/os/": &fakeResult{
    "Package os",
    []string{
      "http://golang.org/",
      "http://golang.org/pkg/",
    },
  },

}