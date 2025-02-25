// Fetch1 prints the content found at a URL, but instead
// of holding whole stream in a buffer using io.ReadAll(resp.Body), it uses io.Copy(os.Stdout, resp.Body)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
