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

	wg := sync.WaitGroup{} // Sync go routines

	var success int
	var fail int

	start := time.Now() // Start timer

	for i := 0; i < numReq; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(resp.Request.URL, "---", resp.Status)
			if resp.StatusCode == 200 {
				success += 1
			} else {
				fail += 1
			}
		}()
	}
	wg.Wait()

	finish := time.Since(start)

	fmt.Println("\nSUCCESSFUL:", success, " --- ", "FAILED:", fail)
	fmt.Println("TIME:", finish)
}
