package waitJob

import (
	"fmt"
	"testing"
	"time"

	"github.com/ivpusic/grpool"
)

func Test_wait_job(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(5, 3)
	defer pool.Release()

	// how many jobs we should wait
	pool.WaitCount(3)

	// submit one or more jobs to pool
	for i := 0; i < 100; i++ {
		count := i

		pool.JobQueue <- func() {
			// say that job is done, so we can know how many jobs are finished
			defer pool.JobDone()

			fmt.Printf("hello %d\n", count)
			time.Sleep(5 * time.Second)
		}
	}

	// wait until we call JobDone for all jobs
	pool.WaitAll()
}

func Test_2_no_wait(t *testing.T) {
	// number of workers, and size of job queue
	pool := grpool.NewPool(2, 10)

	// release resources used by pool
	defer pool.Release()

	// submit one or more jobs to pool
	for i := 0; i < 100; i++ {
		count := i

		pool.JobQueue <- func() {
			fmt.Printf("I am worker! Number %d\n", count)
			time.Sleep(5 * time.Second)
		}
		fmt.Println(" loop waiting....", count)
	}

	// dummy wait until jobs are finished
	fmt.Println(" waiting.....")
	time.Sleep(1 * time.Second)
}
