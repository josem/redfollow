package main

import (
	"fmt"
	"github.com/josem/redfollow/pkg/output"
	"github.com/josem/redfollow/pkg/redirect"
	"net/url"
	"os"
)

const helpText = `Usage: redfollow url

  Redfollow shows the redirects of a URL and its status codes

  Ex. redfollow http://github.com

  Output:
  	301 - http://github.com
		    ↓
	200 - https://github.com/
`

func printUsage() {
	fmt.Fprint(os.Stderr, helpText)
}

func getURL() string {
	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}

	urlToAnalyze := os.Args[1]

	_, err := url.ParseRequestURI(urlToAnalyze)

	if err != nil {
		fmt.Fprintf(os.Stderr, "The URL %s doesn't seem to be valid\n", urlToAnalyze)
		os.Exit(1)
	}

	return urlToAnalyze
}

func main() {
	urlToAnalyze := getURL()
	redirects := redirect.GetRedirects(urlToAnalyze)
	output.PrintRedirects(redirects)
}
