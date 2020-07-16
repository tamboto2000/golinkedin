package linkedin

import "sync"

type threadPool struct {
	mutex *sync.Mutex
	wg    *sync.WaitGroup
	err   error
}

func newThreadPool() *threadPool {
	return &threadPool{
		mutex: new(sync.Mutex),
		wg:    new(sync.WaitGroup),
	}
}

func (th *threadPool) add(c int) {
	th.wg.Add(c)
}

func (th *threadPool) run(runnable func() error) {
	if err := runnable(); err != nil {
		th.mutex.Lock()
		if th.err == nil {
			th.err = err
		}

		th.mutex.Unlock()
	}

	th.wg.Done()
}

func (th *threadPool) wait() {
	th.wg.Wait()
}

func (th *threadPool) getError() error {
	return th.err
}
