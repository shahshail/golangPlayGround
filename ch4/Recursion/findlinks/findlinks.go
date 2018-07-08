package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "find-links: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	// The visit function traverse an HTML node tree, extracts the link from the href attribute of eeach anchor element <a href = "">, appends the link to slice of strngs,
	// and returns resultiung strings
}

func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c) // recursive call
	}

	return links
}
