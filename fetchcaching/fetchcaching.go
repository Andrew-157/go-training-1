// fetchcaching - fetches each url twice to compare the time it took
// to fetch each url for the second time
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	httpUrlPrefix  = "http://"
	httpsUrlPrefix = "https://"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, httpUrlPrefix) || strings.HasPrefix(url, httpsUrlPrefix)) {
			url = httpUrlPrefix + url
		}
		go fetchCache(url, ch, 2)
	}
	for range len(os.Args[1:]) {
		fmt.Println(<-ch)
	}
	fmt.Printf("Total: %.2fs elapsed.\n", time.Since(start).Seconds())
	//Exemplary output:
	// 	root@fedora:~/go_dir# ./fetchcaching.out google.com http://yahoo.com https://youtube.com fjdjdjdffh.com
	// fetchcaching: 1 try fetching url http://fjdjdjdffh.com: Get "http://fjdjdjdffh.com": dial tcp: lookup fjdjdjdffh.com on 127.0.0.53:53: no such host
	// http://google.com: 1 try: [0.25s elapsed|  19261 bytes retrieved], 2 try: [0.11s elapsed|  19280 bytes retrieved]
	// https://youtube.com: 1 try: [0.48s elapsed| 573916 bytes retrieved], 2 try: [0.31s elapsed| 573055 bytes retrieved]
	// http://yahoo.com: 1 try: [1.76s elapsed|1724873 bytes retrieved], 2 try: [0.96s elapsed|1725090 bytes retrieved]
	// Total: 2.72s elapsed.
}

func fetchCache(url string, ch chan<- string, tries int) {
	var start time.Time
	tryTimeElapsed := make(map[int]float64)
	tryBytesRetrieved := make(map[int]int64)
	var triesSlice []int
	for i := 1; i <= tries; i++ {
		// fetch url `tries` times in succession and write time it took for each try
		triesSlice = append(triesSlice, i) // so that we can have sorted printed result in the end
		start = time.Now()
		resp, err := http.Get(url)
		if err != nil {
			ch <- fmt.Sprintf("fetchcaching: %d try fetching url %s: %v", i, url, err)
			return
		}
		// compose filename for a url to where content will be written
		var (
			filename string
			prefix   string
		)
		if strings.HasPrefix(url, httpUrlPrefix) {
			prefix = httpUrlPrefix
		} else if strings.HasPrefix(url, httpsUrlPrefix) {
			prefix = httpsUrlPrefix
		}
		filename = fmt.Sprintf("%s-%d.html", strings.Join(strings.Split(url, prefix)[1:], ""), i)
		f, err := os.Create(filename)
		if err != nil {
			ch <- fmt.Sprintf("fetchcaching: while fetching url %s and creating file: %s: %v", url, filename, err)
			return
		}
		nbytes, err := io.Copy(f, resp.Body)
		resp.Body.Close()
		fileCloseError := f.Close()
		if fileCloseError != nil {
			ch <- fmt.Sprintf("while closing file: %s", filename)
			return
		}
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
			return
		}
		tryTimeElapsed[i] = time.Since(start).Seconds()
		tryBytesRetrieved[i] = nbytes
	}
	resultData := fmt.Sprintf("%s: ", url)
	for index, try := range triesSlice {
		resultData += fmt.Sprintf("%d try: [%.2fs elapsed|%7d bytes retrieved]", try, tryTimeElapsed[try], tryBytesRetrieved[try])
		if index != len(triesSlice)-1 {
			resultData += ", "
		}
	}
	ch <- resultData
}
