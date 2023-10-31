package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func send(ctx context.Context, file int64) error {
	time.Sleep(time.Millisecond)
	if file == 0 {
		return fmt.Errorf("error")
	}
	fmt.Printf("send %d\n", file)
	return nil
}

func TransferFile(ctx context.Context, arr []int64) error {
	eg, egCtx := errgroup.WithContext(ctx)
	for i, _ := range arr {
		func(j int) {
			eg.Go(func() error {
				return send(egCtx, int64(j))
			})
		}(i)
	}
	return eg.Wait()
}

func Query(query string) int {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		r := rand.Intn(5)
		fmt.Println(i, r)
		go func(i, r int) {

			select {
			case <-time.After(time.Duration(r) * time.Second):
				ch <- i
			default:
				fmt.Println("huy tam")
			}
		}(i, r)
	}
	return <-ch
}

func main() {
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	_ = client
}
