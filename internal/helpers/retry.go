package helpers

import "time"

func Rety(f func() error, sleep time.Duration, attempts int) error {
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
