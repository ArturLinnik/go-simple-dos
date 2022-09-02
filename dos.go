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
	fmt.Print("URL: ")
	var url string
	fmt.Scanln(&url)

	// Ask for the number of requests
	fmt.Print("Number of requests: ")
	var numReq int
	fmt.Scanln(&numReq)
	fmt.Println("") // Blank line

	wg := sync.WaitGroup{}

	var success int
	var fail int

	start := time.Now()

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
