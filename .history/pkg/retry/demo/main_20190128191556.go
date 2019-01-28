package main

import (
	"errors"
	"github.com/hiruok/msg-pusher/pkg/retry/backoff"
	"log"
	"time"
)

func main() {
	const successOn = 3
	var i = 0

	// This function is successful on "successOn" calls.
	f := func() error {
		i++
		log.Printf("function is called %d. time\n", i)

		if i == successOn {
			log.Println("OK")
			return nil
		}

		log.Println("error")
		return errors.New("error")
	}
	bkoff := backoff.NewExponentialBackOff()
	err := backoff.Retry(f, bkoff)
	if err != nil {
		log.Println("unexpected error: %s", err.Error())
	}
	if i != successOn {
		log.Println("invalid number of retries: %d", i)
	}
	//type nf backoff.Notify
	//i = 0
	bkoff.MaxElapsedTime = 40 * time.Second
	nf := func(e error, duration time.Duration) {
		if e != nil {
			log.Println("notify%s", e)
		} else {

		}
		log.Println(bkoff.GetElapsedTime())
	}
	backoff.RetryNotify(f, bkoff, nf)
}
