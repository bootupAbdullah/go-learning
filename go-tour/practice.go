package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

tour of go 7/17
func main() {
	fmt.Println(split(17))
}