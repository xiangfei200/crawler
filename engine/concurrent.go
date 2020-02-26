package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan chan Item  //爬虫字段
	RequestProcessor Processor //worker rpc传输
}

//定义worker结构体
type Processor func (Request) (ParseResult,error)

//调度器
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	// 我有一个channel 请问要给哪个worker
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
	//先把worker启动起来，不断的循环去请求页面
	//workcount 表示goroutine的数量
	for i := 0; i < c.WorkerCount; i++ {
		//simple是所有worker共用channel，queued是每个worker一个channel，到底怎么用问Scheduler.WorkerChan()
		//Scheduler.WorkerChan() 初始值是一个nil chan Request，传入createWorker中
		c.createWorker(c.Scheduler.WorkerChan(),out,c.Scheduler)
	}

	//种子页面，先开始
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		//把新的request加入request队列
		c.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		// out赋值给result去循环，然后submit，而submit也是个通道
		result := <-out
		for _, item := range result.Items {
			go func() {
				//保存数据
				c.ItemChan <- item
			}()
			log.Printf("Got profile #%d:%v", itemCount,item)
			itemCount++

			//在这里不能做保存动作，因为他是一个引擎，要赶紧交给worker处理
			//因此我们想把请求和保存请求结果做成一个task
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url){
				continue
			}
			//把新的request加入request队列
			c.Scheduler.Submit(request)
		}
	}
}

//运行种子页面引申出来的各种关联页面
//每个worker 都有 in/out/ready
func (c *ConcurrentEngine)createWorker(in chan Request,out chan ParseResult,ready ReadyNotifier) {
	//每个worker都有channel，因此是他自身的，直接新建即可
	go func() {
		for {
			//告知调取器  我这个worker已经生成需要数量的channel了 可以接收请求了
			//for data := range in{
			//	fmt.Printf("%+v",data)
			//}
			ready.WorkerReady(in)
			//for data := range in{
			//	fmt.Printf("%+v",data)
			//}
			//随着in(即workerChan)的不断改变而改变请求的网页
			// in的来源是 Scheduler告知已经选择了一个worker
			//接收到请求之后赋值给一个变量，从in拿数据
			request := <- in
			//fmt.Printf("%+v",in)

			//做事情，实际去请求页面的动作，在worker中调取到下一个流程的函数
			//从来实现先去实现调用起始页->城市列表->城市->用户列表->用户详情
			//分布式最后一步让worker实现请求rpc服务
			result, err := c.RequestProcessor(request)
			if err != nil {
				continue
			}
			//这里想把result给out,而out本身是一个指针类型，因而改变自身
			//把结果给out
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
