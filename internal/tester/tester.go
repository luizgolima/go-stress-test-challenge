package tester

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Error      error
}

type Report struct {
	TotalTime       time.Duration
	TotalRequests   int
	SuccessCount    int
	StatusCounts    map[int]int
}

func Run(url string, totalRequests int, concurrency int) *Report {
	start := time.Now()

	results := make(chan Result, totalRequests)
	jobs := make(chan struct{}, totalRequests)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	// Create workers
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			client := &http.Client{
				Timeout: 10 * time.Second,
			}
			for range jobs {
				resp, err := client.Get(url)
				if err != nil {
					results <- Result{Error: err}
					continue
				}
				results <- Result{StatusCode: resp.StatusCode}
				resp.Body.Close()
			}
		}()
	}

	// Feed jobs
	go func() {
		for i := 0; i < totalRequests; i++ {
			jobs <- struct{}{}
		}
		close(jobs)
	}()

	// Wait for workers to finish
	wg.Wait()
	close(results)

	report := &Report{
		TotalRequests: totalRequests,
		StatusCounts:  make(map[int]int),
	}

	for res := range results {
		if res.Error != nil {
			// Count errors as a special "status code" 0 or just ignore
			report.StatusCounts[0]++
			continue
		}
		if res.StatusCode == http.StatusOK {
			report.SuccessCount++
		}
		report.StatusCounts[res.StatusCode]++
	}

	report.TotalTime = time.Since(start)
	return report
}

func (r *Report) Print() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("             Stress Test Report                  ")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Total time spent:      %v\n", r.TotalTime)
	fmt.Printf("Total requests:        %d\n", r.TotalRequests)
	fmt.Printf("HTTP 200 responses:    %d\n", r.SuccessCount)
	fmt.Println("Status code distribution:")
	for code, count := range r.StatusCounts {
		if code == 0 {
			fmt.Printf("  Errors:              %d\n", count)
		} else {
			fmt.Printf("  HTTP %d:              %d\n", code, count)
		}
	}
	fmt.Println("--------------------------------------------------")
}
