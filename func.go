package main

// import "fmt"

// func operation(a int, b int)(int, int){
// 	sum := a + b
// 	diff := a - b
// 	return sum, diff
// }

// func main(){
// 	sum, diff := operation(20, 10)
// 	fmt.Println(sum, diff)
// }

//NAMED RETURN VALUES(the output parameters are names in the function declaration)
// func operation(a int, b int)(sum int, diff int){
// 	sum = a + b
// 	diff = a - b
// 	return
// }

// func main(){
// 	sum, difference := operation(20, 10)
// 	fmt.Println(sum, " ", difference)
// }

//VARAIDIC FUCNTION(Allows any number of variables of the same type)
// func sumNumbers(numbers ...int)int{
// 	sum := 0
// 	for _ , value := range numbers {
// 		sum += value
// 	}
// 	return sum

// }

// func main(){
// 	fmt.Println(sumNumbers())
// 	fmt.Println(sumNumbers(10))
// 	fmt.Println(sumNumbers(10, 20, 30, 40, 50))
// }

// VARIADIC FUNCTION EXAMPLE 2

// func printDetails(student string, subjects ...string){
// 	fmt.Println("hey", student, " , here are your subjects - ")
// 	for _, sub := range subjects {
// 		fmt.Printf("%s, ", sub)
// 	}
// }

// func main(){
// 	printDetails("Joe", "Physics", "Biology")
// }

// EXAMPLES(funcs with slices)
// func calcSquare(numbers []int) ([]int, bool) {
// 	squares := []int{}
// 	for _, v := range numbers {
// 			squares = append(squares, v*v)
// 	}
// 	return squares, true

// }

// func main() {
// 	nums := [3]int{10, 20, 15}
// 	fmt.Println(calcSquare(nums[:]))
// }