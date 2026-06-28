package internal

import (
	"net/http"
	"sync"
	"time"
)

// lets create concurency and fix the logic on
// func Get_latency(ch <-chan []string) {
// 	urls := <-ch
// 	go func() {
// 		client := &http.Client{
// 			Timeout: 2 * time.Second,
// 		}

// 		//UrlSlice[0] = "http://google.com"
// 		//resp, err := http.Head(UrlSlice[0])
// 		for _, url := range UrlSlice {
// 			// headobjec, errobject := client.Head(UrlSlice[i])

// 			req, ReqErr := http.NewRequest(http.MethodHead, url, nil)
// 			timenow := time.Now()

// 			resp, responseErr := client.Do(req)

// 			fmt.Println(resp)
// 			thetimesince := time.Since(timenow)

// 			fmt.Println(thetimesince)

// 			if ReqErr != nil && responseErr != nil {
// 				fmt.Printf("some shit happend ! : %s", ReqErr)
// 			}
// 			if resp != nil {
// 				resp.Body.Close()
// 			}
// 		}
// 	}(urls)
// }

// lets write it again with concurrency
// func Getlatency(ch <-chan string ) {
// //////////////////////////////  the logic explaination    ////////////////////////////////
// For each URL received, the worker:
// Creates an HTTP HEAD request (using http.NewRequest).
// Records the start time (time.Now()).
// Executes the request with an http.Client that has a sensible timeout (e.g., 5 seconds).
// Measures the elapsed time (time.Since(start)).
// Constructs a LatencyResult struct containing the URL, the measured latency (or error), and sends that struct into a result channel (resultCh chan LatencyResult).
// i will make it into several fucnctions messuring and workers and sorting
type Resualt struct {
	URL     string
	Latency time.Duration
	Err     error
}

// //////////////////////////////  messuring    ////////////////////////////////
func MessuringTime(url string, timeout time.Duration) Resualt {
	// make a http client and send a head
	client := &http.Client{Timeout: timeout}
	start := time.Now()
	resp, err := client.Head(url)
	latency := time.Since(start)
	if err != nil {
		return Resualt{URL: url, Err: err}
	}
	resp.Body.Close()
	return Resualt{URL: url, Latency: latency}
}

////////////////////////////////    the worker     ////////////////////////////////

func ConcurentWorker(jobs <-chan string, resualt chan<- Resualt, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done()
	// loops over the jobs and creats
	for url := range jobs {
		resualt <- MessuringTime(url, timeout)
	}
}

// //////////////////////////////  rankmirrors    ////////////////////////////////
////////////////////////////////  the logic explaination    ////////////////////////////////
////////////////////////////////  the logic explaination    ////////////////////////////////
////////////////////////////////  the logic explaination    ////////////////////////////////
