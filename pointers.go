package main

// import "fmt"

// func main(){
// 	i := 10
// 	fmt.Printf("%T %v \n", &i, &i)
// 	fmt.Printf("%T %v \n", *(&i), *(&i))
// }

//DECLARING AND INITIALIZING A POINTER USING THREE METHODS
//ONE

// func main(){
// 	i := 10
// 	var ptr_i *int = &i
// 	fmt.Println(ptr_i)
// }

//TWO - HERE WE DO NOT SPECIFY THE DATA TYPE BUT IT WILL BE INTERNALLY DETERMINED BY THE COMPILER

// func main(){
// 	s := "hello"
// 	var ptr_s = &s
// 	fmt.Println(ptr_s)
// }

// THREE - USING THE SHORTHAND OPERATOR

// func main(){
// 	s := "hello"
// 	ptr_s := &s
// 	fmt.Println(ptr_s)
// }

//USING THE DEREFERNCING OPERATOR WITH POINTERS

// func main(){
// 	s := "hello"
// 	fmt.Printf("%T %v \n", s, s)

// 	ps := &s
// 	fmt.Printf("%T \n", &s)
// 	*ps = "world"

// 	fmt.Printf("%T %v \n",s,s)
// }