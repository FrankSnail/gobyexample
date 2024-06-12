package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func PrintWithTrace(ctx context.Context, message string) {
	trace := ""
	trace, _ = ctx.Value("trace").(string)
	fmt.Printf("(Trace id %s) %s.\n", trace, message)
}

func WaitForCancel(ctx context.Context, name string) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Printf("%s quiting...\n", name)
			return
		default:
			fmt.Printf("%s waiting...\n", name)
		}
	}
}

func WithDeadline(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker timeout, %s...\n", ctx.Err())
			return
		default:
			fmt.Printf("Worker is working...\n")
		}
	}
}

func main() {

	// 初始的context对象,其他的context对象要继承自该对象
	ctx := context.Background()

	// WithValue: 传递值
	PrintWithTrace(context.WithValue(ctx, "trace", "0123456"), "A test message...")

	// WithCancel: 可以在外部取消子程序
	ctx2, cancelFUnc := context.WithCancel(ctx)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := range 3 {
		go func() {
			WaitForCancel(ctx2, fmt.Sprintf("worker-%d", i))
			wg.Done()
		}()
	}
	time.Sleep(time.Second * 5)
	cancelFUnc()
	wg.Wait()

	// 带终止时间点, 子routine可以检测是否到时间, 另外一个WithTimeout类似
	ctx3, cancelFunc2 := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
	wg.Add(1)
	go func() {
		WithDeadline(ctx3)
		wg.Done()
	}()
	wg.Wait()
	cancelFunc2()
}
