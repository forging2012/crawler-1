package main

import (
	"fmt"
	"log"

	"github.com/artemnikitin/crawler"
)

func main() {
	urls, err := crawler.GetListOfURLWithLimitedConcurrency("https://www.facebook.com/", 2, 150)
	if err != nil {
		log.Fatal("Can't extract URLs from page ", err)
	}

	fmt.Println("Total # of URL:", len(urls))
	if len(urls) <= 100 {
		fmt.Println("URL:")
		for _, v := range urls {
			fmt.Println(v)
		}
	}
}