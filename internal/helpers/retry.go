package helpers

import (
	"fmt"
	"time"
)

func Retry(f func() error, sleep time.Duration, attempts int) error {
	fmt.Println("retrying...")
	var err error
	for i := 0; i < attempts; i++ {
		time.Sleep(sleep)
		err = f()
		if err == nil {
			return nil
		}
	}

	return err
}
