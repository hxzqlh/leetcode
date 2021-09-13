package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

const (
	N       = 5
	ResOk   = 0
	ResFail = 1
)

type JobRes struct {
	job  string
	code int
	err  error
}

var APIList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

// 最多 N 个并发
// 执行一系列任务
// 任一任务有错误，立即返回
// 等待所有任务执行完成
func main() {
	workerCh := make(chan struct{}, N)
	for i := 0; i < N; i++ {
		workerCh <- struct{}{}
	}

	resCh := make(chan JobRes, len(APIList))
	waitCh := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		var num int
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("res goroutine canceled by main, reason:%s\n", ctx.Err())
				return
			case res := <-resCh:
				num++
				fmt.Printf("job:%v return err:%v\n", res.job, res.err)
				if num == len(APIList) || res.code != ResOk{
					cancel()
					waitCh <- struct{}{}
				}
			}
		}
	}()

	for _, job := range APIList {
		// wait worker available
		<-workerCh
		go workerFn(ctx, workerCh, resCh, job)
	}

	<-waitCh
	time.Sleep(time.Second)
	fmt.Println("main exit")
}

func workerFn(ctx context.Context, workerCh chan struct{}, resCh chan JobRes, job string) {
	fmt.Println("goroutines:", runtime.NumGoroutine())

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("job:%s canceled by main, reason:%s\n", job, ctx.Err())
			return
		default:
			code := ResOk
			err := callApi(job)
			if err != nil {
				code = ResFail
			}

			resCh <- JobRes{job: job, code: code, err: err}

			// reuse worker
			workerCh <- struct{}{}
			return
		}
	}
}

func callApi(a string) error {
	// handler
	time.Sleep(200 * time.Millisecond)
	if a == "e" {
		return errors.New(a)
	}
	return nil
}
