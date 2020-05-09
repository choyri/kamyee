package support

import (
	"math/rand"
	"time"
)

type StopRetry struct {
	error
}

func Retry(attempts int, sleep time.Duration, f func() error) error {
	rand.Seed(time.Now().UnixNano())

	if err := f(); err != nil {
		if s, ok := err.(StopRetry); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep += jitter / 2

			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, f)
		}
		return err
	}

	return nil
}
