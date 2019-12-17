package waitJob

import (
	"fmt"
	"testing"
	"time"

	"github.com/ivpusic/grpool"
)

func Test_(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(100, 50)
	defer pool.Release()

	// how many jobs we should wait
	pool.WaitCount(10)

	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		count := i

		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()

			fmt.Printf("hello %d\n", count)
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()
}

func Test_2_no_wait(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(100, 50)

	// release resources used by pool
	defer pool.Release()

	// submit one or more jobs to pool
	for i := 0; i < 10; i++ {
		count := i

		pool.JobQueue <- func() {
			fmt.Printf("I am worker! Number %d\n", count)
		}
	}

	// dummy wait until jobs are finished
	time.Sleep(1 * time.Second)
}
