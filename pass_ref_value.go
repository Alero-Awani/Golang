package main

// import "fmt"

// PASSING BY REFERENCE

// func modify(s *string){
// 	*s = "world"
// }

// func main(){
// 	a := "hello"

// 	fmt.Println(a)
// 	modify(&a)

// 	fmt.Println(a)
// }

//SLICES AND MAPS ARE BASICALLY PASSED BY REFERENCE BY DEFAULT
// func modify(s []int){
// 	s[0] = 100

// }

// func main(){
// 	slice := []int{10,20,30}
// 	fmt.Println(slice)

// 	modify(slice)
// 	fmt.Println(slice)
// }