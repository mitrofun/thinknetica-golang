package main

import (
	"flag"
	"fmt"
	"index/pkg/crawler"
	"index/pkg/crawler/spider"
	"index/pkg/index"
)

var urls = []string{"https://golang.org", "https://go.dev"}

func scan(urls []string, idx index.Service) ([]crawler.Document, error) {
	var result []crawler.Document
	var counter = 0
	s := spider.New()
	for _, url := range urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			return result, err
		}

		for _, item := range docs {
			item.ID = counter
			idx.Add(item.Title, item.ID)
			result = append(result, item)
			counter++
		}
	}
	return result, nil
}

func main() {
	query := flag.String("s", "", "Search word.")
	flag.Parse()

	if *query == "" {
		flag.PrintDefaults()
		return
	}

	i := index.New()
	docs, _ := scan(urls, *i)
	res := i.Search(*query)

	for _, id := range res {
		fmt.Printf("`%s` found in url: %s\n", *query, docs[id].URL)
	}

}
