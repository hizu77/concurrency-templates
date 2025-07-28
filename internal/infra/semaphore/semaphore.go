package semaphore

type semaphore struct {
	ch chan struct{}
}

func New(limit int) *semaphore {
	return &semaphore{
		ch: make(chan struct{}, limit),
	}
}

func (s *semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.ch
}
