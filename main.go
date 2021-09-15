package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const N = 5
var APIList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

type JobRes struct {
	job  string
	err  error
}

// 控制最多启动 N 个并发
// 执行一系列任务
// 任一任务有错误，立即返回
// 等待所有任务执行完成
func main() {
	workerCh := make(chan struct{}, N)
	for i := 0; i < N; i++ {
		workerCh <- struct{}{}
	}

	respCh := make(chan JobRes, len(APIList))
	waitCh := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		var num int
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("resp goroutine canceled by main, reason:%s\n", ctx.Err())
				return
			case res := <-respCh:
				num++
				fmt.Printf("job:%v result:%v\n", res.job, res.err)
				if num == len(APIList) || res.err != nil{
					cancel()
					if res.err != nil {
						waitCh <- res.err.Error()
					} else {
						waitCh <- "success"
					}
				}
			}
		}
	}()

	for _, job := range APIList {
		// wait worker available
		<-workerCh
		go dealJob(ctx, workerCh, respCh, job)
	}

	fmt.Println("wait all jobs finished or once error occurs...")
	fmt.Println(<-waitCh)
}

func dealJob(ctx context.Context, workerCh chan struct{}, respCh chan JobRes, job string) {
	//fmt.Println("goroutines:", runtime.NumGoroutine())

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("job:%s canceled by main, reason:%s\n", job, ctx.Err())
			return
		default:
			err := callApi(job)
			respCh <- JobRes{job: job, err: err}

			// reuse worker
			workerCh <- struct{}{}
			return
		}
	}
}

func callApi(a string) error {
	// handler
	time.Sleep(100 * time.Millisecond)
	if a == "e" {
		return errors.New("ops!")
	}
	return nil
}
