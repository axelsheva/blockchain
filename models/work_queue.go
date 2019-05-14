package models

type WorkQueue struct {
	queue chan func()
}

func (workQueue *WorkQueue) Push(fn func()) {
	workQueue.queue <- fn
}

func (workQueue *WorkQueue) process() {
	ch := make(chan bool)

	for {
		go func() {
			fn := <-workQueue.queue
			fn()

			ch <- true
		}()
		<-ch
	}
}

func NewWorkQueue(capacity int) *WorkQueue {
	worker := WorkQueue{
		queue: make(chan func(), capacity),
	}
	go worker.process()

	return &worker
}
