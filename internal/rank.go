package internal

import (
	"sort"
	"sync"
	"time"
)

func RankMirrors(urls <-chan string, poolSize int, timeout time.Duration) []Resualt {
	results := make(chan Resualt, 100)
	var wg sync.WaitGroup

	for range poolSize {
		wg.Add(1)
		go ConcurentWorker(urls, results, &wg, timeout)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var ranked []Resualt
	for r := range results {
		if r.Err == nil {
			ranked = append(ranked, r)
		}
	}

	sort.Slice(ranked, func(i, j int) bool {
		return ranked[i].Latency < ranked[j].Latency
	})

	return ranked
}
