package main

import (
	"flag"
	"fmt"
	"index/pkg/crawler"
	"index/pkg/crawler/spider"
	"index/pkg/index"
	"log"
)

var urls = []string{"https://golang.org", "https://go.dev"}
var scanIndex = index.Index{}

func scan(urls []string) ([]crawler.Document, error) {
	var result [] crawler.Document
	var counter = 0
	s := spider.New()
	for  _, url := range urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			log.Fatal(err)
			return result, err
		}

		for _, item := range docs {
			item.ID = counter
			scanIndex.Add(item.Title, item.ID)
			result = append(result, item)
			counter = counter + 1
		}
	}
	return result, nil
}

func main() {
	query := flag.String("s", "", "Search")
	flag.Parse()

	if *query == "" {
		fmt.Println("Need use with flag `s` for search. For example: gosearch -s go")
		return
	}

	scanIndex.New()
	docs, err  := scan(urls)
	if err != nil {
		log.Fatal(err)
		return
	}

	res:= scanIndex.Search(*query)

	for _, id:= range res {
		fmt.Printf("`%s` found in url: %s\n", *query, docs[id].URL)
	}

}
