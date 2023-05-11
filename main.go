package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Define the starting URL
	startURL := "https://www.japscan.lol/"

	// Create a set to store visited URLs
	visited := make(map[string]bool)

	// Initialize the web crawler with the starting URL
	crawl(startURL, visited)
}

func crawl(url string, visited map[string]bool) {
	// Check if the URL has already been visited
	if visited[url] {
		return
	}

	// Mark the URL as visited
	visited[url] = true

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the response body using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Extract links from the document
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && link != "" {
			// Process the link (e.g., validate, normalize, etc.)
			link = normalizeURL(link, url)

			// Crawl the link recursively
			crawl(link, visited)
		}
	})

	// Print the visited URL
	fmt.Println("Visited:", url)
}

func normalizeURL(link, baseURL string) string {
	if strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") {
		return link
	}

	if strings.HasPrefix(link, "/") {
		// Construct absolute URL using the base URL
		return strings.TrimRight(baseURL, "/") + link
	}

	// Handle relative URLs
	return baseURL + "/" + link
}
