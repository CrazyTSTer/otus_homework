package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	// Err counter
	var errCount uint32

	// Create task channel
	taskCh := make(chan Task)

	// Create wait group
	wg := sync.WaitGroup{}
	wg.Add(n)

	// Run tasks in separate go routines
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			// Execute tasks from channel
			for task := range taskCh {
				// Execute separate tasks, if it finished with error, increase error counter
				if err := task(); err != nil {
					atomic.AddUint32(&errCount, 1)
				}
			}
		}()
	}

	// Push all tasks to channel
	for _, task := range tasks {
		// Stop pushing tasks if error counter >= M
		if atomic.LoadUint32(&errCount) >= uint32(m) {
			break
		}
		// Push task to channel
		taskCh <- task
	}

	// Close channel
	close(taskCh)
	// Waite for all go-routines finished
	wg.Wait()

	// Check for error count and return an error if error counter >= M
	if errCount >= uint32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
