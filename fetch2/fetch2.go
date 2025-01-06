// Fetch2 prints the content found at a URL,
// adds prefix 'http://' to the provided url, if it is missing
package main

import (
	"fmt"
	"io"
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
			fmt.Fprintf(os.Stderr, "fetch2: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch2: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
