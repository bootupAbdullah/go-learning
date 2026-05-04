package main

import (
	"cli_timer_akd/pkg/timer"
	"context"
	"fmt"
	"time"
)

func main() {

	now := timer.CurrentTime()
	timer := timer.SetTimer()
	ctx, cancel := context.WithCancel(context.Background())
	for {
		select {
		case <-timer.C:
			if time.Since(now) > 5*time.Second {

				fmt.Println("hello")
				fmt.Println(time.Since(now))
				cancel()

			}
		case <-ctx.Done():
			timer.Stop()
			return
		}
	}
}
