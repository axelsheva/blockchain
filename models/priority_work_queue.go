package models

type PriorityWorkQueue struct {
	mainQueue chan func()
	queue     chan func()
}

func (workQueue *PriorityWorkQueue) PushMain(fn func()) {
	workQueue.mainQueue <- fn
}

func (workQueue *PriorityWorkQueue) Push(fn func()) {
	workQueue.queue <- fn
}

func (workQueue *PriorityWorkQueue) process() {
	ch := make(chan bool)

	for {
		go func() {
			select {
			case mainFn := <-workQueue.mainQueue:
				mainFn()
				ch <- true
				return
			default:
			}

			select {
			case fn := <-workQueue.queue:
				fn()
				ch <- true
			}
		}()

		<-ch
	}
}

func NewPriorityWorkQueue(mainCapacity int, capacity int) *PriorityWorkQueue {
	worker := PriorityWorkQueue{
		mainQueue: make(chan func(), mainCapacity),
		queue:     make(chan func(), capacity),
	}
	go worker.process()

	return &worker
}
