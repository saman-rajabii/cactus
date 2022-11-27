package cactus

import (
	"sync"
)

func Parallel[T any, R any](sources []T, function func(T, chan R), channel chan R , concurrency int) {
	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(channel)
	}()

	concurrencyRate := (len(sources) + concurrency - 1) / concurrency

	for i := 0; i <= concurrencyRate; i+=concurrency {
		wg.Add(concurrency)
		for j := 0; j < concurrency; j++ {
			go func(s T) {
				defer wg.Done()
				function(s, channel)
			}(sources[i+j])
		}
	}

}