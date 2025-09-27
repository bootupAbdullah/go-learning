package main

import "fmt"


func main() {
	
	abdullah := true 
	job := 4

	if !abdullah {
		switch job {
		case 1:
			fmt.Println("It's not looking good")
		case 2:
			fmt.Println("It's looking great!")
	
	}
} else {
	fmt.Println("Time is on your side")
} 
}