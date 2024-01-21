package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"sync"
)

// CrawlWebpage crawls web pages starting from the rootURL up to a specified maxDepth.
func CrawlWebpage(rootURL string, maxDepth int) ([]string, error) {

	// visited keeps track of URLs that have already been visited to avoid duplication.
	visited := make(map[string]struct{})
	var mu sync.Mutex // mu is used to synchronize access to the visited map.

	// crawl is a recursive function that performs the actual crawling.
	var crawl func(string, int) ([]string, error)
	crawl = func(currentURL string, currentDepth int) ([]string, error) {
		// Stop recursion if the current depth exceeds maxDepth.
		if currentDepth > maxDepth {
			return []string{}, nil
		}

		mu.Lock()
		// Check if the current URL has already been visited.
		if _, exists := visited[currentURL]; exists {
			mu.Unlock()
			return []string{}, nil
		}
		// Mark the current URL as visited.
		visited[currentURL] = struct{}{}
		mu.Unlock()

		var results []string
		if currentDepth <= maxDepth {
			results = append(results, currentURL)
		}

		// Fetch and parse the current URL to find hyperlinks.
		links, err := fetchAndParse(currentURL)
		if err != nil {
			return nil, err
		}

		// If the current depth is less than maxDepth, continue to crawl found links.
		if currentDepth <= maxDepth {
			for _, link := range links {
				absoluteLink := resolveLink(link, currentURL)
				if shouldCrawl(absoluteLink, rootURL) {
					deeperLinks, err := crawl(absoluteLink, currentDepth+1)
					if err != nil {
						return nil, err
					}
					results = append(results, deeperLinks...)
				}
			}
		}

		return results, nil
	}

	// Start the crawl with the root URL and depth 0.
	return crawl(rootURL, 0)
}

// fetchAndParse takes a URL as input, fetches its HTML content, and parses it into a slice of strings.
func fetchAndParse(u string) ([]string, error) {
	// Send an HTTP GET request to the URL.
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching URL: %s, status code: %d", u, resp.StatusCode)
	}

	// Parse the HTML content of the response body.
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string

	// traverse is a recursive function to walk through the HTML nodes.
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			// Loop through the attributes of the <a> tag.
			for _, a := range node.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
					break
				}
			}
		}
		// Recursively call traverse for each child of the current node.
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}
	traverse(doc)

	// Return the slice of extracted links and nil error.
	return links, nil
}

func resolveLink(href, base string) string {
	// Parse the base URL
	baseURL, err := url.Parse(base)
	if err != nil {
		return "" // If the base URL is invalid, return an empty string
	}

	// Parse the link
	link, err := url.Parse(href)
	if err != nil {
		return "" // If the link is invalid, return an empty string
	}

	// Resolve the link against the base URL
	resolvedLink := baseURL.ResolveReference(link)

	return resolvedLink.String()
}

func shouldCrawl(link, rootURL string) bool {
	parsedRootURL, err := url.Parse(rootURL)
	if err != nil {
		return false
	}
	parsedLink, err := url.Parse(link)
	if err != nil {
		return false
	}
	if parsedLink.Scheme != "http" && parsedLink.Scheme != "https" {
		return false
	}
	return parsedLink.Host == parsedRootURL.Host
}
