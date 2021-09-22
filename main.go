package main

import (
	"context"
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
	jobCh := make(chan string, len(APIList))
	respCh := make(chan JobRes, len(APIList))
	waitCh := make(chan string)

	// produce job
	for _,  v := range APIList {
		jobCh <- v
	}
	close(jobCh)

	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= N; i++ {
		// consume job
		go workerLoop(ctx, i, jobCh, respCh)
	}

	// collect job resp
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

	fmt.Println("wait all jobs finished or once error occurs...")
	fmt.Println(<-waitCh)
}

func workerLoop(ctx context.Context, index int, jobCh chan string, respCh chan JobRes) {
	//fmt.Println("goroutines:", runtime.NumGoroutine())

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker:%v canceled by main, reason:%s\n", index, ctx.Err())
			return
		case job, ok := <-jobCh:
			if !ok {
				fmt.Println("no job any more!")
				return
			}
			err := callApi(job)
			respCh <- JobRes{job: job, err: err}
		}
	}
}

func callApi(a string) error {
	// handler
	time.Sleep(100 * time.Millisecond)
	//if a == "e" {
	//	return errors.New("ops!")
	//}
	return nil
}
