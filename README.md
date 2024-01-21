# Web Crawler in Go
## Overview
This project is a simple web crawler written in Go. It is designed to crawl web pages starting from a given URL, following links up to a specified depth. The crawler fetches web pages, parses them for links, and continues this process recursively until the maximum depth is reached.

## Features

Recursive Crawling: Follows links found on web pages to a specified depth.

Depth Control: Limits crawling to a user-specified depth to avoid extensive crawling.

Visited URL Tracking: Keeps track of visited URLs to avoid duplicate processing.


## Getting Started
### Prerequisites
Go (version 1.18 or higher)

### Installation
Clone the repository to your local machine:
```
https://github.com/ShindeSatish/webcrawler.git

cd webcrawler
```

### Usage
Run the crawler with the following command:
`go run cmd/main.go <URL> <Depth>` replace URL abd Depth values

Example 
```
go run main.go https://www.w3schools.com 1
```

NOTE - This program require some time to fetch all the URLs according to the links presents. So Please wait until it fetches all the URLs.