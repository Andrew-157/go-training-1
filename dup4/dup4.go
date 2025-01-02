// dup4 searches for lines repeating more than once
// either from filenames provided as command-line args or stdin
// and then prints lines that appeared more than once along with their count
// and files where they appeared
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s - %v\n", n, line, filenames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !stringSliceContains(filenames[line], f.Name()) {
			filenames[line] = append(filenames[line], f.Name())
		}
	}
}

func stringSliceContains(slice []string, elem string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == elem {
			return true
		}
	}
	return false
}
