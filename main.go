package main

import (
	"fmt"
	"github.com/ShindeSatish/webcrawler/crawler"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <URL> <DEPTH>")
		return
	}

	url := os.Args[1]
	depth := os.Args[2]

	maxDepth, err := strconv.Atoi(depth)
	if err != nil {
		fmt.Println("Invalid depth value provided, Error:", err)
		return
	}

	fmt.Println("Please wait while we crawl the webpage...")
	links, err := crawler.CrawlWebpage(url, maxDepth)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, link := range links {
		fmt.Println(link)
	}

	fmt.Println("Total number of links:", len(links))
}
