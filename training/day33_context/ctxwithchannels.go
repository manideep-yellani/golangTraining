package main

import (
	"context"
	"log"
	"time"
)

func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())

	go func() {
		<-time.After(time.Microsecond * 300)
		cancelFunc()
	}()

	//ctx, _ := context.WithTimeout(context.Background(), time.Microsecond*300)

	Fibonacci(ctx)
}

func Fibonacci(ctx context.Context) {
	n := 0
	a := uint64(0)
	b := uint64(1)
	for {
		select {
		case <-ctx.Done():
			return

		default:
			log.Println(a + b)
			a, b = b, a+b
			n++
		}
	}
}
