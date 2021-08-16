package request

import (
	"time"
)

// retryWithBackoff calls the provided function repeatedly until it succeeds or
// until the retry duration is up.
// modified from https://github.com/maxmind/geoipupdate/blob/main/pkg/geoipupdate/internal/retry.go
func retryWithBackoff(fn func() error, retryFor time.Duration) error {
	start := time.Now()

	for i := uint(0); ; i++ {
		err := fn()
		if err == nil {
			return nil
		}

		currentDuration := time.Since(start)

		waitDuration := 200 * time.Millisecond * (1 << i)

		if currentDuration+waitDuration > retryFor {
			return err
		}

		time.Sleep(waitDuration)
	}
}
