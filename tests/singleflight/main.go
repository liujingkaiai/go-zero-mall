package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	time.Sleep(9 * time.Second)
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	var g singleflight.Group
	const n = 5
	waited := int32(n)
	done := make(chan struct{})
	key := "https://weibo.com/1227368500/H3GIgngon"
	for i := 0; i < n; i++ {
		go func(j int) {
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})

			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}
		}(i)
	}

	select {
	case <-done:
		fmt.Println("done")
	}
}
