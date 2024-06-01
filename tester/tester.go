package tester

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Run(url string, requests int, concurrency int) {
	workChannel := make(chan struct{}, requests)
	resultsChannel := make(chan int, requests)

	var waitGroup sync.WaitGroup
	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		waitGroup.Add(1)
		go worker(&waitGroup, url, workChannel, resultsChannel)
	}

	for i := 0; i < requests; i++ {
		workChannel <- struct{}{}
	}
	close(workChannel)

	waitGroup.Wait()
	close(resultsChannel)
	totalTime := time.Since(startTime)

	showReport(resultsChannel, totalTime)
}

func worker(waitGroup *sync.WaitGroup, url string, workChannel chan struct{}, resultsChannel chan int) {
	defer waitGroup.Done()
	for range workChannel {
		resp, err := http.Get(url)
		if err != nil {
			resultsChannel <- 0
			continue
		}
		resultsChannel <- resp.StatusCode
	}
}

func showReport(results chan int, totalTime time.Duration) {
	totalRequests := 0
	statusCount := make(map[int]int)

	for status := range results {
		totalRequests++
		statusCount[status]++
	}

	fmt.Printf("Relatório do Teste de Stress\n")
	fmt.Printf("Tempo total: %v\n", totalTime)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("Requests com status 200: %d\n", statusCount[200])
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for status, count := range statusCount {
		if status != 200 {
			fmt.Printf("%d: %d\n", status, count)
		}
	}
}
