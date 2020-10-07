package scheduler

import "goSetUp/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
	panic("implement me")
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//send request to worker
func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func() { s.workerChan <- r }()
}
