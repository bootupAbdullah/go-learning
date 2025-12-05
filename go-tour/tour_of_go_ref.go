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


//!! ----> EXAMPLE:
// Tour of Go: 9/17
// Initialize variables
// var i, j int = 1, 2


// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

// result: 
// 1 2 true false "no!"

//!! ----> EXAMPLE:
// Tour of Go: 9/17
// var (
// 	ToBe bool 		= false
// )

// func main() {
// 	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
// }

// result:
// Type: bool Value: false


//!! ----> EXAMPLE:
// Tour of Go: 12/17
// func main() {
// 	var i int
// 	var f float64
// 	var b bool
// 	var s string
// 	fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }

// return:
// 0 0 false ""