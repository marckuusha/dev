package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
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

func main() {
	err := TransferFile(context.Background(), []int64{1, 2, 3, 4, 5})
	if err != nil {
		log.Error().Msg(err.Error())
	}

}
