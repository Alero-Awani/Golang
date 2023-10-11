package main

// import (
// 	"fmt"
// 	"strings"
// )

// DEFER STATEMENT - defers the execution of a function until the surrounding functions return

// func printName(str string){
// 	fmt.Println(str)
// }

// func printRollNo(rno int){
// 	fmt.Println(rno)
// }

// func printAddress(adr string){
// 	fmt.Println(adr)
// }

// func main(){
// 	printName("Alero")
// 	defer printRollNo(23)
// 	printAddress("Street-32")
// }

//EXAMPLE 1
// in the func definition we can see we are expecting two strings
// func getString(str string) (string, string) {
// 	return strings.ToLower(str), strings.ToUpper(str)
// }
// // this is a trick, the one that returns the upper was named lower
// func main() {
// 	_, lower := getString("BROWSER")
// 	fmt.Println(lower)
// }