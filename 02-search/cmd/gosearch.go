package main

import (
	"flag"
	"fmt"
	"log"
	"search/pkg/crawler"
	"search/pkg/crawler/spider"
	"strings"
)

type searcher struct{
	spawn crawler.Interface
}

var store []crawler.Document
var urls = []string{"https://golang.org", "https://go.dev"}

func scanUrls(urls []string, store *[]crawler.Document) {
	s := searcher{}
	s.spawn = spider.New()

	for  _, url := range urls {
		docs, err := s.spawn.Scan(url, 2)
		if err != nil {
			log.Fatal(err)
		}
		for _, i := range docs {
			*store = append(*store, i)
		}
	}
}

func main() {
	var query string
	flag.StringVar(&query, "s", "", "Search")
	flag.Parse()

	if query == "" {
		fmt.Println("Need use with flag `s` for search. For example: gosearch -s go")
		return
	}

	scanUrls(urls, &store)

	for _, v := range store {
		if strings.Contains(strings.ToLower(v.Title), strings.ToLower(query)) {
			fmt.Printf("`%s` found in url: %s\n", query, v.URL)
		}
	}
}
