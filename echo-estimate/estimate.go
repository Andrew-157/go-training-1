/*
Package that compares the efficiency of echo1, echo2 and echo3, based on the
time of execution
*/
package echoestimate

import (
	"example.com/echo1"
	"example.com/echo2"
	"example.com/echo3"
	"fmt"
	"time"
)

// Runs echo1.Echo(), echo2.Echo(), echo3.Echo() and calculates time each one took and
// prints it to stdout
func EstimateEcho() {
	var (
		firstImpTook  time.Duration
		secondImpTook time.Duration
		thirdImpTook  time.Duration
	)
	// Calculate first implementation
	start := time.Now()
	echo1.Echo()
	firstImpTook = time.Since(start)

	// Calculate second implementation
	start = time.Now()
	echo2.Echo()
	secondImpTook = time.Since(start)

	// Calculate third implementation
	start = time.Now()
	echo3.Echo()
	thirdImpTook = time.Since(start)

	fmt.Printf("First Implementation of Echo took: %s\n", firstImpTook)
	fmt.Printf("Second Implementation of Echo took: %s\n", secondImpTook)
	fmt.Printf("Third Implementation of Echo took: %s\n", thirdImpTook)

	/*
		Sample output:

		root@fedora:~/go_dir/echo-estimate# go run . Hello WHHHD hfhfhfh fhfhfhf hhfhf hfhvhhv kdksk jdjd ddkdkd jdjfhv hfhhfhf hfhfhf
		Hello WHHHD hfhfhfh fhfhfhf hhfhf hfhvhhv kdksk jdjd ddkdkd jdjfhv hfhhfhf hfhfhf
		Hello WHHHD hfhfhfh fhfhfhf hhfhf hfhvhhv kdksk jdjd ddkdkd jdjfhv hfhhfhf hfhfhf
		Hello WHHHD hfhfhfh fhfhfhf hhfhf hfhvhhv kdksk jdjd ddkdkd jdjfhv hfhhfhf hfhfhf
		First Implementation of Echo took: 90.702µs
		Second Implementation of Echo took: 11.7µs
		Third Implementation of Echo took: 10.901µs
	*/
}
