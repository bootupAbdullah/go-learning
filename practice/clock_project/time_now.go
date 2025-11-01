package main

import (
	"fmt" 
	"time"
)

// 1. Create function to check current time.
func current_time () {
 
	// 2. Declare var to hold value (time)
	// var current_time_VA int32

	t := time.Now().Local()

	// tz, _ := time.LoadLocation("America/New_York")
	// vaHour := t.In(tz).Hour()

	// return int32(vaHour)

	fmt.Println(t)
}


func main() {
	// 3. Call time_now function
	current_time()

}