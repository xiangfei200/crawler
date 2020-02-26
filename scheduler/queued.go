package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	// 请求队列
	requestChan chan engine.Request
	//Worker队列 队列每个元素又是一个channel类型的请求 去fetch目标服务器内容
	//把所有的worker channel塞到 workerchan变量的一个Master channel中
	// chan engine.Request 是worker所需要的类型  第一个chan是调度器和worker之间通信所需
	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	//这里希望每个worker有自己的channel，因此返回一个chan
	return make(chan engine.Request)
}

//所有的请求来了之后先放到请求队列中
func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

// 从外部告诉有任务已经准备好，可以负责去接收request了
func (q *QueuedScheduler) WorkerReady (w chan engine.Request){
	q.workerChan <- w
}

//因为有生成，并且改变数值的动作，因此都用指针操作
// 在队列里 两个请求队列
func (q *QueuedScheduler) Run() {
	//生成结构体数据
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ  []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//request 和worker都在排队的时候 都在运行的时候，触发调取器调度动作，出队列
			if len(requestQ) >0 && len(workerQ) >0 {
				//不能在这里把请求通过通道发给worker，因为可也能会卡住
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			// request队列和worker队列是两个独立的任务，没有先后之分，所以要用select去选择
			select {
			case r := <-q.requestChan:
				// 收到就排队 send r to a ? worker
				requestQ = append(requestQ,r)
			case w := <-q.workerChan:
				// 收到就排队 send ? next request to w
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				//初始activeWorker 此时也只能是传参过来的in  因为worker队列除了in没有其他东西
				//然后in 就接收到种子页面的request  值就改变了, 可以触发worker的实际操作
				//出队列，把request给了指定的worker 那么worker本身的值就改变了
				//即engine.createWorker函数中的 in参数值改变了 goroutine中 就可以从in拿数据
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}


		}
	}()
}



