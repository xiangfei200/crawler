package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//调度器
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	//所有worker暂用同一个输入
	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChan(in)
	//先把worker启动起来
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	//种子页面，先开始
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		// out赋值给result去循环，然后submit，而submit也是个通道
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got Item #%d:%v", itemCount,item)
			itemCount++
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

//运行种子页面引申出来的各种关联页面
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			//有循环等待
			//这里想把result给out
			out <- result
		}
	}()
}
