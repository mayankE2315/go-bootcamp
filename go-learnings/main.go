package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result holds the URL and its status
type Result struct {
	URL    string
	Status string
}

func check(ctx context.Context, url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- Result{URL: url, Status: "DOWN"}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		results <- Result{URL: url, Status: "DOWN"}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		results <- Result{URL: url, Status: "UP"}
	} else {
		results <- Result{URL: url, Status: "DOWN"}
	}
}

func printer(results <-chan Result, printerDone chan struct{}) {
	defer close(printerDone)
	for result := range results {
		if result.Status == "UP" {
			fmt.Printf("✅ %s is UP\n", result.URL)
		} else {
			fmt.Printf("❌ %s is DOWN\n", result.URL)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
		"https://instagram.com",
		"https://site-not-present.io",
		"https://youtube.com",
		"https://linkedin.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
	}

	results := make(chan Result, len(urls))

	printerDone := make(chan struct{})

	var wg sync.WaitGroup

	go printer(results, printerDone)

	for _, url := range urls {
		wg.Add(1)
		go check(ctx, url, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	select {
	case <-printerDone:
		fmt.Println("All checks completed!")
	case <-ctx.Done():
		fmt.Println("Timeout reached! Some checks may still be running...")
	}
}
