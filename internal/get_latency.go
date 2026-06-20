package internal

import (
	"fmt"
	"net/http"
	"time"
)

func Get_latency() {
	UrlSlice := make([]string, 1)

	timenow := time.Now()

	UrlSlice[0] = "http://google.com"
	resp, err := http.Head(UrlSlice[0])
	if err != nil {
		fmt.Printf("some shit happend ! : %s", err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	thetimesince := time.Since(timenow)
	fmt.Println(thetimesince)
}
