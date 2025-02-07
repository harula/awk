package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func timertest() {
	fmt.Println(time.Now())
	time.Sleep(time.Second)
	fmt.Println(time.Now())

	timer := time.NewTimer(time.Second)
	fmt.Println(<-timer.C)

	fmt.Println(<-time.After(time.Second))

	//ctx,cancel := context.WithCancel(context.Background())
	tiker := time.NewTicker(time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println("timer test", i, <-tiker.C)
	}
}

func hw(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("hw ctx done", i)
			return
		default:
			fmt.Println("hw", i, time.Now())
			time.Sleep(time.Second)
		}
	}

}
func ctxtest() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//ctx,cancle := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go hw(ctx, &wg, i)
	}
	wg.Wait()
	fmt.Println("ctx exit...", time.Now())
	//time.Sleep(20*time.Second)

}
func main() {
	fmt.Println("hw")

	//timertest()
	ctxtest()

}
