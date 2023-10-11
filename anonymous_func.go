package main

// import (
// 	"fmt"
// )

// ANONYMOUS FUNCTION DEFINED OUTSIDE THE MAIN FUNC
// var (
// 	cube = func(i int) string {
// 			c := i * i * i
// 			return strconv.Itoa(c)
// 	}
// )

// func main() {
// 	x := cube(8)
// 	fmt.Printf("%T %v", x, x)
// }

// ANONYMOUS FUNCTION DEFINED INSIDE FUNC MAIN,

// func main(){
// 	x := func(l int, b int) int {
// 		return l * b
// 	}

// 	fmt.Printf("%T \n", x)
// 	fmt.Println(x(20,30))
// }

// ANONYMOUS FUNCTION WITH WITH ARGUMENTS CALLED DIRECTLY(here the x stores the values that are returned by the function)

// func main(){
// 	x := func(l int, b int) int{
// 		return l * b
// 	}(20, 30)
// 	fmt.Printf("%T \n", x)
// 	fmt.Println(x)
// }
