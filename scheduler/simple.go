package scheduler

import "crawler/engine"

//send request down to worker chan
type SimpleScheduler struct {
	workerChan chan engine.Request
}

//只改变值本身
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	//在这里再开一个goroutine 解决循环等待问题
	go func() {
		s.workerChan <- r
	}()
}