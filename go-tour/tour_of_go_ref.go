package main

import "fmt"



//!! ----> EXMAPLE:
// Tour of Go: 6/17 ---> package level
// func baller(x,y string) (string, string) {
//   return y, x
// }

// func main() { ----> function level
// 	fmt.Println(baller("Hello", "World"))
// }


//!! ----> EXAMPLE:  
// Tour of Go: 7/17
// example: Go return values may be named.
// func split(sum int) (x,y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// }

//!! ----> EXAMPLE:
//  Tour of Go: 8/17
// This var statement declares a list of variables at the "package" level
// var c, python, java bool

// func main() {
	// This var is created at "function" level
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

// Result:
// 0 false false false
// All Go variables have zero values when declared without initialization.




