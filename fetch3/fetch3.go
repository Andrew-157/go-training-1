// Fetch3 prints status code when GET request is made
// for each provided URL
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const urlPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, urlPrefix) {
			url = urlPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch3: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Url %s returned %s status code\n", url, resp.Status)
	}
}
