package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	// Ask for the URL of the webpage
	var url string
	fmt.Print("URL: ")
	_, err := fmt.Scanln(&url)
	if err != nil {
		log.Fatal(err)
	}

	// Ask for the number of requests
	var numReq int
	fmt.Print("Number of requests: ")
	_, err = fmt.Scanln(&numReq)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("") // Blank line

	// Create request structure
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now() // Start timer

	wg := sync.WaitGroup{} // Sync go routines

	// Send requests
	var success int
	var fail int

	for i := 0; i < numReq; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(resp.Request.URL, "---", resp.Status)
			if resp.StatusCode != http.StatusOK {
				fail += 1
			} else {
				success += 1
			}
		}()
	}
	wg.Wait()

	finish := time.Since(start) // Stop timer

	// Print results
	fmt.Println("\nSUCCESSFUL:", success, " --- ", "FAILED:", fail)
	fmt.Println("TIME:", finish)
}
