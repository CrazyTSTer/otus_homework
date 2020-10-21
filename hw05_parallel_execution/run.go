package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	//Err counter
	var errCount uint32

	//Create task channel
	taskCh := make(chan Task)

	//Create wait group
	wg := sync.WaitGroup{}
	wg.Add(N)

	//Run tasks in separate go routines
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()

			//Execute tasks from channel
			for task := range taskCh {
				//Execute separate tasks, if it finished with error, increase error counter
				if err := task(); err != nil {
					atomic.AddUint32(&errCount, 1)
				}
			}
		}()
	}

	//Push all tasks to channel
	for _, task := range tasks {
		//Stop pushing tasks if error counter >= M
		if atomic.LoadUint32(&errCount) >= uint32(M) {
			break
		}
		//Push task to channel
		taskCh <- task
	}

	//Close channel
	close(taskCh)
	//Waite for all go-routines finished
	wg.Wait()

	//Check for error count and return an error if error counter >= M
	if errCount >= uint32(M) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
