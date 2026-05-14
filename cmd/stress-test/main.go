package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/luizgolima/go-stress-test-challenge/internal/tester"
)

func main() {
	url := flag.String("url", "", "URL of the service to be tested")
	totalRequests := flag.Int("requests", 0, "Total number of requests to be performed")
	concurrency := flag.Int("concurrency", 0, "Number of simultaneous calls")

	flag.Parse()

	if *url == "" || *totalRequests <= 0 || *concurrency <= 0 {
		fmt.Println("Usage: stress-test --url=<url> --requests=<number> --concurrency=<number>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Starting stress test on %s...\n", *url)
	fmt.Printf("Total requests: %d, Concurrency: %d\n", *totalRequests, *concurrency)

	report := tester.Run(*url, *totalRequests, *concurrency)
	report.Print()
}
