package main

import (
	"context"
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
// 只要有错误，立即返回
// 等待所有任务执行完成

func main() {
	workerCh := make(chan struct{}, N)
	for i := 0; i < N; i++ {
		workerCh <- struct{}{}
	}

	jobs := make(chan string, len(APIList))
	resCh := make(chan JobRes, len(APIList))
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case res := <-resCh:
				fmt.Printf("job:%v return err:%v\n", res.job, res.err)
				if res.code != ResOk {
					cancel()
					return
				}

			case job, ok := <-jobs:
				if !ok {
					fmt.Println("all jobs sent")
					jobs = nil
				}

				// wait worker available
				<-workerCh

				go workerFn(ctx, workerCh, resCh, job)
			}
		}
	}()

	for _, v := range APIList {
		jobs <- v
	}
	close(jobs)

	time.Sleep(time.Second * 5)
	fmt.Println("ok")
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
	//if a == "e" {
	//	return errors.New(a)
	//}
	return nil
}
