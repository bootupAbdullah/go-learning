package main

import (
	"fmt"
	"time"
)

func main () {
	timer := time.NewTimer(1 * time.Minute)

	<-timer.C
	fmt.Println("One minue has passed.")
}
