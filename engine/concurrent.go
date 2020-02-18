package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan chan Item
}

//调度器
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	//输入用了队列
	out := make(chan ParseResult)
	// 在simple中 run实现了chan Request
	c.Scheduler.Run()
	//先把worker启动起来
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.WorkerChan(),out,c.Scheduler)
	}

	//种子页面，先开始
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		c.Scheduler.Submit(r)
	}
	//itemCount := 0
	for {
		// out赋值给result去循环，然后submit，而submit也是个通道
		result := <-out
		for _, item := range result.Items {
			go func() {
				c.ItemChan <- item
			}()
			//log.Printf("Got profile #%d:%v", itemCount,item)
			//itemCount++

			//在这里不能做保存动作，因为他是一个引擎，要赶紧交给worker处理
			//因此我们想把请求和保存请求结果做成一个task
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url){
				continue
			}
			c.Scheduler.Submit(request)
		}
	}
}

//运行种子页面引申出来的各种关联页面
func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier) {
	//每个worker都有channel，因此是他自身的，直接新建即可
	go func() {
		for {
			//告知调取器有 worker准备好了可以接收请求了
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			//这里想把result给out
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool {
	//判断当前url是否在已经请求过的历史切片中
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
